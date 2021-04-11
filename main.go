package main

import (
	"bufio"
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	maxWidth, speed, words, position int
	bookFile                         string
)

//go:embed ATaleOfTwoCitiesByCharlesDickens.txt
var embeded embed.FS

func main() {

	defer unhideCursor()

	// parse flags
	flag.IntVar(&speed, "s", 200, "Reading speed in words per minute")
	flag.IntVar(&maxWidth, "lw", 80, "Line width - max number of characters per line")
	flag.IntVar(&words, "w", 2, "Field width - max number of words displayed")
	flag.IntVar(&position, "p", 0, "Position in book - line number")
	flag.StringVar(&bookFile, "f", "", "Book file")
	flag.Parse()

	// hide cursor. Unhide if interupt received
	fmt.Printf("\033[?25l")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		bookMark()
		unhideCursor()
		os.Exit(0)
	}()

	var file fs.File
	var err error

	file, err = embeded.Open("ATaleOfTwoCitiesByCharlesDickens.txt")
	if bookFile == "" {
		file, err = embeded.Open("ATaleOfTwoCitiesByCharlesDickens.txt")

		// if specific possition has not been provided start on chapter I
		// in the embeded book
		customPosition := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == "p" {
				customPosition = true
			}
		})
		if !customPosition {
			position = 105
		}

	} else {
		file, err = os.Open(bookFile)
	}
	if err != nil {
		unhideCursor()
		log.Fatalf("Error opening file. %s\n", err.Error())
	}
	defer file.Close()

	fmt.Printf("Baseline speed: %d words per minute.\n\n", speed)

	scanner := bufio.NewScanner(file)
	var lineCount int
	for scanner.Scan() {

		// seek line
		if position-1 > lineCount {
			lineCount++
			continue
		}

		if line := scanner.Text(); line != "" {
			position = lineCount + 1
			printLine(line)
		}

		lineCount++
	}
	if err := scanner.Err(); err != nil {
		unhideCursor()
		log.Fatal(err)
	}
}

// count returns number of runes in a string
func count(s string) int {
	r := []rune(s)
	return len(r)
}

func printLine(s string) {

	words := strings.Fields(s)
	var currentLine []string
	var lineLength int

	for _, word := range words {

		if lineLength+count(word) >= maxWidth {
			printPacer(currentLine)
			lineLength = 0
			currentLine = nil
		}

		lineLength += count(word) + 1
		currentLine = append(currentLine, word)
	}
	printPacer(currentLine)

}

func printPacer(line []string) {

	var prefixPad, currentLine, postfixPad string

	var begin, end int

	for i := range line {

		end++

		currentLine = strings.Join(line[begin:end], " ")

		if i > 0 && len(line) > 1 {
			if i >= words {
				prefixPad = strings.Repeat(" ", count(strings.Join(line[:begin], " "))+1)
			}
		}

		postfixPad = strings.Repeat(" ", maxWidth-count(prefixPad+currentLine))
		t := int(math.Round(60.0 / float64(speed) * 1000))
		fmt.Printf(" %s%s%s\r", prefixPad, currentLine, postfixPad)

		// slow down on new line, comma and full stop (double time)
		if strings.HasSuffix(currentLine, ".") || strings.HasSuffix(currentLine, ",") || i == 0 {
			time.Sleep(time.Duration(t) * time.Millisecond)
		}

		// slow down on long words
		x := count(line[i]) - 5
		if x > 0 {
			// 50ms for each letter above 5
			time.Sleep(time.Duration(x*50) * time.Millisecond)
		}

		// sleep duration between words
		time.Sleep(time.Duration(t) * time.Millisecond)

		if end-begin >= words {
			begin++
		}

	}
}

// unhideCursor sends terminal sequence to unhide cursor
func unhideCursor() {
	fmt.Printf("\033[?25h")
}

func bookMark() {

	const bm = `
Current possition in the book %d
If you would like to resume: %s -p %d
`
	print := fmt.Sprintf(bm, position, os.Args[0], position)
	fmt.Println(print)
}
