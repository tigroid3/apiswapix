package main

import (
	"github.com/joho/godotenv"
	"github.com/tigroid3/apiswapix/v4"
)

func main() {
	_ = godotenv.Load(".env")

	v4.Run()
}
