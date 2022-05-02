package main

import (
	"fmt"
	"github.com/Khalturin/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
