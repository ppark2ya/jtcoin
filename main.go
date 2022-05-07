package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ppark2ya/jtcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	// template에서 사용하려면 접근제어자가 public 이어야함
	PageTitle string
	Blocks []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	// Must -> args으로 받은 template이 없으면 error 처리를해줌
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
