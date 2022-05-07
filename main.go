package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ppark2ya/jtcoin/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "templates/"
)
var templates *template.Template

type homeData struct {
	// template에서 사용하려면 접근제어자가 public 이어야함
	PageTitle string
	Blocks []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	// Must -> args으로 받은 template이 없으면 error 처리를해줌
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
