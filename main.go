package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates./*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

// 	firstName := r.FormValue("first")
// 	lastName := r.FormValue("last")
// 	dateofbirth := r.FormValue("dob")
// 	phonenumber := r.FormValue("phone")
// 	emailid := r.FormValue("email")
// 	cvuplode := r.FormFile("cv")

// d := struct {
// 	FirstName string
// 	LastName  string
// 	Dateofbirth string
// 	Phonenumber string
// 	Emailid string
// 	Cv multipart.File
// }
// {
// 	First: firstName,
// 	Last: lastName,
// 	Dateofbirth: dateofbirth,
// 	Phonenumber: phonenumber,
// 	Emailid: emailid,
// 	Cv: cvuplode,
// }

// tpl.ExecuteTemplate(w, "processer.gohtml")

// }
