package main

import "github.com/dev-vinicius-andrade/go-cobra-example/cmd/foo"

func main() {
	foo.CreateCommand().Execute()

}
