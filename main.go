package main

import (
	_ "fmt"
	"log"
	"os"
	"strings"
"strconv"
	_"embed"
)

var (
	//go:embed input.txt
	input string

	//go:embed "lines per slide.txt"
	linesPerSlide string

	outputFile = "output.txt"
)

func lines(v string) []string {

	return strings.Split(v, "\n")
}

// clean remove empty lines and spaces around the none empty lines
func clean(ls []string) []string {

	rs := []string{}

	for _, ln := range ls {
		ln := strings.TrimSpace(ln)
		if len(ln) != 0 {
ln = ln + "\n"
			rs = append(rs, ln)
		}
	}

	return rs
}

// ashara: break `ls` into sets of len `count` and introduces `br` after evening
func ashara(ls []string, count int, br string) []string {

	if len(ls) < count {
		return ls
	}

	rs := []string{}
	rs = append(rs, ls[:count]...)
	rs = append(rs, br)
	rs = append(rs, ashara(ls[count:], count, br)...)

	return rs
}

func write(ls []string) {

	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(strings.Join(ls, ""))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	log.Println("started...")

	var (
		ls = lines(input)

		// number of inputs before br
		count = 2

		// br: stands for break
		// is a divider b/2n each set of line
		br = "\n"
	)

	count, err := strconv.Atoi(linesPerSlide)
if err != nil {
	log.Fatal(err)
}

	log.Println("input loaded")

	ls = clean(ls)
	log.Println("input cleaned")

	ls = ashara(ls, count, br)
	log.Println("input organised")

	write(ls)
	log.Print("write to file successful")
}
