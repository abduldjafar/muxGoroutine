package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func getInfoNoconc() []string {
	var datas []string
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	log.Println("processing file")


	for scanner.Scan() {
		datas = append(datas, scanner.Text())
	}

	return datas
}

func InfoNoconc( w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, 200,getInfoNoconc())
}
