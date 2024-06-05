package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	InfoLogger.Println("Starting Service!")
	t := time.Now()
	readFromFile()
	fmt.Println(len(trans), time.Since(t))
	InfoLogger.Printf("JSON file with %d length took %s time\n", len(trans), time.Since(t))
}

var (
	fileName       = "data.json"
	loggerFileName = "logger.text"
	trans          []StoreProcessor
	InfoLogger     *log.Logger
	WarnLogger     *log.Logger
	ErrorLogger    *log.Logger
)

type StoreProcessor struct {
	Processtime string `json:"processtime"`
	Storeid     string `json:"storeid"`
	Storecode   string `json:"storecode"`
}

// Logging the Errors in logger file for readability purpose.
func init() {
	logFile, err := os.OpenFile(loggerFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Unable to create Logger file:", err.Error())
		return
	}
	log.SetOutput(logFile)
	InfoLogger = log.New(logFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(logFile, "Warn:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func readFromFile() {
	// Opening the JSON file
	file, err := os.Open(fileName)
	if err != nil {
		ErrorLogger.Println(err.Error())
		return
	}
	defer file.Close()

	//Readinf line by line from JSON input file
	buff := bufio.NewReader(file)
	for {
		b, _, err := buff.ReadLine()
		if err != nil {
			//Stopping the process when it reaches end of the file
			if io.EOF == err {
				WarnLogger.Println(err.Error())
				return
			}
			return
		}

		//Serializing the json data into struct
		var t StoreProcessor
		errD := json.NewDecoder(strings.NewReader(string(b))).Decode(&t)
		if errD != nil {
			WarnLogger.Println("Error from worker:", err.Error())
			continue
		}
		trans = append(trans, t)
	}
}
