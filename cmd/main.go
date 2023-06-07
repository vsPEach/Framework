package main

import (
	"github.com/vsPEach/Framework/internal/repository/sql"
	internalhttp "github.com/vsPEach/Framework/internal/server/http"
	"log"
)

func main() {
	storage, err := sql.NewStorage()
	if err != nil {
		log.Fatalln(err)
	}
	s := internalhttp.NewServer(storage)
	_ = s.Start()
}
