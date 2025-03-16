package service

import (
	"GoCart/config"
	"GoCart/internal/domain"
	"GoCart/internal/dto"
	"GoCart/internal/helper"
	"GoCart/internal/repository"
	"errors"
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

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {

	exitCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, err
	}
	if len(input.Name) > 0 {
		exitCat.Name = input.Name
	}
	if input.ParentId > 0 {
		exitCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		exitCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		exitCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(exitCat)
	return updatedCat, nil
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		return errors.New("category does not exist")
	}
	return nil
}

func (s CatalogService) GetCategories() ([]*domain.Category, error) {

	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, err
	}

	return categories, err
}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {

	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, err
	}
	return cat, nil
}
