package models

import (
	"Gomgo/pkg/setting"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session
var database *mgo.Database

func init() {
	/*配置mongodb的josn文件,配置内容如下:
	 {
	   "hosts": "localhost",
	   "database": "user"
	  }*/

	var err error

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{setting.Host},
		Direct:    false,
		Timeout:   time.Second * 1,
		Username:  setting.User,
		Password:  setting.Password,
		PoolLimit: 4096, // Session.SetPoolLimit
		//创建一个维护套接字池的session
	}
	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Println(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)

	database = GetMgo().DB("chords") } //使用指定数据库

func GetMgo() *mgo.Session {
	return session
}
func GetDataBase() *mgo.Database {
	return database
}
func GetErrNotFound() error {
	return mgo.ErrNotFound
}

//onclick=getChordDetail("'+obj.ID+'")