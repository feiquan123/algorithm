package main

type Msg struct {
	string
}

func main() {
	Debug(Msg{"msg xxx"})
}

// debug env
//go:generate go run -tags debug .

// produce env
//go:generate go run .
