package handler

import "net/http"

// DelBook deletes the clicked book
func (l *Lib) DelBook(rw http.ResponseWriter, r *http.Request) {

	l.l.Println("Delete Handler")
	if _, err := DB.Exec("delete from books where pk=?", r.FormValue("pk")); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.WriteHeader(http.StatusOK)
}
