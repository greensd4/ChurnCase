// Author: Daniel Greenspan
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Costumer struct {
	Id string   `json:"id"`
	Gender  string   `json:"gender"`
	Segment   string   `json:"segment"`
	Years_customer   string   `json:"years_costumer"`
	No_of_complaints   string   `json:"no_of_complaints"`
	Contract_value   string   `json:"contract_value"`
}

func sendData(w http.ResponseWriter, r* http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/api")
	if csvFile, error := os.Open("./resources" + message); error != nil {
		fmt.Println("Unknown file\n")
		panic(error)
	} else {
		reader := csv.NewReader(bufio.NewReader(csvFile))
		reader.Comma = ';'
		var cArray []Costumer
		_, error := reader.Read()
			if error == io.EOF {
				return
			} else if error != nil {
				log.Fatal(error)
			}
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			ctmp := &Costumer{
				Id:               line[0],
				Gender:           line[1],
				Segment:          line[2],
				Years_customer:   line[3],
				No_of_complaints: line[4],
				Contract_value:   line[5],
			}
			cArray = append(cArray, *ctmp)
		}
		var costumersJson, err= json.Marshal(cArray)
		if err != nil {
			log.Fatal("cant cast to jason")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers",
			"{$_SERVER['HTTP_ACCESS_CONTROL_REQUEST_HEADERS']}")
		w.Write(costumersJson)
		fmt.Println("Sent data to Client...")
	}
}

func main() {
	http.HandleFunc("/", sendData)
	fmt.Println("Start churn server..")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}