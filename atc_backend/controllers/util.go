package controllers

import (
	"atc_backend/models"
	"log"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Return struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Leaf  string `json:"leaf"`
}

type JsonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Operations about Users
type UtilController struct {
	beego.Controller
}

var logger *log.Logger

func init() {
	logger = logs.GetLogger()
}

func (u *UtilController) GetCompanies() {

	companies := models.GetCompaniesByRole(u.GetString("role"))

	var companiesret []Return
	for _, company := range *companies {
		ret := Return{}
		ret.Value = company.Name
		ret.Label = company.CNname
		ret.Leaf = "level >= 1"
		companiesret = append(companiesret, ret)
	}

	//ret.Code = 1000
	//ret.Data = companiesret

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: companiesret,
	}

	u.ServeJSON()
}

func (u *UtilController) GetCheckCompanies() {
	companies := models.GetCompaniesByRole("0")

	var companiesret []Return
	for _, company := range *companies {
		ret := Return{}
		ret.Value = company.Name
		ret.Label = company.CNname
		companiesret = append(companiesret, ret)
	}

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: companiesret,
	}
	u.ServeJSON()
}

func (u *UtilController) GetCheckPublishers() {
	company_names := strings.Split(u.GetString("company"), ",")
	var publishersret []Return

	for _, company_name := range company_names {
		publishers := models.GetPublishersByCompany(company_name)
		var publishers_ret_temp []Return
		for _, publisher := range *publishers {
			ret := Return{}
			ret.Value = publisher.Userid
			ret.Label = publisher.Userid
			publishers_ret_temp = append(publishers_ret_temp, ret)
		}
		publishersret = append(publishersret, publishers_ret_temp...)
	}

	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: publishersret,
	}
	u.ServeJSON()
}
