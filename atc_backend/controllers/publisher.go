package controllers

import (
	"atc_backend/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type PublisherController struct {
	beego.Controller
}

func (u *PublisherController) PostAtc() {
	transid, err := models.PostAtc(u.GetString("userid"), u.GetString("msgtype"), u.GetString("content"))

	var ret_code int
	var ret_data string

	if err != nil {
		ret_code = 1001
		ret_data = ""
	} else {
		ret_code = 1000
		ret_data = transid
	}

	u.Data["json"] = JsonResponse{
		Code: ret_code,
		Data: ret_data,
	}
	u.ServeJSON()
}
