package mongo

import (
	"fmt"
	"gin-app-start/app/config"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Session
var Db string

// 初始化
func Init() {
	config := config.Conf.Mongo
	s, err := mgo.Dial(config.Url)
	if err != nil {
		fmt.Println("mongo connection error: ", err)
		panic(err.Error())
	}
	s.SetMode(mgo.Monotonic, true)
	fmt.Println("mongo connection open to: ", config.Url)
	Session = s
	Db = config.Database
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	// 每次操作copy一份 Session,避免每次创建Session
	ms := Session.Copy()
	c := ms.DB(db).C(collection)
	return ms, c
}

/**
 * 插入数据
 * db:操作的数据库
 * collection:操作的文档(表)
 * doc:要插入的数据
 */
func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(doc)
}

/**
 * 查询数据
 * db:操作的数据库
 * collection:操作的文档(表)
 * query:查询条件
 * selector:需要过滤的数据(projection)
 * result:查询到的结果
 */
func FindOne(db, collection string, query, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	// return c.Find(query).Select(selector).One(result)
	return c.Find(query).One(result)
}

// 查询列表
func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
	// return c.Find(query).All(result)
}

/**
 * 分页查询
 * db:操作的数据库
 * collection:操作的文档(表)
 * page:当前页面
 * limit:每页的数量值
 * query:查询条件
 * selector:需要过滤的数据(projection)
 * result:查询到的结果
 */
func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

/**
 * 更新数据
 * db:操作的数据库
 * collection:操作的文档(表)
 * selector:更新条件
 * update:更新的操作
 */
func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

//更新，如果不存在就插入一个新的数据 `upsert:true`
func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

// `multi:true`
func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

/**
 * 删除数据
 * db: 操作的数据库
 * collection: 操作的文档(表)
 * selector: 删除条件
 */
func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

// 数量统计
func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Count()
}

// 是否存在
func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}
