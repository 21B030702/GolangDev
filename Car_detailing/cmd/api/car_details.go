package main

import (
	"car_detailing.arsennusip.net/internal/data"
	"car_detailing.arsennusip.net/internal/validator"
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
	carDetail := &data.CarDetail{
		Title:            input.Title,
		Description:      input.Description,
		DateOfProduction: input.DateOfProduction,
		Material:         input.Material,
		Price:            input.Price,
		Weight:           input.Weight,
	}
	v := validator.New()

	if data.ValidateCarDetail(v, carDetail); !v.Valid() {
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
	carDetail := data.CarDetail{
		ID:               id,
		Title:            "Car Detail",
		Description:      "web0fgheafbjkwnakedewdjncas",
		DateOfProduction: "07.10.2023",
		Weight:           1.5,
		Material:         "Aluminium",
		Price:            30,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"car_detail": carDetail}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
