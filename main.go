package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var csvName string
var answers []string

func main() {
	flag.StringVar(&csvName, "filename", "problems.csv", "Give the name of the csv file.")
	flag.Parse()

	csvFile, err := os.Open(csvName)
	if err != nil {
		log.Fatal(err)
	}

	rd := csv.NewReader(csvFile)

	questions, err := rd.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bufioRd := bufio.NewReader(os.Stdin)
	for i, v := range questions {
		fmt.Print(i+1, ". ", v[0], "\n")
		ans, err := bufioRd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answers = append(answers, strings.TrimSpace(ans))
		bufioRd.Reset(os.Stdin)
	}

	fmt.Println()
	fmt.Println(answers)
}
