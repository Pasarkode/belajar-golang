package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/coopernurse/gorp"
	"strings"
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
	checkErr(err, "Create tables failed ")

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
		os.Exit(1)
	}

	doc := soup.HTMLParse(res)

	getList := doc.Find("ul", "class", "g_mod_bd content-list").FindAll("a")

	for _, link := range getList {

		url := "https:"+link.Attrs()["href"]

		go ParseNovel(url)
	}

}

func ParseNovel(url string) {
	var dbmap = InitDB()
	res, err := soup.Get(url)

	if err != nil {
		os.Exit(1)
	}

	data := soup.HTMLParse(res)

	getNovel := data.Find("div", "class", "cha-page j_contentWrap")
	getTitle := getNovel.Find("div", "class", "cha-tit").FindAll("h3")
	getContent := getNovel.Find("div", "class", "cha-content").FindAll("p")

	for _, title := range getTitle{
		chaTit := strings.Split(title.Text(), ": ")
		chap := chaTit[0]
		tit := chaTit[1]
		fmt.Println(tit)
		var temp []string
		for _, content := range getContent {
			temp = append(temp, "<p>"+content.Text()+"</p>")
		}
		novels := Novels{
			Chapter:chap,
			Title:tit,
			Content:strings.Join(temp, ""),
		}

		err := dbmap.Insert(&novels)
		checkErr(err, "Gagal memasukkan novel")
		fmt.Println("Done!")
	}

}

func main() {

	var url,text string
	//url = "https://www.webnovel.com/book/7176992105000305"
	fmt.Print("Masukkan URL halaman TOC dari https://webnovel.com: ")
	fmt.Scanln(&url)
	go parse(url)
	//go ParseNovel(url)
	fmt.Scanln(&text)

}