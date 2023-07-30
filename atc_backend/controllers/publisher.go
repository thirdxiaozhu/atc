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
	transid, err := models.PostAtc(u.GetString("userid"), u.GetString("msgtype"), u.GetString("content"), u.GetString("timestamp"), u.GetString("flight"))

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

func (u *PublisherController) PostMultiAtc() {
	transids, err := models.PostMultiAtc(u.GetString("userid"), u.GetString("msgs"), u.GetString("timestamp"))
	//transid, err := models.PostAtc(u.GetString("userid"), u.GetString("msgs"), u.GetString("timestamp"))

	var ret_code int
	var ret_data []string

	if err != nil {
		ret_code = 1001
		ret_data = nil
	} else {
		ret_code = 1000
		ret_data = transids
	}

	u.Data["json"] = JsonResponse{
		Code: ret_code,
		Data: ret_data,
	}
	u.ServeJSON()
}

func (u *PublisherController) GetAtcs() {
	paralist := [5]string{u.GetString("publisher"), u.GetString("company"), u.GetString("starttime"), u.GetString("endtime"), u.GetString("flight")}

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

func (u *PublisherController) GetArns() {
	arn_p := models.GetArns()
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: &arn_p,
	}
	u.ServeJSON()
}

func (u *PublisherController) GetRoutes() {
	routes_p := models.GetRoutes()
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: &routes_p,
	}
	u.ServeJSON()
}
func (u *PublisherController) GetAuths() {
	auth_p := models.GetAuths()
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: &auth_p,
	}
	u.ServeJSON()
}

func (u *PublisherController) GetLinkOptions() {

	options := models.GetLinkOptions()
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: &options,
	}
	u.ServeJSON()
}

func (u *PublisherController) PostRegLink() {
	ret := models.RegistLink(u.GetString("opt_arn"), u.GetString("opt_route"), u.GetString("opt_auth"))

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: ret,
	}
	u.ServeJSON()
}

func (u PublisherController) GetRegLink() {
	ret := models.GetRegistLink()
	logger.Println(ret)
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: ret,
	}
	u.ServeJSON()
}
