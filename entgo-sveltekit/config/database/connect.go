package database

import (
	"context"
	"errors"
	"fmt"

	"app/entity"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(name string) (client *entity.Client, err error) {
	client, err = entity.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", name))
	if err != nil {
		return nil, errors.New("failed opening connection to sqlite:" + err.Error())
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, errors.New("failed creating schema resources:" + err.Error())
	}
	return
}
