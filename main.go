package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var csvName string
var timelimit int
var total int
var total_correct int
var ch chan bool

func main() {
	flag.StringVar(&csvName, "filename", "problems.csv", `Give the name of the csv file. Defaults to "problems.csv".`)
	flag.IntVar(&timelimit, "timelimit", 30, "Give the time limit in seconds for the whole quiz. Defaults to 30.")
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

	ch = make(chan bool)

	go func() {
		bufioRd := bufio.NewReader(os.Stdin)
		for i, v := range questions {
			fmt.Print(i+1, ". ", v[0], "\n")
			ans, err := bufioRd.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			if strings.TrimSpace(ans) == v[1] {
				total_correct++
			}
			total++
			bufioRd.Reset(os.Stdin)
		}
		ch <- true
	}()

	select {
	case <-ch:
		fmt.Printf("\nTotal: %d\nCorrect: %d\n", total, total_correct)
	case <-time.After(time.Second * time.Duration(timelimit)):
		fmt.Print("\n\nTime ran out!\n")
	}
}
