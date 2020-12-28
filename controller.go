package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"sync"
)

func readFIle(waitGroup *sync.WaitGroup,result chan *bufio.Scanner)  {
	defer waitGroup.Done()

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	result <-scanner
}

func getInfo(waitGroup *sync.WaitGroup, scannerChan chan *bufio.Scanner, result chan []string) {
	defer waitGroup.Done()
	var datas []string

	scanner := <- scannerChan


	log.Println("processing file")


	for scanner.Scan() {
		datas = append(datas, scanner.Text())
	}

	result <- datas
}

func proc() []string {
	var wg sync.WaitGroup
	c1 := make(chan []string)
	c2 := make(chan  *bufio.Scanner)

	wg.Add(1)
	go readFIle(&wg,c2)

	wg.Add(1)
	go getInfo(&wg,c2,c1)

	result := <-c1
	wg.Wait()
	close(c1)

	return result
}

func Info(waitGroup *sync.WaitGroup, w http.ResponseWriter, r *http.Request) {
	defer waitGroup.Done()

	RespondJSON(w, 200,proc() )
}
