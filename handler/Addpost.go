package handler

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/BHrithik/474telstra/data"
)

//AddPost for displaying the search results
func (l *Lib) AddPost(rw http.ResponseWriter, r *http.Request) {

	l.l.Println("ADD HANDLER")
	var results data.Search
	var err error

	if results, err = search(r.FormValue("search")); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	if err := results.ToJSON(rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func search(query string) (data.Search, error) {
	var c ClassifySearchResponse
	body, err := ClassifyAPI("http://classify.oclc.org/classify2/Classify?summary=true&title=" + url.QueryEscape(query))

	if err != nil {
		return data.Search{}, err
	}

	err = xml.Unmarshal(body, &c)
	return c.Results, err
}
