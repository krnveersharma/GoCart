package service

import (
	"GoCart/config"
	"GoCart/internal/helper"
	"GoCart/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}
