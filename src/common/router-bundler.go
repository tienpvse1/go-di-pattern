package common

import (
	"reflect"
)

func (bundler Bundler) Bundle() {
	servicePools := map[string]any{}
	for _, v := range bundler.Services {
		serviceMetadata := reflect.ValueOf(v)
		serviceName := serviceMetadata.Elem().Type().Name()
		// perform the dependencies injection
		for i := 0; i < serviceMetadata.Elem().Type().NumField(); i++ {
			// alias the field variable
			field := serviceMetadata.Elem().Type().Field(i)
			injectTag := field.Tag.Get("inject")
			fieldName := field.Name
			if injectTag == "sqlc_queries" {
				serviceMetadata.Elem().FieldByName(fieldName).Set(reflect.ValueOf(bundler.Queries))
        continue
			}
		}
		// v are now has been injected with sqlc queries, add it to the service poll
		servicePools[serviceName] = v
	}
	for _, controller := range bundler.Controllers {
		controllerMetadata := reflect.ValueOf(controller)
		for i := 0; i < controllerMetadata.Elem().Type().NumField(); i++ {
			field := controllerMetadata.Elem().Type().Field(i)
			injectTag := field.Tag.Get("inject")
			fieldName := field.Name
			if len(injectTag) == 0 {
				continue
			}

			controllerMetadata.Elem().FieldByName(fieldName).Set(reflect.ValueOf(servicePools[injectTag]))
		}
    // inject the app instance if need  
		if isFieldExist(controllerMetadata, "Router") {
			controllerMetadata.Elem().FieldByName("Router").Set(reflect.ValueOf(bundler.Router))
		}
    controller.Create()
	}
}

func isFieldExist(value reflect.Value, fieldName string) bool {
	_, ok := value.Elem().Type().FieldByName(fieldName)
	return ok
}
