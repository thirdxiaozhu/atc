package models

import (
	"service"

	"github.com/google/uuid"
)

type User struct {
	ID       uint   `orm:"column(id);auto"`
	Userid   string `orm:"column(name);unique;size(20)"`
	Password string `orm:"column(password);size(16)"`
	Company  string `orm:"column(company);size(20)"`
	UUID     string `orm:"column(uuid);size(64)"`
}

type Login_Return struct {
	User      User   `json:"user"`
	CNcompany string `json:"cncompany"`
}

var (
	cryptpath_map  map[string]string
	configpath_map map[string]string
	setup_map      map[string]*service.ServiceSetup
)

func init() {
	cryptpath_map = map[string]string{
		"eastchina":  "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/",
		"northchina": "/root/go/src/atc/fabric/crypto-config/peerOrganizations/northchina.atc.com/users/User1@northchina.atc.com/tls/",
		"airchina":   "/root/go/src/atc/fabric/crypto-config/peerOrganizations/airchina.atc.com/users/User1@airchina.atc.com/tls/",
		"csair":      "/root/go/src/atc/fabric/crypto-config/peerOrganizations/csair.atc.com/users/User1@csair.atc.com/tls/",
	}
	configpath_map = map[string]string{
		"eastchina":  "../../config/config_eastchina.yaml",
		"northchina": "../../config/config_northchina.yaml",
		"airchina":   "../../config/config_airchina.yaml",
		"csair":      "../../config/config_csair.yaml"}
	setup_map = make(map[string]*service.ServiceSetup)
}

func Login(userid, password string) *Login_Return {
	//users := []User{}
	user := User{}
	err := DBH.QueryOneByField(&user, "user", "name", userid)
	if err == nil {
		if user.Password != password {
			return nil
		}
		/* 初始化服务 */
		setup := service.InitSetup(configpath_map[user.Company])
		if setup == nil {
			return nil
		}
		setup_map[user.Userid] = setup

		//生成UUID，作为token
		u, _ := uuid.NewRandom()
		user.UUID = u.String()
		DBH.Update(&user)

		company := Company{}
		err := DBH.QueryOneByField(&company, "company", "name", user.Company)
		if err != nil {
			return nil
		}

		login_ret := Login_Return{
			User:      user,
			CNcompany: company.CNname,
		}

		return &login_ret
	}
	return nil
}

func Signup(u *User) bool {
	//var uu User
	res := DBH.Insert(u)
	return res
}

func DeleteUser(uid string) {
	return
}
