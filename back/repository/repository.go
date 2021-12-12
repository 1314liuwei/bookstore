package repository

import (
	"back/repository/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var once sync.Once
var client *ent.Client

func InitDBClient(dataSource string) *ent.Client {
	once.Do(func() {
		var err error
		client, err = ent.Open("mysql", dataSource)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return client
}

func CloseDBClient() {
	once.Do(func() {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	})
}

func GetDBClient() *ent.Client {
	if client != nil {
		return client
	} else {
		return InitDBClient(GetDataSource())
	}
}

func GetDataSource() string {
	return "root:104494lw!@tcp(116.62.229.23:3306)/bookstore?parseTime=True"
}
