package main

import "embed"

//go:embed test/content
var content embed.FS

func main() {
	println(content.Name()) // Output: test/content
}

