package main

import (
	"fmt"
)

func main() {
	prs, _ := getPullRequests()
	fmt.Println(prs[0].title, prs[0].branch, prs[0].conditions[0].success())
}
