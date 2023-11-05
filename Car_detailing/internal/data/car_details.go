package data

import (
	"car_detailing.arsennusip.net/internal/validator"
	"database/sql"
	"errors"
	"github.com/lib/pq"
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
	//v.Check(detail.Description != "", "description", "must be provided")
	v.Check(detail.Material != "", "material", "must be provided")
	v.Check(detail.Weight != 0, "weight", "must be more than zero")
	v.Check(detail.Price != 0, "price", "must be more than zero")
}

type CarDetailModel struct {
	DB *sql.DB
}

func (c CarDetailModel) Insert(detail *CarDetail) error {

	query := `
			INSERT INTO car_details (title, dateofproduction, weight, material)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at, description`
	args := []interface{}{detail.Title, detail.DateOfProduction, detail.Weight, pq.Array(detail.Material)}
	return c.DB.QueryRow(query, args...).Scan(&detail.ID, &detail.CreatedAt, &detail.Description)
}

func (c CarDetailModel) Get(id int64) (*CarDetail, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
SELECT id, created_at, title, description, dateofproduction, weight, material, price
FROM car_details
WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var detail CarDetail // Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	err := c.DB.QueryRow(query, id).Scan(
		&detail.ID,
		&detail.CreatedAt,
		&detail.Title,
		&detail.Description,
		&detail.DateOfProduction,
		&detail.Weight,
		pq.Array(&detail.Material),
		&detail.Price,
	)
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound
	// error instead.
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &detail, nil
}
func (c CarDetailModel) Update(detail *CarDetail) error {
	query := `
UPDATE car_details
SET title = $1, dateofproduction = $2, weight = $3, material = $4, price = price + 300
WHERE id = $5
RETURNING price`
	args := []interface{}{
		detail.Title,
		detail.DateOfProduction,
		detail.Weight,
		pq.Array(detail.Material),
		detail.ID,
	}
	return c.DB.QueryRow(query, args...).Scan(&detail.Price)
	return nil
}
func (c CarDetailModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
DELETE FROM car_details
WHERE id = $1`
	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil

}
