package service

import (
	models "fulka-api/models/country"
	"fulka-api/repository"
)

type CountryService interface {
	GetAllCountries(pageSize, offset int) ([]models.Country, error)
	CountAllCountry() (int, error)
	GetByIdCountries(id string) (models.Country, error)
}

type countryService struct {
	repo repository.CountryRepository
}

func NewCountryService(repo repository.CountryRepository) CountryService {
	return &countryService{repo}
}

func (s *countryService) GetAllCountries(pageSize, offset int) ([]models.Country, error) {
	return s.repo.GetAllCountries(pageSize, offset)
}

func (s *countryService) CountAllCountry() (int, error) {
	return s.repo.CountAllCountry()
}

func (s *countryService) GetByIdCountries(id string) (models.Country, error) {
	return s.repo.GetByIdCountries(id)
}
