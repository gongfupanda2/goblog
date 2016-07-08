package main

import (
    "github.com/astaxie/beego"
    "fmt"
  
)

type MainController struct {
    beego.Controller
}


func (index *MainController) Get() {
	 sess := index.StartSession()
    sess.Set("username","asd")
    username := sess.Get("username")
    fmt.Println(username)
    if username == nil || username == "" {
        index.Ctx.WriteString("hello world")
    } else {
         index.Ctx.WriteString("hello world123")
    }

    
}

func main() {
	
	beego.Router("/", &MainController{})
    beego.Run()

}