package controllers

import (
	"atc_backend/models"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	//var us models.UserSearch
	userid := u.GetString("id")

	user := models.User{}
	id, _ := strconv.Atoi(userid)
	user.ID = uint(id)
	models.DBH.QueryByPK(&user)
	u.Data["json"] = user.Userid

	u.ServeJSON()

}

func (u *UserController) TestQuery() {

	//var us models.UserSearch
	userid := u.GetString("id")

	user := models.User{}
	id, _ := strconv.Atoi(userid)
	user.ID = uint(id)
	models.DBH.QueryByPK(&user)
	u.Data["json"] = user.Userid

	u.ServeJSON()
}

func (u *UserController) TestCreate() {
	user := models.User{}
	user.Userid = u.GetString("username")
	user.Password = u.GetString("password")
	models.DBH.Insert(&user)
	u.Data["json"] = "succeed to add a user"

	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	//user := models.User{
	//	Userid:   u.GetString("userid"),
	//	Password: u.GetString("userid"),
	//}
	logger.Println(u.GetString("userid"))
	log_ret := models.Login(u.GetString("userid"), u.GetString("password"))
	u.Data["json"] = JsonResponse{
		Code: 1000,
		Data: *log_ret,
	}
	u.ServeJSON()
}

func (u *UserController) Signup() {
	logger := logs.GetLogger()
	logger.Println(u.GetString("userid"), u.GetString("password"), u.GetString("company"))
	user := models.User{
		Userid:   u.GetString("userid"),
		Password: u.GetString("password"),
		Company:  u.GetString("company"),
	}
	var res JsonResponse
	result := models.Signup(&user)
	if result == false {
		res.Code = 1001
		res.Data = "该用户已被注册"
	} else {
		res.Code = 1000
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
