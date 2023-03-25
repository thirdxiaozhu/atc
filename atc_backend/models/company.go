package models

import "github.com/beego/beego/v2/core/logs"

type Company struct {
	ID     uint   `orm:"column(id);auto"`
	Name   string `orm:"column(name);unique;size(20)"`
	CNname string `orm:"column(cnname);size(20)"`
	Role   int    `orm:"column(role)"`
}

func init() {
	logger := logs.GetLogger()
	logger.Println(getCompanies())
}

func getCompanies() []Company {
	companies := []Company{
		{
			Name:   "eastchina",
			CNname: "华东空管局",
			Role:   0,
		},
		{
			Name:   "northchina",
			CNname: "华北空管局",
			Role:   0,
		},
		{
			Name:   "airchina",
			CNname: "中国国际航空",
			Role:   1,
		},
		{
			Name:   "csair",
			CNname: "中国南方航空",
			Role:   1,
		},
	}

	return companies
}
