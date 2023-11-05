package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/car_details", app.listCarDetailsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/car_details", app.createCarDetailsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/car_details/:id", app.showCarDetailsHandler)
	router.HandlerFunc(http.MethodPut, "/v1/car_details/:id", app.updateCarDetailHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/car_details/:id", app.deleteCarDetailHandler)

	return app.recoverPanic(app.rateLimit(router))
}
