package data

import (
	"encoding/json"
	"io"

	//Driver for sql
	_ "github.com/mattn/go-sqlite3"
)

//Book for getting the data from the db
type Book struct {
	PK             int
	Title          string
	Author         string
	Classification string
}

//Page sample struct
type Page struct {
	Books []Book
}

//SearchResults used to store the searched value and display it on the html
type SearchResults struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

//Search for storing []Searchresults
type Search []SearchResults

//ToJSON write to the html file
func (s *Search) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}
