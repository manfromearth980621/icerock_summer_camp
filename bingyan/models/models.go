package models

import(
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknown/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME		="data/bingyan.db"	//lujing
	_SQLITE3_DRIVER ="sqlite3"	//shujukumingzi
)

/*type Category struct {
	Id  int64
	Title	string
	Create time.Time   `orm: "index"`
	View	int64		`orm: "index"`
	TopicTime time.Time `orm:"index"`
	TopicCount	int64 
	TopicLastUserId int64
}*/

type Topic struct {
	Id		int64
	Title	string
	Content string
	Createtime    time.Time   `orm: "index"`
	TopicTime time.Time `orm:"index"`
}

/*type Message struct {
	Id 			int64
	//Uid 		int64
	Title 		string
	Content 	string     `orm:"size(5000)"`
	Attachment 	string 
	Create	    time.Time   `orm: "index"`
	View		int64		`orm: "index"`
	Author  	string
}*/

type Comment struct {
	comment 	[]string
}

fun AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title: title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_,err := o.Insert(topic)
	return err
}


func RegisterDB(){
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.DIR(_DB_NAME),os.MdelPerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModle(new(Topic),new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) 
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o:= orm.NewOrm()

	topics := make([]*Topic,0)

	qs := o.QueryTable("topic")

	var err error 

	if isDesc {
		_,err=qs.OrderBy("-Created").All($topics)
	}else{
		_,err := qs.All(&topics)
	}
	
	return topics, err
}

