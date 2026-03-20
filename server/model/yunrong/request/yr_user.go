package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type YrUserSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Username       *string     `json:"username" form:"username"`
	request.PageInfo
}

type AdminChangePassword struct {
	UserID   uint   `json:"userID"`
	Password string `json:"password"`
}

type YrUserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
