package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	md "github.com/Skarlso/html-to-markdown"
)

func main() {
	url := "https://blog.golang.org/godoc-documenting-go-code"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	content := doc.Find("#content")

	conv := md.NewConverter(md.DomainFromURL(url), true, nil)
	markdown := conv.Convert(content)

	fmt.Println(markdown)
}
