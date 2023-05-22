package common

import (
	"reflect"
	sqlc "tienpvse1/go-fiber-server/src/modules/generated/sql"
)

func (bundler Bundler) Bundle(router *Router, queries *sqlc.Queries) {

  var queriesInstance *sqlc.Queries
  var routerInstance *Router
  if bundler.Queries != nil {
    queriesInstance = bundler.Queries
  }else {
    queriesInstance = queries
  }
  if bundler.Router != nil {
    routerInstance = bundler.Router
  }else {
    routerInstance = router
  }
	if bundler.Imports != nil && len(bundler.Imports) > 0 {
		for _, importInstance := range bundler.Imports {
			importInstance.Bundle(bundler.Router, bundler.Queries)
		}
	}
	// service pool is where we store dependencies-injectED services
	// use key-pair value store to make use of the injected service later
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
				serviceMetadata.Elem().FieldByName(fieldName).Set(reflect.ValueOf(queriesInstance))
				continue
			}
		}
		// v are now has been injected with sqlc queries, add it to the service poll
		servicePools[serviceName] = v
	}

	// loop through list of controller, inject service if need
	for _, controller := range bundler.Controllers {
		controllerMetadata := reflect.ValueOf(controller)
		for i := 0; i < controllerMetadata.Elem().Type().NumField(); i++ {
			field := controllerMetadata.Elem().Type().Field(i)
			injectTag := field.Tag.Get("inject")
			fieldName := field.Name
			if len(injectTag) == 0 {
				continue
			}

			// inject the service
			controllerMetadata.Elem().FieldByName(fieldName).Set(reflect.ValueOf(servicePools[injectTag]))
		}

		// inject the router builder instance if need
		if isFieldExist(controllerMetadata, "Router") {
			controllerMetadata.Elem().FieldByName("Router").Set(reflect.ValueOf(routerInstance))
		}
		controller.Create()
	}
}

// check either the field exist in the struct type
func isFieldExist(value reflect.Value, fieldName string) bool {
	_, ok := value.Elem().Type().FieldByName(fieldName)
	return ok
}
