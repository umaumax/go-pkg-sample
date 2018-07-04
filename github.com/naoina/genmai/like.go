package main

import (
	"fmt"
	"log"
	"os"

	//	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

//	_ "github.com/mattn/go-sqlite3"
//	go installを用いるべし
/*
  147782  687547 5239373 sqlite3.c
    7478   53408  360297 sqlite3.h
     487    1567   26110 sqlite3ext.h
*/

type Keyword struct {
	Id  int    `db:"pk" column:"id"`
	Key string `column:"key"`
}

func main() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.CreateTable(&Keyword{}); err != nil {
		log.Fatalln(err)
	}

	db.SetLogOutput(os.Stdout)

	//	["]の処理を行う
	//	objectに対するメソッでであるので、SQL文構築時に内部で用いられているかも...
	//	むやみに使用しないほうが良いかも
	//	fmt.Println(db.Quote("%go\"og'le%"))

	tbl := Keyword{
		Key: "google",
	}
	_, err = db.Insert(&tbl)
	if err != nil {
		log.Fatalln(err)
	}

	//	NOTE Whereではsliceを用いなければならない
	var tbls []Keyword
	//	query := db.Where("key").Like("\" OR \"1\" = \"1\"")
	query := db.Where("key").Like("go%le")
	err = db.Select(&tbls, query)
	//	fmt.Printf("%+v\n", query)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tbls)
}
