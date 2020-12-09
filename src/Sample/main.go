package main

import (
	"SampleGo/src/Sample/controller"
	"SampleGo/src/Sample/model"
	"os"

	"github.com/withmandala/go-log"

	"database/sql"
	"html/template"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
)

const passwordSalt = "a99VVoWzmd1C9ujcitK0fIVNE0I5I61AC47C852RoLTsHDyLCltvP+ZHEkIl/2hkzTOW90c3ZEjtYRkdfTWJ1Q=="

type detail struct {
	Name         string
	ID           string
	Children     []string
	Years        []int
	Availability []availability
	LastUpdate   string
	Description  string
}

type availability struct {
	Year      string
	Levels    []int
	Quarterly string
}

func main() {
	log := log.New(os.Stdout)
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Start(templates)
	log.Infof("App start on: %v:%v", "127.0.0.1", "8080")
	http.ListenAndServe(":8080", nil)
}

func sayHello(name string) string {
	return "Hello " + name + ":)"
}

func connectToDatabase() *sql.DB {
	log := log.New(os.Stdout)
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Errorf("Unable to connect to database: %v", fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func test() {
	//Statystyki malzenstwach
	resp2, err2 := http.Get("https://bdl.stat.gov.pl/api/v1/subjects/G535?lang=pl&format=json")

	body2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		panic(err2.Error())
	}
	var data2 detail
	json.Unmarshal(body2, &data2)
	color.Cyan("\nResults: %v\n\n", body2)
	fmt.Printf("Results: [%v]\n", len(data2.Availability))
	for _, s := range data2.Availability {
		fmt.Println(s.Year, s.Quarterly)
	}
}
