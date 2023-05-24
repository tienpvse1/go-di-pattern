package user

import "tienpvse1/go-fiber-server/src/common"

var UserModule = common.Bundler {
  Controllers: []common.IController{
    new(UserController),
  },
  Services: []interface {} {
    new(UserService),  
  },
}
