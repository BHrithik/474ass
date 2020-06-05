package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/BHrithik/474ass/handler"
	"github.com/gorilla/mux"
)

//DB database

func main() {

	l := log.New(os.Stdout, "474_508_OnlineLib", log.LstdFlags)
	lh := handler.NewlibraryHandler(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", lh.GetPosts)
	addRouter := sm.Methods("POST").Subrouter()
	addRouter.HandleFunc("/search", lh.AddPost)
	disRouter := sm.Methods("PUT").Subrouter()
	disRouter.HandleFunc("/books/add", lh.DisPost)
	disRouter.Use(lh.VerifyDatabase)
	delRouter := sm.Methods("DELETE").Subrouter()
	delRouter.HandleFunc("/books/delete", lh.DelBook)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		ErrorLog:     l,
	}

	// n := negroni.Classic()
	// n.Use(negroni.HandlerFunc(handler.VerifyDatabase))
	// n.UseHandler(sm)
	go func() {
		l.Println("Starting server on port :9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
		// n.Run(":9090")
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println(sig)
	l.Println("Got Signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	l.Println("Shutting down")
	os.Exit(1)

}
