package main

import (
	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New().String()
	println(len(uuid))
}
