package handler

import (
	"net/http"

	"github.com/BHrithik/474telstra/data"
)

//GetPosts Initial landing page
func (l *Lib) GetPosts(rw http.ResponseWriter, r *http.Request) {
	l.l.Println("Post handler")
	p := data.Page{[]data.Book{}}
	rows, _ := DB.Query("select pk,title,author,classification from books")
	for rows.Next() {
		var b data.Book
		rows.Scan(&b.PK, &b.Title, &b.Author, &b.Classification)
		p.Books = append(p.Books, b)
	}
	tpl := Temp(rw)
	if err := tpl.Execute(rw, p); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
