package routers

import (
	"github.com/astaxie/beego"
	"../controllers/api"
	"../controllers/fis"
	"../controllers/xmlrpc"
)

func init() {
	beego.Router("/api/get/user", &api.CurrentUserController{})
	beego.Router("/api/get/total_article_number", &api.TotalArticleNumberController{})
	beego.Router("/api/get/total_user_number", &api.TotalUserNumberController{})
	beego.Router("/api/get/server_time", &api.ServerTimeController{})
	beego.Router("/map.json", &api.MapJsonController{})

	beego.Router("/xmlrpc", &xmlrpc.XmlrpcController{})

	beego.Router("/fis", &fis.FisController{}, "*:Receiver")

}
