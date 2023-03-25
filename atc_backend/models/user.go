package models

import (
	"service"

	"github.com/google/uuid"
)

type User struct {
	ID       uint   `orm:"column(id);auto"`
	Userid   string `orm:"column(name);unique;size(20)"`
	Password string `orm:"column(password);size(16)"`
	Company  int    `orm:"column(company);size(20)"`
	UUID     string `orm:"column(uuid);size(64)"`
}

var (
	cryptpath_map  map[int]string
	configpath_map map[int]string
	setup_map      map[string]*service.ServiceSetup
)

func init() {
	cryptpath_map = map[int]string{
		1: "/root/go/src/atc/fabric/crypto-config/peerOrganizations/eastchina.atc.com/users/User1@eastchina.atc.com/tls/",
		2: "/root/go/src/atc/fabric/crypto-config/peerOrganizations/northchina.atc.com/users/User1@northchina.atc.com/tls/",
		3: "/root/go/src/atc/fabric/crypto-config/peerOrganizations/airchina.atc.com/users/User1@airchina.atc.com/tls/",
		4: "/root/go/src/atc/fabric/crypto-config/peerOrganizations/csair.atc.com/users/User1@csair.atc.com/tls/",
	}
	configpath_map = map[int]string{
		1: "../../config/config_eastchina.yaml",
		2: "../../config/config_northchina.yaml",
		3: "../../config/config_airchina.yaml",
		4: "../../config/config_csair.yaml"}
	setup_map = make(map[string]*service.ServiceSetup)
}

func Login(userid, password string) *User {
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

		return &user
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
