// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"atc_backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/publisher",
			beego.NSRouter(
				"/postatc",
				&controllers.PublisherController{},
				"post:PostAtc",
			),
			beego.NSRouter(
				"/getatc",
				&controllers.PublisherController{},
				"get:GetAtcs",
			),
		),

		beego.NSNamespace("/sundries"),

		beego.NSNamespace("/user",
			//beego.NSInclude(
			//	&controllers.UserController{},
			//),
			beego.NSRouter(
				"/search",
				&controllers.UserController{},
				"get:TestQuery",
			),
			beego.NSRouter(
				"/create",
				&controllers.UserController{},
				"post:TestCreate",
			),
			beego.NSRouter(
				"/signup",
				&controllers.UserController{},
				"post:Signup",
			),
			beego.NSRouter(
				"/login",
				&controllers.UserController{},
				"post:Login",
			),
		),
		//beego.NSRouter("/login", &controllers.UserController{}, "post: Login"),
		beego.NSNamespace("/util",
			beego.NSRouter(
				"/company",
				&controllers.UtilController{},
				"get:GetCompanies",
			),
			beego.NSRouter(
				"/getcheckcompanies",
				&controllers.UtilController{},
				"get:GetCheckCompanies",
			),
			beego.NSRouter(
				"/getcheckpublishers",
				&controllers.UtilController{},
				"get:GetCheckPublishers",
			),
			beego.NSRouter(
				"/getcert",
				&controllers.UtilController{},
				"get:GetCertificate",
			),
			beego.NSRouter(
				"/getpriv",
				&controllers.UtilController{},
				"get:GetPrivateKey",
			),
		),
	)
	beego.AddNamespace(ns)
}
