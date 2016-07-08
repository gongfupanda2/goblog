package main

import (
	"github.com/astaxie/beego"
	_ "./initial"
	_ "./routers"
)

func main() {
	
	beego.Run()
}
