package repository

import (
	"back/repository/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var client *ent.Client

func InitDBConnect(dataSource string) *ent.Client {
	var err error
	client, err = ent.Open("mysql", dataSource)
	if err != nil {
		return nil
	}
	return client
}

func CloseDBConnect() {
	err := client.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func GetDataSource() string {
	return "root:104494lw!@tcp(116.62.229.23:3306)/bookstore?parseTime=True"
}
