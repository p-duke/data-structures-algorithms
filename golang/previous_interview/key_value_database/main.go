package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	
	for s.Scan() {
		text := s.Text()
		fmt.Println("Token:", text)
	}

	if err := s.Err(); err != nil {
		// Handle/report any error (EOF will not be reported)
		log.Fatal(err)
	}
}
