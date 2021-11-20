package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `email:"email"`
}

func main() {
	http.HandleFunc("/", index_page)
	http.HandleFunc("/hotels/", hotels_page)
	http.HandleFunc("/contacts/", contact_page)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index_page(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 0, FirstName: "Cassiano", LastName: "Yasumitsu", Email: "cassiano@email.com"}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, user)
}

func hotels_page(w http.ResponseWriter, r *http.Request) {

}

func contact_page(w http.ResponseWriter, r *http.Request) {

}
