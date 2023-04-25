package main

import (
	"github.com/go-summer-dev/defaultcase"
	_ "github.com/lib/pq"
	"log"
	"my-arch/internal/handler"
	"my-arch/internal/pkg/config"
)

func main() {
	r := handler.InitApi()

	log.Fatalln(r.Run(config.GetPort()))
}

func init() {
	defaultcase.SetDefaultCase(defaultcase.Snak_case)
}
