package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"service"

	"github.com/google/uuid"
)

type User struct {
	ID       uint   `orm:"column(id);auto"`
	Userid   string `orm:"column(name);unique;size(20)"`
	Password string `orm:"column(password);size(256)"`
	Company  string `orm:"column(company);size(20)"`
	UUID     string `orm:"column(uuid);size(64)"`
}

type Login_Return struct {
	User    User    `json:"user"`
	Company Company `json:"company"`
	//CNcompany string  `json:"cncompany"`
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
		"eastchina":  "/root/go/src/atc/config/config_eastchina.yaml",
		"northchina": "/root/go/src/atc/config/config_northchina.yaml",
		"airchina":   "/root/go/src/atc/config/config_airchina.yaml",
		"csair":      "/root/go/src/atc/config/config_csair.yaml"}
	setup_map = make(map[string]*service.ServiceSetup)
}

func Login(userid, password string) *Login_Return {
	//users := []User{}
	user := User{}
	err := DBH.QueryOneByField(&user, "user", "name", userid)
	if err == nil {
		h := md5.New()
		io.WriteString(h, password)
		if user.Password != fmt.Sprintf("%x", h.Sum(nil)) {
			return nil
		}
		/* 初始化服务 */
		setup := service.InitSetup(configpath_map[user.Company], user.Company)
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
			User:    user,
			Company: company,
		}

		return &login_ret
	}
	return nil
}

func Signup(userid, password, company string) bool {
	//var uu User

	h := md5.New()
	io.WriteString(h, password)

	user := User{
		Userid:   userid,
		Password: fmt.Sprintf("%x", h.Sum(nil)),
		Company:  company,
	}
	logger.Println(user)
	res := DBH.Insert(&user)
	return res
}

func DeleteUser(uid string) {
	return
}
