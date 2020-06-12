package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// create("sample", "csv")
	// read("sample", "csv")
}

// func create
func create(fname, fformat string) {
	f, err := os.Create(fmt.Sprint(fname + "." + fformat))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := csv.NewWriter(f)

	data := []string{"ini", "data", "kita", "punya"}

	errWriter := writer.Write(data)
	if errWriter != nil {
		log.Fatal(errWriter)
	}
	writer.Flush()
}

// func read
func read(fname, fformat string) {
	csvf, err := os.Open(fmt.Sprint(fname + "." + fformat))
	if err != nil {
		log.Fatal(err)
	}

	reader, err := csv.NewReader(csvf).ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range reader {
		fmt.Println(r)
	}

}
