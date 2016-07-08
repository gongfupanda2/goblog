package article

import (
	"fmt"
	"github.com/astaxie/beego"
	"../../controllers"
	. "../../models"
	"github.com/gogather/com/log"
	"strconv"
)

// 添加文章
type AddArticleController struct {
	controllers.BaseController
}

func (this *AddArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJSON()
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")
	abstract := this.GetString("abstract")

	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	if "" == title {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "title can not be empty", "refer": "/"}
		this.ServeJSON()
		return
	}

	username := user.(string)

	id, err := AddArticle(title, content, keywords, abstract, username)
	if nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "success added, id " + fmt.Sprintf("[%d] ", id), "refer": nil}
	} else {
		log.Warnln(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "added failed", "refer": nil}
	}
	this.ServeJSON()
}

// 获取文章
type ArticleController struct {
	controllers.BaseController
}

func (this *ArticleController) Get() {
	id, err := this.GetInt("id")
	uri := this.Ctx.Input.Param(":uri")

	log.Blueln("[uri]", uri)

	var art Article
	if nil == err && id != 0 {
		art, err = GetArticle(int(id))
	} else if "" != uri {
		art, err = GetArticleByUri(uri)
	} else {
		this.Abort("404")
		this.TplName = "error/404.tpl"
		return
	}

	if 0 == art.Id {
		this.Abort("404")
		this.TplName = "error/404.tpl"
		return
	}

	maps, err := CountByMonth()
	if nil == err {
		this.Data["count_by_month"] = maps
	}

	hottest, err := HottestArticleList()
	if nil == err {
		this.Data["hottest"] = hottest
	}

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
		this.ServeJSON()
	}

	if 0 != art.Id {
		UpdateCount(art.Id)
	}
	

	maps, err = GetComment(int(art.Id))
	/*if nil == err && id != 0 {
		maps, err = GetComment(int(art.Id))
	} else {
		this.Data["commentlist"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJSON()
	} */
	this.Data["id"] = art.Id
	this.Data["title"] = art.Title
	this.Data["uri"] = art.Uri
	this.Data["content"] = art.Content
	this.Data["author"] = art.Author
	this.Data["time"] = art.Time
	this.Data["count"] = art.Count
	this.Data["keywords"] = art.Keywords
	this.Data["description"] = art.Title
	this.Data["duoshuo"] = beego.AppConfig.String("duoshuo_short_name")
	this.Data["commentlist"] = maps
	this.TplName = "article.tpl"
}

func (this *ArticleController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJSON()
}

// 修改文章
type UpdateArticleController struct {
	controllers.BaseController
}

func (this *UpdateArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJSON()
}

func (this *UpdateArticleController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	id, err := this.GetInt64("id")
	uri := this.Ctx.Input.Param(":uri")

	newTitle := this.GetString("title")
	newContent := this.GetString("content")
	newKeywords := this.GetString("keywords")

	if "" == newTitle {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "title can not be empty", "refer": "/"}
		this.ServeJSON()
	}

	var art Article

	if nil == err {
		art, err = GetArticle(int(id))
	} else if "" != uri {
		art, err = GetArticleByUri(uri)
	} else {
		this.Ctx.WriteString("not found")
	}

	art.Title = newTitle
	art.Content = newContent
	art.Keywords = newKeywords

	err = UpdateArticle(id, uri, art)

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "update failed", "refer": "/"}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "update success", "refer": "/"}
		this.ServeJSON()
	}

}

// 删除文章
type DeleteArticleController struct {
	controllers.BaseController
}

func (this *DeleteArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJSON()
}

func (this *DeleteArticleController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	id, err := this.GetInt64("id")
	title := this.Ctx.Input.Param(":title")

	if err != nil {
		id = 0
	}

	num, err := DeleteArticle(id, title)

	if nil != err {
		log.Fatal(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "delete faild", "refer": nil}
		this.ServeJSON()
	} else if 0 == num {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "articles dose not exist", "refer": nil}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": fmt.Sprintf("[%d]", num) + " articles deleted", "refer": nil}
		this.ServeJSON()
	}
}

// 文章列表页
type ArticleListPageController struct {
	controllers.BaseController
}

func (this *ArticleListPageController) Get() {
	s := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPageFlag, _, err := ListPage(int(page), 30)
	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJSON()
	} else {
		this.Data["title"] = "文章列表"
		this.Data["keywords"] = this.Data["title"]
		this.Data["description"] = this.Data["title"]
		this.Data["list"] = maps
		this.Data["prev_page"] = page - 1
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = page + 1
		this.Data["next_page_flag"] = nextPageFlag
		this.TplName = "list.tpl"
	}
}

func (this *ArticleListPageController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request, only get is avalible", "refer": "/"}
	this.ServeJSON()
}

// 按月归档-按月文章列表
type ArchiveController struct {
	controllers.BaseController
}

func (this *ArchiveController) Get() {
	s := this.Ctx.Input.Param(":year")
	year, err := strconv.Atoi(s)
	if nil != err || year < 0 {
		year = 1970
	}

	s = this.Ctx.Input.Param(":month")
	month, err := strconv.Atoi(s)
	if nil != err || month < 0 {
		month = 1
	}

	s = this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPageFlag, pages, err := ListByMonth(year, month, page, 10)

	if pages < int(page) {
		page = pages
	}

	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["prev_page"] = fmt.Sprintf("/archive/%d/%d/%d", year, month, page-1)
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = fmt.Sprintf("/archive/%d/%d/%d", year, month, page+1)
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	hottest, err := HottestArticleList()

	if nil == err {
		this.Data["hottest"] = hottest
	}
	monthMaps, err := CountByMonth()

	if nil == err {
		this.Data["count_by_month"] = monthMaps
	}

	this.Data["title"] = fmt.Sprintf("- %d年%d月", year, month)

	this.TplName = "index.tpl"
}

func (this *ArchiveController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request, only get is avalible", "refer": "/"}
	this.ServeJSON()
}
