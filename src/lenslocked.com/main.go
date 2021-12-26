package main

import (
	"net/http"
	
	"lenslocked.com/views"
	"lenslocked.com/controllers"
	
	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w,nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w,nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {	
	homeView = views.NewView("bootstrap","views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")
	
	usersC := controllers.NewUsers()
	
	router := mux.NewRouter()
	
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8501", router)
}