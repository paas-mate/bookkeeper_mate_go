package main

import "bookkeeper_mate_go/pkg/bk"

func main() {
	err := bk.Config()
	if err != nil {
		panic(err)
	}
}
