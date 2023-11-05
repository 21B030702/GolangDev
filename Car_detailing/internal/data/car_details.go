package data

import (
	"car_detailing.arsennusip.net/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.DB.QueryRowContext(ctx, query, args...).Scan(&detail.ID, &detail.CreatedAt, &detail.Description)
}

func (c CarDetailModel) Get(id int64) (*CarDetail, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
SELECT  id, created_at, title, description, dateofproduction, weight, material, price
FROM car_details
WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var detail CarDetail // Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
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
WHERE id = $5 AND price = $6
RETURNING price`
	args := []interface{}{
		detail.Title,
		detail.DateOfProduction,
		detail.Weight,
		pq.Array(detail.Material),
		detail.ID,
		detail.Price,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&detail.Price)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
func (c CarDetailModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
DELETE FROM car_details
WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := c.DB.ExecContext(ctx, query, id)
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
func (c CarDetailModel) GetAll(title string, genres []string, filters Filters) ([]*CarDetail, Metadata, error) {
	// Construct the SQL query to retrieve all movie records.
	query := fmt.Sprintf(`
SELECT id, created_at, title, dateofproduction, weight, material, price
FROM car_details
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
AND (material @> $2 OR $2 = '{}')
ORDER BY %s %s, id ASC
LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()
	// Initialize an empty slice to hold the movie data.
	totalRecords := 0
	details := []*CarDetail{}
	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Initialize an empty Movie struct to hold the data for an individual movie.
		var detail CarDetail

		err := rows.Scan(
			&totalRecords,
			&detail.ID,
			&detail.CreatedAt,
			&detail.Title,
			&detail.DateOfProduction,
			&detail.Weight,
			pq.Array(&detail.Material),
			&detail.Price,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		details = append(details, &detail)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return details, metadata, nil

}
