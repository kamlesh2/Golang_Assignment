package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/edgedb/edgedb-go"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Resume struct {
	edgedb.Optional
	fname string    `edgedb:"fname"`
	lname string    `edgedb:"lname"`
	phone string    `edgedb:"phone"`
	email string    `edgedb:"email"`
	dob   time.Time `edgedb:"dob"`
	file  time.Time `edgedb:"file"`
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fname := r.FormValue("first")
	last := r.FormValue("last")
	dob := r.FormValue("dob")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	file, handler, err := r.FormFile("cv")
	file_str := ""
	if err != nil {
		log.Fatal(err)
		file := false
		fmt.Printf("%s\n", file)
	} else {
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Print("---------------------------------------------------------------------\n")
		// resp_body, err := ioutil.ReadAll(file)
		// fmt.Printf("%s\n", err)
		// file_str := string(resp_body)
		// fmt.Printf("%s\n", file_str)
		fmt.Print("---------------------------------------------------------------------\n")
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}

	fmt.Printf("%s\n", file)
	fmt.Printf("%s\n", err)

	var inserted struct{ id edgedb.UUID }
	err = client.QuerySingle(ctx, `
	INSERT Resume {
		fname := <str>$0,
		lname := <str>$1,
		phone := <str>$2,
		email := <str>$3,
		dob := <str>$4,
		file := <str>$5,
	}
	`, &inserted, fname, last, phone, email, dob, file_str)

	var resume []Resume
	args := map[string]interface{}{"fname": fname}
	query := "SELECT Resume {fname} FILTER .fname = <str>$fname"
	err = client.Query(ctx, query, &resume, args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resume)

	// Mail
	// from := "from@gmail.com"
	// password := "<Email Password>"
	// to := []string{
	// 	"sender@example.com",
	// }
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"
	// message := []byte("This is a test email message.")
	// auth := smtp.PlainAuth("", from, password, smtpHost)
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Email Sent Successfully!")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
