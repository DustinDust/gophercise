package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	score := 0

	var path = flag.String("input-file", "problem.csv", "Path to the input file")
	var timer = flag.Int("timer", 5, "Time limit")
	flag.Parse()

	csvFile, err := os.Open(*path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	t := time.NewTimer(time.Duration(*timer) * time.Second)
	go func() {
		<-t.C
		exit(fmt.Sprintf("\n\nTime out! You scored %v out of %v", score, len(csvLines)))
	}()

	for _, qa := range csvLines {
		fmt.Printf("%v = ", qa[0])
		var awns string
		fmt.Scanln(&awns)

		fmtedAnswer := strings.TrimSpace(strings.ToLower(awns))
		if fmtedAnswer == qa[1] {
			score += 1
		}
	}

	exit(fmt.Sprintf("\n\nYou score is %v out of %v", score, len(csvLines)))
}

func exit(msg string) {
	fmt.Printf("%v", msg)
	os.Exit(0)
}
