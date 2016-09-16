package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	re0 := regexp.MustCompile("(?U)^.+/")
	re1 := regexp.MustCompile("[^a-z0-9]")
	re2 := regexp.MustCompile("-$")
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)

	}
	result := string(in)
	result = re0.ReplaceAllString(result, "")
	result = re1.ReplaceAllString(result, "-")
	maxLength := 24
	if len(result) > maxLength {
		result = result[:24]
	}
	result = re2.ReplaceAllString(result, "")
	fmt.Println(result)
}
