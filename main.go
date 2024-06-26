package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var csvName string
var timelimit int
var total int
var total_correct int
var ch chan bool

func shuffle(questions [][]string) {
	n := len(questions)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}
}

func main() {
	flag.StringVar(&csvName, "path", "problems.csv", `Relative path to the csv quiz file`)
	flag.IntVar(&timelimit, "timelimit", 30, "Time limit in seconds for the quiz")
	flag.Parse()

	csvFile, err := os.Open(csvName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	rd := csv.NewReader(csvFile)

	questions, err := rd.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	shuffle(questions)

	ch = make(chan bool)

	go func() {
		bufioRd := bufio.NewReader(os.Stdin)
		for i, v := range questions {
			fmt.Print(i+1, ". ", v[0], " = ")
			ans, err := bufioRd.ReadString('\n')
			fmt.Println()
			if err != nil {
				log.Fatal(err)
			}
			cleanAns := strings.ToLower(strings.TrimSpace(ans))
			if cleanAns == v[1] {
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
