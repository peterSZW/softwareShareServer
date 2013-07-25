package main

import (
	"database/sql"
	//"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	//"time"
)

/*

1: …../SetShareCode?MyHardCode= ….&Software=…..&ShareCode=……//返回设置是否成功
2:…../ActiveShareCode?UserHardCode=….&Software=…..&ShareCode=……
3:…../ActiveSoftware?MyHardCode=….&Software=…..&ShareCode=……//返回已提交的用户数

CREATE TABLE `softdata` (
`uid` INTEGER PRIMARY KEY AUTOINCREMENT,
	`myhardcode` VARCHAR(64) NULL,
	`software` VARCHAR(64) NULL,
	`sharecode` VARCHAR(64) NULL,
        `activecount` INTEGER ,
	`created` DATE NULL
);
CREATE TABLE `sharedata` (
	`uid` INTEGER PRIMARY KEY AUTOINCREMENT,
        `userhardcode` VARCHAR(64) NULL,
	`software` VARCHAR(64) NULL,
	`sharecode` VARCHAR(64) NULL,
	`created` DATE NULL
);
*/

var orm beedb.Model

type Softdata struct {
	Uid         int `beedb:"PK"`
	Myhardcode  string
	Software    string
	Sharecode   string
	Activecount int
	Created     string
}

type Sharedata struct {
	Uid          int `beedb:"PK"`
	Userhardcode string
	Software     string
	Sharecode    string
	Created      string
}

func loadDB() {
	db, err := sql.Open("sqlite3", "./softwareserver.s3db")
	if err != nil {
		panic(err)
	}
	orm = beedb.New(db)

	/*
		insert()
		insertsql()
		a := selectone()
		fmt.Println(a)

		b := selectall()
		fmt.Println(b)

		update()

		updatesql()

		findmap()

		groupby()

		jointable()

		deleteone()

		deletesql()

		deleteall()*/
}

/*
func insert() {
	//save data
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now().Format("2006-01-02 15:04:05")
	orm.Save(&saveone)
	fmt.Println(saveone)
}

func insertsql() {
	// add one
	add := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2012-12-02"
	orm.SetTable("userinfo").Insert(add)
}

func selectone() Userinfo {
	//get one info
	var one Userinfo
	orm.Where("uid=?", 1).Find(&one)
	return one
}

func selectall() []Userinfo {
	//get all data
	var alluser []Userinfo
	orm.Limit(10).Where("uid>?", 1).FindAll(&alluser)
	return alluser
}

func update() {
	// //update data
	var saveone Userinfo
	saveone.Uid = 1
	saveone.Username = "Update Username"
	saveone.Departname = "Update Departname"
	saveone.Created = time.Now().Format("2006-01-02 15:04:05")
	orm.Save(&saveone)
	fmt.Println(saveone)
}

func updatesql() {
	//original SQL update
	t := make(map[string]interface{})
	t["username"] = "updateastaxie"
	//update one
	orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)
	//update batch
	orm.SetTable("userinfo").Where("uid>?", 3).Update(t)
}

func findmap() {
	//Original SQL Backinfo resultsSlice []map[string][]byte
	//default PrimaryKey id
	c, _ := orm.SetTable("userinfo").SetPK("uid").Where(2).Select("uid,username").FindMap()
	fmt.Println(c)
}

func groupby() {
	//Original SQL Group By
	b, _ := orm.SetTable("userinfo").GroupBy("username").Having("username='updateastaxie'").FindMap()
	fmt.Println(b)
}

func jointable() {
	//Original SQL Join Table
	a, _ := orm.SetTable("userinfo").Join("LEFT", "userdeatail", "userinfo.uid=userdeatail.uid").Where("userinfo.uid=?", 1).Select("userinfo.uid,userinfo.username,userdeatail.profile").FindMap()
	fmt.Println(a)
}

func deleteone() {
	// // //delete one data
	saveone := selectone()
	orm.Delete(&saveone)
}

func deletesql() {
	//original SQL delete
	orm.SetTable("userinfo").Where("uid>?", 2).DeleteRow()
}

func deleteall() {
	// //delete all data
	alluser := selectall()
	orm.DeleteAll(&alluser)
}
*/
