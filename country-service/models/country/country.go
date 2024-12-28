package models

import "time"

type Country struct {
	ID              string            `json:"id"`
	ZoneId          string            `json:"zone_id"`
	Name            string            `json:"name"`
	Code            string            `json:"code"`
	Code2           string            `json:"code2"`
	PhoneCode       string            `json:"phone_code"`
	Image           string            `json:"image"`
	IsConflict      bool              `json:"is_conflict"`
	DataLink        *string           `json:"data_link"`
	ProhibitedItems *string           `json:"prohibited_items"`
	TotalDiaspora   *int              `json:"total_diaspora"`
	TotalUsers      *int              `json:"total_users"`
	Shipment        *float64          `json:"shipment"`
	ShipmentAverage *float64          `json:"shipment_average"`
	ShipmentValue   *float64          `json:"shipment_value"`
	MarketShare     *float64          `json:"market_share"`
	IsAfrica        *int              `json:"is_africa"`
	IsNorthAmerica  *int              `json:"is_north_america"`
	IsSouthAmerica  *int              `json:"is_south_america"`
	IsOceania       *int              `json:"is_oceania"`
	IsAsia          *int              `json:"is_asia"`
	IsEurope        *int              `json:"is_europe"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       *time.Time        `json:"deleted_at"`
	Zone            *Zone             `json:"zone,omitempty"`
	ShippingReviews []*ShippingReview `json:"shipping_reviews"`
}

type Zone struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ShippingReview struct {
	ID              *string `json:"id"`
	UserID          *string `json:"user_id"`
	CountryID       string  `json:"country_id"`
	CombinedOrderID *string `json:"combined_order_id"`
	Rating          *int    `json:"rating"`
	Message         *string `json:"message"`
	IsShow          *int    `json:"is_show"`
}
