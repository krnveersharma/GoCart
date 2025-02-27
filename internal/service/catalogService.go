package service

import (
	"GoCart/config"
	"GoCart/internal/domain"
	"GoCart/internal/dto"
	"GoCart/internal/helper"
	"GoCart/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatalogService) EditCategory(input any) error {

	return nil
}

func (s CatalogService) DeleteCategory(input any) error {

	return nil
}

func (s CatalogService) GetCategories(input any) error {

	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil
	}

	return categories, err
}

func (s CatalogService) GetCategory(input any) error {

	return nil
}
