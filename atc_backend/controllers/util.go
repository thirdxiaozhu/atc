package controllers

import (
	"atc_backend/models"
	"log"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type CompanyRet struct {
	Value uint   `json:"value"`
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
	companies := []models.Company{}
	var companiesret []CompanyRet
	models.DBH.QueryAllByField(&companies, "company", "role", u.GetString("role"))

	for _, company := range companies {
		ret := CompanyRet{}
		ret.Value = company.ID
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
