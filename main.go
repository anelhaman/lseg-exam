package main

import (
	"log"
	"lseg-exam/config"
	"lseg-exam/service"
)

func main() {

	config.InitEnv()

	svc, err := service.NewEC2()
	if err != nil {
		log.Fatalf("init error: %v", err)
	}

	instances, err := svc.DescribeInstances()
	if err != nil {
		log.Fatalf("describe error: %v", err)
	}

	for _, inst := range instances {
		log.Printf("%+v", inst)
	}
}
