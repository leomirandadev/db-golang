package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leomirandadev/db-golang/configs"
	"github.com/leomirandadev/db-golang/entities"
	"github.com/leomirandadev/db-golang/repositories"
)

func main() {
	repos := repositories.New(repositories.Options{
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderGorm: configs.GetReaderGorm(),
		WriterGorm: configs.GetWriterGorm(),
	})

	for i := 0; i < 10000; i++ {

		err := repos.User.Create(context.Background(), entities.User{
			Name:     "Leonardo Miranda",
			Email:    "email@email.com",
			NickName: "leomirandadev",
			Password: "e8d95a51f3af4a3b134bf6bb680a213a",
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Success")
}
