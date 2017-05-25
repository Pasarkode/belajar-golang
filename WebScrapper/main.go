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

func ParseNovel(url string) {
	//var dbmap = InitDB()
	res, err := soup.Get(url)

	if err != nil {
		os.Exit(1)
	}

	data := soup.HTMLParse(res)

	getNovel := data.Find("div", "class", "cha-page j_contentWrap")
	getTitle := getNovel.Find("div", "class", "cha-tit").FindAll("h3")
	getContent := getNovel.Find("div", "class", "cha-content").FindAll("p")

	for _, title := range getTitle{
		//for _, content := range getContent{
		chaTit := strings.Split(title.Text(), ": ")
		/*chap := chaTit[0]
		tit := chaTit[1]*/
		//novel := strings.Join(getContent, "")
		fmt.Println(chaTit)
		var temp []string
		for _, content := range getContent {
			temp = append(temp, content.Text())
		}
		fmt.Println("__________________________________________________________________")
		fmt.Print(strings.Join(temp, ""))
		/*novels := Novels{
			Chapter:chap,
			Title:tit,
			Content:,
		}*/

		//err := dbmap.Insert(&novels)
		//checkErr(err, "Gagal memasukkan novel")
		//}
	}

}

func parse(url string) {
	res, err := soup.Get(url)

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(res)

	getList := doc.Find("ul", "class", "g_mod_bd content-list").FindAll("a")

	go func() {
		for _, link := range getList {

			getChapter, err := soup.Get("https:"+link.Attrs()["href"])

			if err != nil {
				os.Exit(1)
			}

			body := soup.HTMLParse(getChapter)

			go func() {
				pages := body.Find("div", "class", "cha-page j_contentWrap")

				titles := pages.Find("div", "class", "cha-tit").FindAll("h3")

				for _, title := range titles{
					fmt.Println(title.Text()+" Link: https:"+link.Attrs()["href"])
				}
			}()

		}
	}()
}

func main() {

	var url,text string
	//url = "https://www.webnovel.com/book/7176992105000305"
	fmt.Println("Masukkan URL dari Webnovel.com")
	fmt.Scanln(&url)
	//go parse(url)
	go ParseNovel(url)
	defer fmt.Println("Scrapping Novel The King's Avatar")
	fmt.Scanln(&text)

}