package main

import (
	"compass.com/go-homework/comment_count/counter"
	"compass.com/go-homework/comment_count/reader"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		printHelp()
	} else {
		dir := args[0]
		if err := countCommentLines(dir); err != nil {
			fmt.Println(err)
		}
	}
}

func printHelp() {
	fmt.Println("usage: \n\tgo run . <directory>")
}

func countCommentLines(dir string) error {
	debugMode := strings.ToLower(strings.TrimSpace(os.Getenv("DEBUG"))) == "true"
	c := counter.NewCounter()
	c.SetDebug(debugMode)
	sortReader := reader.NewSortReader()
	sortedList, err := sortReader.Read(dir, reader.NewSuffixFilter([]string{".c", ".cpp", ".h", ".hpp"}))
	if err != nil {
		return err
	}
	for _, file := range sortedList {
		res, err := c.Count(file)
		if err != nil {
			log.Fatalf("Count err:%s", err.Error())
		}
		fmt.Println(res.ToPrintString())
	}
	return nil
}
