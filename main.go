package main

import (
	"fmt"
	"log"
	"lseg-exam/service"
)

func main() {

	fetcher, err := service.NewEC2MetaFetcher()
	if err != nil {
		log.Fatalf("init error: %v", err)
	}

	meta, err := fetcher.GetMetadata()
	if err != nil {
		log.Fatalf("fetch error: %v", err)
	}

	fmt.Printf("Metadata: %+v\n", meta)
}
