package app

import (
	"github.com/vsPEach/Framework/internal/services"
)

type App struct {
	repo services.RepoService
}

func NewApp() (*App, error) {
	return nil, nil
}

func (a *App) Run() {

}
