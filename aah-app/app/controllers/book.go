package controllers

import (
	"aah-app/app/models"

	aah "aahframework.org/aah.v0"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type Book struct {
	*aah.Context
}

type Hasil map[string]interface{}

var dbmap = InitDB("books")

func (a *Book) Index() {
	var books []models.Books
	_, err := dbmap.Select(&books, "SELECT * FROM books ORDER BY id")
	checkErr(err, "Select data fail ")

	content := Hasil{}
	for k, v := range books {
		content[strconv.Itoa(k)] = v
	}

	a.Reply().Ok().JSON(content)
}

func (a *Book) Create() {
	a.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	Judul := a.Req.FormValue("judul")
	Pengarang := a.Req.FormValue("pengarang")
	Ringkasan := a.Req.FormValue("ringkasan")
	Created := time.Now()

	book := Store(Judul, Pengarang, Ringkasan)
	if book.Judul == Judul {
		a.Reply().Ok().JSON(models.Books{
			Judul:Judul,
			Pengarang:Pengarang,
			Ringkasan:Ringkasan,
			Created:Created,
		})
	} else {
		a.Reply().InternalServerError().JSON("Gagal Input Buku")
	}
}

func Store(Judul, Pengarang, Ringkasan string) models.Books {
	book := models.Books{
		Judul:Judul,
		Pengarang:Pengarang,
		Ringkasan:Ringkasan,
		Created:time.Now(),
	}
	err := dbmap.Insert(&book)
	checkErr(err, "Insert failed")

	return book
}

func (a *Book) GetBook() {
	bookId, err := strconv.Atoi(a.Req.PathValue("id"))
	if err != nil {
		panic("Gagal...")
	}
	book := OneBook(bookId)
	a.Reply().Ok().JSON(models.Books{
		Id:book.Id,
		Judul:book.Judul,
		Pengarang:book.Pengarang,
		Ringkasan:book.Ringkasan,
		Created:book.Created,
	})
}

func OneBook(bookId int) models.Books {
	book := models.Books{}
	err := dbmap.SelectOne(&book, "SELECT * FROM books WHERE id=?", bookId)
	checkErr(err, "Select One Failed ")
	return book
}

func (a *Book) Update() {
	bookId, err := strconv.Atoi(a.Req.PathValue("id"))
	Judul := a.Req.FormValue("judul")
	Pengarang := a.Req.FormValue("pengarang")
	Ringkasan := a.Req.FormValue("ringkasan")

	if err != nil {
		panic("Gagal...")
	}

	book := Patch(Judul, Pengarang, Ringkasan, bookId)
	checkErr(err, "Gagal Update data")
	a.Reply().Ok().JSON(models.Books{
		Id:book.Id,
		Judul:book.Judul,
		Pengarang:book.Pengarang,
		Ringkasan:book.Ringkasan,
		Created:book.Created,
	})

}

func Patch(Judul, Pengarang, Ringkasan string, bookId int) models.Books {
	book := models.Books{
		Judul:Judul,
		Pengarang:Pengarang,
		Ringkasan:Ringkasan,
		Id:bookId,
	}
	//query, err := dbmap.Prepare("UPDATE books SET judul = ? , pengarang = ? , ringkasan = ? WHERE id = ?")
	_, err := dbmap.Update(&book)
	checkErr(err, "Gagal update buku... ")
	bookOne := OneBook(bookId)

	return bookOne
}

func (a *Book) Delete() {
	bookId, err := strconv.Atoi(a.Req.PathValue("id"))
	cari := OneBook(bookId)
	data:= models.Books{
		Id:cari.Id,
		Judul:cari.Judul,
		Pengarang:cari.Pengarang,
		Ringkasan:cari.Ringkasan,
	}
	checkErr(err, "Gagal ambil data untuk dihapus... ")
	hasil, err := dbmap.Delete(&data)
	if hasil != 1 {
		a.Reply().Ok().Text("Gagal hapus data")
	}

	a.Reply().Ok().Text("Berhasil hapus data")
}