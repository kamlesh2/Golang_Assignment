package main

import (
	"html/template"
	"net/http"
	"fmt"
)

import "github.com/gofiber/fiber/v2"

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/edgedb/edgedb-go"
// )

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// func main() {
// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/process", processor)
// 	http.ListenAndServe(":8080", nil)
// }


func main() {
    app := fiber.New()

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.SendString("Hello, World 👋!")
    // })
	
	// DB

	// ctx := context.Background()
    // client, err := edgedb.CreateClient(ctx, edgedb.Options{})
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer client.Close()

    // var (
    //     users []struct {
    //         ID   edgedb.UUID `edgedb:"id"`
    //         FName string      `edgedb:"fname"`
    //         LName string      `edgedb:"lname"`
    //     }
    // )

    // query := "SELECT User{name} FILTER .age = <int64>$0"
    // err = client.Query(ctx, query, &users, age)

	// DB

	app.Static("/", "./public") 
	
	app.Post("/form/submit", func(c *fiber.Ctx) error {
		fname := c.FormValue("first")
		last := c.FormValue("last")
		dob := c.FormValue("dob")
		Phone := c.FormValue("Phone")
		email := c.FormValue("email")
		file, err := c.FormFile("cv")
		fmt.Printf("%s\n",fname)
		fmt.Printf("%s\n",last)
		fmt.Printf("%s\n",dob)
		fmt.Printf("%s\n",Phone)
		fmt.Printf("%s\n",email)
		fmt.Printf("%s\n",file)
		fmt.Printf("%s\n",err)
		return c.SendString("I'm a POST request!")
	})

    app.Listen(":3000")
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
