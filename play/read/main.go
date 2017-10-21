package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(os.Args[0])
	fmt.Println(filepath.Base(os.Args[0]))
	fmt.Println(filepath.HasPrefix(filepath.Base(os.Args[0]), "m"))
}
