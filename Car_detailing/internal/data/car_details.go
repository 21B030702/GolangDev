package data

import (
	"car_detailing.arsennusip.net/internal/validator"
	"time"
)

type CarDetail struct {
	ID               int64     `json:"id"`
	CreatedAt        time.Time `json:"-"`
	Title            string    `json:"title,omitempty"`
	Description      string    `json:"description,omitempty"`
	DateOfProduction string    `json:"date_of_production,omitempty"`
	Weight           Weight    `json:"weight,omitempty"`
	Material         string    `json:"material,omitempty"`
	Price            int64     `json:"price,omitempty"`
}

func ValidateCarDetail(v *validator.Validator, detail *CarDetail) {
	v.Check(detail.Title != "", "title", "must be provided")
	v.Check(len(detail.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(detail.DateOfProduction != "", "year", "must be provided")
	v.Check(detail.Description != "", "description", "must be provided")
	v.Check(detail.Material != "", "material", "must be provided")
	v.Check(detail.Description != "", "description", "must be provided")
	v.Check(detail.Weight != 0, "weight", "must be more than zero")
	v.Check(detail.Price != 0, "price", "must be more than zero")
}
