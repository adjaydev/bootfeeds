package main

import (
	"bootfeeds/internal/config"
	"fmt"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("jay")
	fmt.Println(cfg.DbUrl)

}
