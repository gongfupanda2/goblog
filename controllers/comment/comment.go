package comment

import (
	"fmt"
	"github.com/gogather/com/log"
	"github.com/astaxie/beego/orm"
	"../../controllers"
	. "../../models"
	"strconv"
)

type CommentController struct {
	controllers.BaseController
}

func (this *CommentController) Get() {
	id, err := this.GetInt("id")
	var maps  []orm.Params

	
	if nil == err && id != 0 {
		maps, err = GetComment(int(id))
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJSON()
	}
	this.Data["list"] = maps	
	this.TplName = "listcomment.tpl"
	

	

	

	
}

func (this *CommentController) Post() {
	//Article_id :=  this.Ctx.Input.Param(":id")
	Article_id := this.GetString(":id")
	fmt.Printf("\n\n\n"+Article_id)
	content := this.GetString("comment")
	author := this.GetString("name")
	id_a,_ :=strconv.Atoi(Article_id)
		id, err := AddComment( id_a, content, author)
	if nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "success added, id " + fmt.Sprintf("[%d] ", id), "refer": nil}
	} else {
		log.Warnln(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "added failed", "refer": nil}
	}
	
	this.Redirect("/article/golang", 302)
}
