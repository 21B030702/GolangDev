package main

import (
	"car_detailing.arsennusip.net/internal/data"
	"car_detailing.arsennusip.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createCarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title            string      `json:"title"`
		Description      string      `json:"description"`
		DateOfProduction string      `json:"date_of_production"`
		Weight           data.Weight `json:"weight"`
		Material         string      `json:"material"`
		Price            int64       `json:"price"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	detail := &data.CarDetail{
		Title:            input.Title,
		Description:      input.Description,
		DateOfProduction: input.DateOfProduction,
		Material:         input.Material,
		Price:            input.Price,
		Weight:           input.Weight,
	}
	v := validator.New()

	if data.ValidateCarDetail(v, detail); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showCarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	detail, err := app.models.CarDetails.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"car_detail": detail}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	//detail := &data.CarDetail{
	//	ID:               id,
	//	Title:            "Car Detail",
	//	Description:      "web0fgheafbjkwnakedewdjncas",
	//	DateOfProduction: "07.10.2023",
	//	Weight:           1.5,
	//	Material:         "Aluminium",
	//	Price:            30,
	//}
	//err = app.models.CarDetails.Insert(detail)
	//if err != nil {
	//	app.serverErrorResponse(w, r, err)
	//}
	//headers := make(http.Header)
	//headers.Set("Location", fmt.Sprintf("/v1/car_details/%d", detail.ID))
	//err = app.writeJSON(w, http.StatusOK, envelope{"car_detail": detail}, nil)
}

func (app *application) updateCarDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	detail, err := app.models.CarDetails.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Title            string      `json:"title"`
		DateOfProduction string      `json:"year"`
		Weight           data.Weight `json:"weight"`
		Material         string      `json:"material"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	detail.Title = input.Title
	detail.DateOfProduction = input.DateOfProduction
	detail.Weight = input.Weight
	detail.Material = input.Material

	v := validator.New()

	if data.ValidateCarDetail(v, detail); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.CarDetails.Update(detail)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"detail": detail}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteCarDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.CarDetails.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "detail successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
