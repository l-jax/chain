package main

import (
	"fmt"
	"log"
)

const testUrl = "https://github.com/l-jax/chain/pull/1"

func main() {
	checkTestPr()
	removeLabel(testUrl, "test")
	checkTestPr()
	addLabel(testUrl, "test")
	viewPr(testUrl)
}

func checkTestPr() error {
	pr, err := getPr(testUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pr.Title, pr.Body, pr.Mergeable, pr.Labels, pr.State)
	return err
}
