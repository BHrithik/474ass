package handler

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yosssi/ace"

	"github.com/BHrithik/474telstra/data"
)

//DB database
var (
	DB, _ = sql.Open("sqlite3", "dev.db")
)

//Lib is handler
type Lib struct {
	l *log.Logger
}

//NewlibraryHandler creates a handler
func NewlibraryHandler(l *log.Logger) *Lib {
	return &Lib{l}
}

//Temp is temp
func Temp(rw http.ResponseWriter) *template.Template {
	tpl, err := ace.Load("templates/index", "", nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	return tpl
}

//ClassifyBookResponse for storing the book response from the id passed as form value
type ClassifyBookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		ID     string `xml:"owi,attr"`
	} `xml:"work"`
	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendations>ddc>mostPopular"`
}

//ClassifySearchResponse for displaying the search results
type ClassifySearchResponse struct {
	Results data.Search `xml:"works>work"`
}

//ClassifyAPI searches for the url reads and sends the byte[]
func ClassifyAPI(url string) ([]byte, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
