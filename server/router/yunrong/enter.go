package yunrong

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ YrUserRouter }

var yrUserApi = api.ApiGroupApp.YunrongApiGroup.YrUserApi
