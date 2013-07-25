package main

import (
	"testing"
	"time"
)

func Test_Benchmark_getAdmin(t *testing.T) {
	loadDB()
	testDB()
}

func testDB() {
	var saveone Softdata
	saveone.Myhardcode = "aa"
	saveone.Software = "assad"
	saveone.Sharecode = "assad"
	saveone.Created = time.Now().Format("2006-01-02 15:04:05")
	saveone.Activecount = 0

	Trace("in set2")

	orm.Save(&saveone)

}
