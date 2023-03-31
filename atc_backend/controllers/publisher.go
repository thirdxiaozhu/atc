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
	transid, err := models.PostAtc(u.GetString("userid"), u.GetString("msgtype"), u.GetString("content"), u.GetString("timestamp"))

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

func (u *PublisherController) GetAtcs() {
	paralist := [4]string{u.GetString("publisher"), u.GetString("company"), u.GetString("starttime"), u.GetString("endtime")}

	atcs := models.GetAtcs(u.GetString("userid"), paralist[:])

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: &atcs,
	}
	u.ServeJSON()
}

func (u *PublisherController) PostEdit() {
	paralist := [4]string{u.GetString("userid"), u.GetString("content"), u.GetString("ID"), u.GetString("timestamp")}

	transid := models.EditAtc(paralist[:])

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: transid,
	}
	u.ServeJSON()
}

func (u *PublisherController) PostDelete() {
	transid := models.DeleteAtc(u.GetString("userid"), u.GetString("ID"))
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: transid,
	}
	u.ServeJSON()
}
