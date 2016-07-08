package routers

import (
	"github.com/astaxie/beego"
	"../controllers"
	"../controllers/admin"
	"../controllers/article"
	"../controllers/comment"
	"../controllers/index"
	"../controllers/project"
)

func init() {
	beego.Router("/", &index.MainController{})
	beego.Router("/page/:page", &index.MainController{})
	beego.Router("/tag/:tag/:page", &index.TagController{})
	beego.Router("/search/:keywords/:page", &index.SearchController{})
	beego.Router("/article/:uri", &article.ArticleController{})
	beego.Router("/article", &article.ArticleController{})
	beego.Router("/article/comment/add/:id", &comment.CommentController{})
	
	beego.Router("/archive/:year/:month/:page", &article.ArchiveController{})
	beego.Router("/list", &article.ArticleListPageController{})
	beego.Router("/list/:page", &article.ArticleListPageController{})
	beego.Router("/project", &project.ProjectListController{}, "*:PageProjects")
	beego.Router("/project/:page", &project.ProjectListController{}, "*:PageProjects")
	beego.Router("/about/statistics", &index.StatisticsController{})
	beego.Router("/about/blog", &index.AboutBlogController{})
	beego.Router("/about/resume", &index.ResumeController{})
	beego.Router("/logo", &index.LogoController{})
	beego.Router("/favicon", &index.SiteIconController{})

	beego.ErrorController(&controllers.ErrorController{})

	model := beego.AppConfig.String("runmode")
	if "dev" == model {
		beego.Router("/test", &admin.TestController{})
	}

}
