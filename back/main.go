package main

import (
	"back/repository"
	"context"
	"log"
)

func main() {
	dataSource := repository.GetDataSource()
	client := repository.InitDBConnect(dataSource)
	defer repository.CloseDBConnect()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
