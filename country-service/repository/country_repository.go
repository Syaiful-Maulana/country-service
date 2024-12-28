package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	models "fulka-api/models/country"
)

type CountryRepository interface {
	GetAllCountries(page int, pageSize int) ([]models.Country, error)
	GetByIdCountries(id string) (models.Country, error)
	CountAllCountry() (int, error)
}

type countryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) CountryRepository {
	return &countryRepository{db}
}

func (r *countryRepository) CountAllCountry() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM countries WHERE deleted_at IS NULL"
	if err := r.db.QueryRow(query).Scan(&count); err != nil {
		log.Printf("Error executing count query: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *countryRepository) GetAllCountries(page int, pageSize int) ([]models.Country, error) {
	count, err := r.CountAllCountry()
	if err != nil {
		return nil, err
	}

	fmt.Println("count:", count)

	offset := (page - 1) * pageSize

	query := `
		SELECT
			c.id AS country_id,
			c.zone_id,
			c.name,
			c.code,
			c.code2,
			c.phone_code,
			c.image,
			c.is_conflict,
			c.data_link,
			c.prohibited_items,
			c.total_diaspora,
			c.total_users,
			c.shipment,
			c.shipment_average,
			c.shipment_value,
			c.market_share,
			c.is_africa,
			c.is_north_america,
			c.is_south_america,
			c.is_oceania,
			c.is_asia,
			c.is_europe,
			c.created_at AS country_created_at,
			c.updated_at AS country_updated_at,
			c.deleted_at AS country_deleted_at,
			z.id AS zone_id,
			z.name AS zone_name,
			z.code AS zone_code,
			z.created_at AS zone_created_at,
			z.updated_at AS zone_updated_at,
			z.deleted_at AS zone_deleted_at
		FROM countries c
		LEFT JOIN zones z ON c.zone_id = z.id
		WHERE c.deleted_at IS NULL
		ORDER BY c.name ASC
		LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var countriesList []models.Country
	for rows.Next() {
		var country models.Country
		var deletedAt sql.NullTime
		var zone models.Zone

		err := rows.Scan(
			&country.ID, &country.ZoneId, &country.Name, &country.Code, &country.Code2,
			&country.PhoneCode, &country.Image, &country.IsConflict, &country.DataLink,
			&country.ProhibitedItems, &country.TotalDiaspora, &country.TotalUsers,
			&country.Shipment, &country.ShipmentAverage, &country.ShipmentValue, &country.MarketShare,
			&country.IsAfrica, &country.IsNorthAmerica, &country.IsSouthAmerica, &country.IsOceania,
			&country.IsAsia, &country.IsEurope, &country.CreatedAt, &country.UpdatedAt,
			&deletedAt, &zone.ID, &zone.Name, &zone.Code, &zone.CreatedAt, &zone.UpdatedAt, &zone.DeletedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		if deletedAt.Valid {
			country.DeletedAt = &deletedAt.Time
		}
		country.Zone = &zone
		countriesList = append(countriesList, country)
	}

	// Mengambil shipping reviews
	if len(countriesList) > 0 {
		countryIDs := []string{}
		for _, c := range countriesList {
			countryIDs = append(countryIDs, fmt.Sprintf("'%s'", c.ID))
		}

		reviewsQuery := fmt.Sprintf(`
		SELECT sr.id, sr.user_id, sr.country_id
		FROM shipping_reviews sr
		WHERE sr.country_id IN (%s)`, strings.Join(countryIDs, ","))

		reviewRows, err := r.db.Query(reviewsQuery)
		if err != nil {
			log.Printf("Error executing reviews query: %v", err)
			return nil, err
		}
		defer reviewRows.Close()

		reviewMap := make(map[string][]*models.ShippingReview)
		for reviewRows.Next() {
			var review models.ShippingReview
			if err := reviewRows.Scan(&review.ID, &review.UserID, &review.CountryID); err != nil {
				log.Printf("Error scanning review row: %v", err)
				continue
			}
			reviewMap[review.CountryID] = append(reviewMap[review.CountryID], &review)
		}

		for i, c := range countriesList {
			if reviews, exists := reviewMap[c.ID]; exists {
				countriesList[i].ShippingReviews = reviews
			} else {
				countriesList[i].ShippingReviews = []*models.ShippingReview{}
			}
		}
	}

	return countriesList, nil
}

func (r *countryRepository) GetByIdCountries(id string) (models.Country, error) {
	query := `
		SELECT
			c.id AS country_id,
			c.zone_id,
			c.name,
			c.code,
			c.code2,
			c.phone_code,
			c.image,
			c.is_conflict,
			c.data_link,
			c.prohibited_items,
			c.total_diaspora,
			c.total_users,
			c.shipment,
			c.shipment_average,
			c.shipment_value,
			c.market_share,
			c.is_africa,
			c.is_north_america,
			c.is_south_america,
			c.is_oceania,
			c.is_asia,
			c.is_europe,
			c.created_at,
			c.updated_at,
			c.deleted_at,
			z.id AS zone_id,
			z.name AS zone_name,
			z.code AS zone_code,
			z.created_at AS zone_created_at,
			z.updated_at AS zone_updated_at,
			z.deleted_at AS zone_deleted_at
		FROM countries c
		LEFT JOIN zones z ON c.zone_id = z.id
		WHERE c.id = ?
		LIMIT 1`

	var country models.Country
	var deletedAt sql.NullTime
	var zone models.Zone

	err := r.db.QueryRow(query, id).Scan(
		&country.ID, &country.ZoneId, &country.Name, &country.Code, &country.Code2,
		&country.PhoneCode, &country.Image, &country.IsConflict, &country.DataLink,
		&country.ProhibitedItems, &country.TotalDiaspora, &country.TotalUsers,
		&country.Shipment, &country.ShipmentAverage, &country.ShipmentValue, &country.MarketShare,
		&country.IsAfrica, &country.IsNorthAmerica, &country.IsSouthAmerica, &country.IsOceania,
		&country.IsAsia, &country.IsEurope, &country.CreatedAt, &country.UpdatedAt,
		&deletedAt, &zone.ID, &zone.Name, &zone.Code, &zone.CreatedAt, &zone.UpdatedAt, &zone.DeletedAt,
	)
	if err != nil {
		return models.Country{}, err
	}

	if deletedAt.Valid {
		country.DeletedAt = &deletedAt.Time
	}

	country.Zone = &zone

	reviewsQuery := `
		SELECT
			id, user_id, country_id, combined_order_id, rating, message, is_show
		FROM shipping_reviews
		WHERE country_id = ?`

	rows, err := r.db.Query(reviewsQuery, id)
	if err != nil {
		return models.Country{}, fmt.Errorf("error querying reviews: %w", err)
	}
	defer rows.Close()

	var reviews []*models.ShippingReview
	for rows.Next() {
		var review models.ShippingReview
		err := rows.Scan(
			&review.ID, &review.UserID, &review.CountryID,
			&review.CombinedOrderID, &review.Rating, &review.Message, &review.IsShow,
		)
		if err != nil {
			return models.Country{}, fmt.Errorf("error scanning review row: %w", err)
		}
		reviews = append(reviews, &review)
	}

	country.ShippingReviews = reviews

	return country, nil
}
