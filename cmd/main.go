package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/vsPEach/Framework/internal/entity"
	"github.com/vsPEach/Framework/internal/repository/sql"
	"time"
)

func main() {
	//a, err := app.NewApp()
	//if err != nil {
	//	return
	//}
	//a.Run()
	st := sql.NewStorage()
	st.Create(context.Background(), entity.User{
		ID:          uuid.New(),
		Username:    "Peach",
		Email:       "cklx2000@mail.ri",
		Role:        0,
		IsConfirmed: false,
		Password:    "ccc",
		CreatedAt:   time.Now(),
	})
	st.Create(context.Background(), entity.Article{
		ID:        uuid.New(),
		AuthorID:  uuid.UUID{},
		Title:     "Article",
		Text:      "article",
		CreatedAt: time.Now(),
	})
	st.Create(context.Background(), entity.Comment{
		ID:        uuid.UUID{},
		AuthorID:  uuid.UUID{},
		ArticleID: uuid.UUID{},
		Text:      "Text",
		CreatedAt: time.Now(),
	})
}
