package service

import (
	"database/sql"
	models "fulka-api/models/country"
	"fulka-api/repository"
)

type CountryService interface {
	GetAllCountries(pageSize, offset int) ([]models.Country, error)
	CountAllCountry() (int, error)
	GetByIdCountries(id string) (models.Country, error)
	CreateCountry(country *models.Country) error
}

type countryService struct {
	repo repository.CountryRepository
	db   *sql.DB
}

func NewCountryService(repo repository.CountryRepository, db *sql.DB) CountryService {
	return &countryService{
		repo: repo,
		db:   db,
	}
}

func (s *countryService) GetAllCountries(pageSize, offset int) ([]models.Country, error) {
	countries, err := s.repo.GetAllCountries(pageSize, offset)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *countryService) CountAllCountry() (int, error) {
	return s.repo.CountAllCountry()
}

func (s *countryService) GetByIdCountries(id string) (models.Country, error) {
	return s.repo.GetByIdCountries(id)
}

func (s *countryService) CreateCountry(country *models.Country) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = s.repo.CreateCountry(country, tx)
	if err != nil {
		return err
	}

	return nil
}
