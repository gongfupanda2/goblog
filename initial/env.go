package initial

import (
	"github.com/astaxie/beego"
	//"fmt"
)

func InitEnv() {
	runmode := beego.AppConfig.String("runmode")
	//fmt.Printf("\n"+runmode+"\n\n")
	if runmode == "dev" {
		//fmt.Printf("\n\nstatic\n\n")
		beego.SetStaticPath("/static", "static")
	}
}
