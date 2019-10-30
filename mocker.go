package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func main() {
	var inputString string
	var outputFmt string
	flag.StringVar(&outputFmt, "o", "flat", "Output format")

	// Checks if data is being sent to stdin from pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputString += scanner.Text()
		}

		flag.Parse()
	// Otherwise, parses inputString flag

	} else {
		flag.StringVar(&inputString,"i", "", " Input text to be transformed")
		flag.Parse()
		// Shows help text if no flag or stdin input is set
		if len(inputString) < 1 {
			flag.Usage()
		}
	}

	var newString = []rune{}

	for i := 0; i < len(inputString); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) != 0 {
			newString = append(newString, unicode.ToUpper(rune(inputString[i])))
		} else {
			newString = append(newString, unicode.ToLower(rune(inputString[i])))
		}
	}

	if outputFmt == "json" {
		outMap := map[string]string{
			"status": "ok",
			"text": string(newString),
		}
		out, err := json.Marshal(outMap)
		if err != nil {
			outMap["status"] = "error"
			outMap["text"] = "error converting input string to json"
		}
		fmt.Println(string(out))
	} else {
		fmt.Println(string(newString))
	}

	os.Exit(1)
}
