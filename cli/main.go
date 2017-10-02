package main

import (
	"github.com/yaskoo/chrombo"
)

func main() {
	b, err := chrombo.NewBrowser("http://localhost:9999")
	if err != nil {
		panic(err)
	}

	p := b.Pages[0]

	p.Navigate("https://9gag.com/fresh")
	p.Evaluate("JSON.stringify({test:123})")

	b.NewPage("")
}
