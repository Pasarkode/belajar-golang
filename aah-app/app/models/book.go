package models

import "time"

type Books struct {
	Id int `json:"id"`
	Judul string `json:"judul"`
	Pengarang string `json:"pengarang"`
	Ringkasan string `json:"ringkasan"`
	Created time.Time `json:"created"`
}
