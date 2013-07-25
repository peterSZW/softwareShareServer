package main

import (
	//	"container/list"
	//	"crypto/rand"
	//	"encoding/hex"
	//	"encoding/json"
	"fmt"
	"html/template"
	//	"io/ioutil"
	"net/http"
	//	"os"
	//	"runtime"
	"strconv"
	"strings"

	//	"sync"
	"time"
)

//var UserDB map[string]*User //get obj from map[string]User will copy struct only
//var ChatRoom map[string]*list.List

var ch chan int

func Handler(w http.ResponseWriter, r *http.Request) {

	//	fmt.Fprintf(w, "<h1>welcome to go chat server %s!</h1>", r.URL.Path[1:])
	//	return

	File := r.URL.Path[1:]
	t, err := template.ParseFiles(File)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
	return

}

// StringReplace -- replaces all occurences of rep with sub in src
func StringReplace(src, rep, sub string) (n string) {
	// make sure the src has the char we want to replace.
	if strings.Count(src, rep) > 0 {
		runes := src // convert to utf-8 runes.
		for i := 0; i < len(runes); i++ {
			l := string(runes[i]) // grab our rune and convert back to string.
			if l == rep {
				n += sub
			} else {
				n += l
			}
		}
		return n
	}
	return src
}

func quit(w http.ResponseWriter, r *http.Request) {
	ch <- 0
	fmt.Fprintf(w, "OK")
}

func count(w http.ResponseWriter, r *http.Request) {
	Userhardcode := r.FormValue("Userhardcode")
	Software := r.FormValue("Software")
	Sharecode := r.FormValue("Sharecode")
	Trace("in count")

	var saveone Softdata
	var check Sharedata

	//查找朋友信息
	orm.Where("Software=? and Sharecode=? and Userhardcode=?", Software, Sharecode, Userhardcode).Find(&check)

	if check.Uid != 0 {
		Trace("如果找到就说已经加过了")
		fmt.Fprintf(w, "already count")
		return
	}
	Trace("找原纪录")

	orm.Where("Software=? and Sharecode=?", Software, Sharecode).Find(&saveone)

	if saveone.Uid != 0 {

		//计数器加一，并保存
		saveone.Activecount++
		orm.Save(&saveone)

		//保存朋友信息，
		check.Userhardcode = Userhardcode
		check.Software = Software
		check.Sharecode = Sharecode
		orm.Save(&check)

		fmt.Fprintf(w, strconv.Itoa(saveone.Activecount))
	} else {
		fmt.Fprintf(w, Userhardcode)
	}

}

func get(w http.ResponseWriter, r *http.Request) {
	Myhardcode := r.FormValue("Myhardcode")
	Software := r.FormValue("Software")
	Sharecode := r.FormValue("Sharecode")
	Trace("in set")

	var saveone Softdata

	orm.Where("Myhardcode=? and Software=? and Sharecode=?",
		Myhardcode, Software, Sharecode).Find(&saveone)

	if saveone.Uid == 0 {

		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, strconv.Itoa(saveone.Activecount))
	}

}
func set(w http.ResponseWriter, r *http.Request) {
	Myhardcode := r.FormValue("Myhardcode")
	Software := r.FormValue("Software")
	Sharecode := r.FormValue("Sharecode")
	Trace("in set")

	var saveone Softdata

	orm.Where("Myhardcode=? and Software=? and Sharecode=?",
		Myhardcode, Software, Sharecode).Find(&saveone)

	if saveone.Uid == 0 {

		saveone.Myhardcode = Myhardcode
		saveone.Software = Software
		saveone.Sharecode = Sharecode
		saveone.Created = time.Now().Format("2006-01-02 15:04:05")
		saveone.Activecount = 0

		Trace("Save New")

		orm.Save(&saveone)
		fmt.Fprintf(w, "OK")
	} else {
		fmt.Fprintf(w, "Exists")
	}

	//msg = StringReplace(msg, " ", "+")

}

func main() {
	//UserDB = make(map[string]*User)
	//ChatRoom = make(map[string]*list.List)

	SetLevel(0)
	//readOptions()
	loadDB()

	ch = make(chan int, 10)

	http.Handle("/html/", http.FileServer(http.Dir("")))

	http.HandleFunc("/", Handler)
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/count", count)
	http.HandleFunc("/quit", quit)

	Trace("Running")
	http.ListenAndServe(":19999", nil)

	<-ch
	Trace("Exited")
}

/*
127.0.0.1:19999/set?Myhardcode=123&Software=TextFinding&Sharecode=12
127.0.0.1:19999/get?Myhardcode=123&Software=TextFinding&Sharecode=12
127.0.0.1:19999/count?Userhardcode=456&Software=TextFinding&Sharecode=12

*/
