package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jakab922/phone_storage/utils"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	batchSize = 100
)

func Clean(input string) string {
	reg, err := regexp.Compile("[^+0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	ret := reg.ReplaceAllString(input, "")

	// Here I'm assuming the the number is either in the
	// 0123456789 or in the +44123456789 format
	if strings.HasPrefix(ret, "+44") {
		return ret
	} else {
		return fmt.Sprintf("+44%v", strings.TrimPrefix(ret, "0"))
	}
}

func Send(batch []utils.PhoneData) {
	log.Printf("Sending batch to the server: %v", batch)
	serverAddress := os.Getenv("SERVER_ADDRESS")
	fullAddress := fmt.Sprintf("http://%v/store", serverAddress)

	serialized, err := json.Marshal(batch)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fullAddress, "application/json", bytes.NewBuffer(serialized))
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("The call to the server was not successful received an error code %v", resp.StatusCode)
	}
}

func main() {
	filePath := os.Getenv("FILE_PATH")

	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	batch := make([]utils.PhoneData, 0)

	line, err := reader.Read()
	if err != nil {
		log.Fatal("Failed to read the header row")
	}
	for {
		line, err = reader.Read()
		if err == io.EOF {
			if len(batch) != 0 {
				Send(batch)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		}

		data := utils.PhoneData{
			Name:        line[1],
			PhoneNumber: Clean(line[3]),
		}
		batch = append(batch, data)

		if len(batch) == batchSize {
			Send(batch)
			batch = batch[:0] // Cleaning the batch slice
		}
	}
}
