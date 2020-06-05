package handler

import (
	"net/http"
)

// VerifyDatabase verifies database
func (l Lib) VerifyDatabase(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// db, _ := sql.Open("sqlite3", "dev.db")

		if err := DB.Ping(); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			l.l.Printf("Error connecting to database %v", err.Error())
			return
		}
		next.ServeHTTP(rw, r)
	})
}

//VerifyDatabase verifies database
// func VerifyDatabase(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

// 	if err := DB.Ping(); err != nil {
// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	next.ServeHTTP(rw, r)
// }
