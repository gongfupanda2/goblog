package models
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	
	// "github.com/gogather/com"

	"time"
)

type Comment struct {
	Id       int
	Article_id int
	Content  string
	Author   string
	Time     time.Time
	
}

func (this *Comment) TableName() string {
	return "Comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}

func AddComment(Article_id int, content string, author string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")

	sql := "insert into comment(article_id, content, author) values(?, ?, ?)"
	res, err := o.Raw(sql, Article_id,  content, author).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

// 通过id获取comment-cached
func GetComment(Article_id int) ([]orm.Params,  error) {
	var err error
	
	sql1 := "select * from comment where article_id = ? order by time desc "
	var maps  []orm.Params
	o := orm.NewOrm()
	_, err = o.Raw(sql1, fmt.Sprintf("%d", Article_id), ).Values(&maps)
	
	if err == nil {
		return maps,  nil
	} else {
		return nil,  err
	}
}
