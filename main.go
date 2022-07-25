package main

import (
	"context"
	"log"

	"github.com/m4rc0nd35/db-golang/configs"
	"github.com/m4rc0nd35/db-golang/entities"
	"github.com/m4rc0nd35/db-golang/repositories"
)

func main() {
	repos := repositories.New(repositories.Options{
		WriterGorm: configs.GetWriterGorm(),
		ReaderGorm: configs.GetReaderGorm(),
	})

	for i := 0; i < 10; i++ {
		err := repos.User.Create(context.Background(), entities.User{
			Name:     "user-" + string(i),
			Email:    "-email@gmail.com",
			Password: "password",
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
