package main

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/jdkato/prose.v2"
)


func goProse(docs string) (tag [][]string, entities [][]string){
	doc, err := prose.NewDocument(docs)
    if err != nil {
        log.Fatal(err)
	}

    for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
		var t []string
		t = append(t, tok.Text, tok.Tag)
		tag = append(tag, t)
	}

	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		var e []string
		e = append(e, ent.Text, ent.Label)
		entities = append(entities, e)
	}
	
	return tag, entities

}

type resultType struct {
	WhoamI    string
	Tag [][]string
	Ent [][]string
  }
  
func postHandler(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(string(body))

			var resultTag [][]string
			var resultEnt [][]string

		    resultTag,resultEnt  = goProse(string(body))

			result := resultType{"👋 from Pingzhou| ⛵", resultTag, resultEnt}

			fmt.Println(result)

			js, err := json.Marshal(result)
  			if err != nil {
    			http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
  			}

			w.Header().Set("Content-Type", "application/json")
  			w.Write(js)
		}
		
	
  }
  

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/go-nlp", postHandler)
	http.ListenAndServe(":"+os.Getenv("api_port"), mux)
 }
