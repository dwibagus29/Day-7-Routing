package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("GET")
	route.HandleFunc("/project-detail", projectDetail).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
		// return nil, err
	}

	// w.Write([]byte("addBlog"))
	tmpl.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}
