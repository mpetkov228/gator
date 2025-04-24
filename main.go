package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("mihail")

	cfg = config.Read()
	fmt.Printf("%v\n", cfg.DbUrl)
	fmt.Printf("%v\n", cfg.CurrentUserName)
}
