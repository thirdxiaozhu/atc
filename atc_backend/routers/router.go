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
				"/postmultiatc",
				&controllers.PublisherController{},
				"post:PostMultiAtc",
			),
			beego.NSRouter(
				"/getatc",
				&controllers.PublisherController{},
				"get:GetAtcs",
			),
			beego.NSRouter(
				"/postedit",
				&controllers.PublisherController{},
				"post:PostEdit",
			),
			beego.NSRouter(
				"/postdelete",
				&controllers.PublisherController{},
				"post:PostDelete",
			),
			beego.NSRouter(
				"/getarns",
				&controllers.PublisherController{},
				"get:GetArns",
			),
			beego.NSRouter(
				"/getroutes",
				&controllers.PublisherController{},
				"get:GetRoutes",
			),
			beego.NSRouter(
				"/getauths",
				&controllers.PublisherController{},
				"get:GetAuths",
			),

			beego.NSRouter(
				"/getlinkoptions",
				&controllers.PublisherController{},
				"get:GetLinkOptions",
			),
			beego.NSRouter(
				"/postreglink",
				&controllers.PublisherController{},
				"post:PostRegLink",
			),
			beego.NSRouter(
				"/getreglink",
				&controllers.PublisherController{},
				"get:GetRegLink",
			),
			//beego.NSRouter(
			//	"/getlinks",
			//	&controllers.PublisherController{},
			//	"get:GetLinks",
			//),
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
			//beego.NSRouter(
			//	"/getcompanies",
			//	&controllers.UtilController{},
			//	"get:GetCheckCompanies",
			//),
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
