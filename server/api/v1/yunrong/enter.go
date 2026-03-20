package yunrong

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ YrUserApi }

var yrUserService = service.ServiceGroupApp.YunrongServiceGroup.YrUserService
