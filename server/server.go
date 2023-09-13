package server

import "fmt"

type Server interface {
}

func RunServer() {

	err := getRouter()

	if err != nil {
		_ = fmt.Errorf("ERROR START SERVER")
		panic(err)
	}
}
