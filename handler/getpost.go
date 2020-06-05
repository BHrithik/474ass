package handler

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/BHrithik/474telstra/data"
)

//DisPost for saving the selcted book in the database
func (l *Lib) DisPost(rw http.ResponseWriter, r *http.Request) {
	l.l.Println("Getpost")
	var book ClassifyBookResponse
	var err error

	if book, err = find(r.FormValue("id")); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	result, err := DB.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)",
		nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	pk, _ := result.LastInsertId()
	b := data.Book{
		PK:             int(pk),
		Title:          book.BookData.Title,
		Author:         book.BookData.Author,
		Classification: book.Classification.MostPopular,
	}
	if err = json.NewEncoder(rw).Encode(b); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

func find(id string) (ClassifyBookResponse, error) {
	var c ClassifyBookResponse
	body, err := ClassifyAPI("http://classify.oclc.org/classify2/Classify?summary=true&owi=" + url.QueryEscape(id))

	if err != nil {
		return ClassifyBookResponse{}, err
	}
	err = xml.Unmarshal(body, &c)
	return c, err
}
