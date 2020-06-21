package main

import (
	"bytes"
	"fmt"
	"go/format"
	"gtats/fileparser"
	"log"
)

func main() {

	file, fs, f, err := fileparser.ParseFile("example/datatest.go")
	if err != nil {
		log.Fatalf("failed to parse file %v", err)
	}

	var buf bytes.Buffer
	for _, m := range file {
		BeautifyTag(&m)

	}
	err = format.Node(&buf, fs, f)
	if err != nil {
		log.Fatalf("failed to format node: %v\n", err)
	}
	fmt.Println(buf.String())

}
