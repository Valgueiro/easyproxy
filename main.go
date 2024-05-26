package main

import (
	"easyproxy/pkg/server"
)

func main() {
//   cmd.Execute()
	s := server.New()
	s.Run()
}
