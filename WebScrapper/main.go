package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {

	fmt.Println("Scrapping URL Novel China The King's Avatar")

	res, err := soup.Get("https://www.webnovel.com/book/7176992105000305")

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(res)

	getList := doc.Find("ul", "class", "g_mod_bd content-list").FindAll("a")

	for _, link := range getList {

		getChapter, err := soup.Get("https:"+link.Attrs()["href"])

		if err != nil {
			os.Exit(1)
		}

		body := soup.HTMLParse(getChapter)

		pages := body.Find("div", "class", "cha-page j_contentWrap")

		titles := pages.Find("div", "class", "cha-tit").FindAll("h3")

		for _, title := range titles{
			fmt.Println("Chapter Title: "+title.Text()+" Link: https:"+link.Attrs()["href"])
		}

	}
}