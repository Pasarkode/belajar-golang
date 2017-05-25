package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/coopernurse/gorp"
	"strings"
	"os"
)

type Novels struct {
	Chapter string `json:"chapter"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func InitDB() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "database.sqlite3")
	checkErr(err, "Database Open failed ")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Novels{}, "novels")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed... ")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func parse(url string) {
	res, err := soup.Get(url)

	if err != nil {
		panic(err)
	}

	doc := soup.HTMLParse(res)

	getList := doc.Find("ul", "class", "g_mod_bd content-list").FindAll("a")

	for _, link := range getList {

		url := "https:"+link.Attrs()["href"]

		ParseNovel(url)
	}

}

func ParseNovel(url string) {
	var dbmap = InitDB()
	res, err := soup.Get(url)

	if err != nil {
		os.Exit(1)
	}

	data := soup.HTMLParse(res)

	getTitle := data.Find("div", "class", "cha-tit").Find("h3")
	getContent := data.Find("div", "class", "cha-content").FindAll("p")

	chaTit := strings.Split(getTitle.Text(), ": ")
	chap := chaTit[0]
	tit := chaTit[1]

	var temp []string

	fmt.Println(chap+": "+tit)


	for _, content := range getContent {

		if content.Text() != "" {
			temp = append(temp, "<p>"+content.Text()+"</p>")
		}

	}

	novels := Novels{
		Chapter:chap,
		Title:tit,
		Content:strings.Join(temp, ""),
	}

	err = dbmap.Insert(&novels)
	checkErr(err, "Gagal memasukkan novel")

}

func main() {

	var url,text string
	fmt.Print("Masukkan URL halaman TOC dari https://webnovel.com: ")
	fmt.Scanln(&url)
	go parse(url)
	fmt.Println("Done!")
	fmt.Scanln(&text)

}