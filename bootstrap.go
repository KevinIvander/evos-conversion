package main

import (
	"encoding/json"
	"evos-conversion/handlers"
	"evos-conversion/repositories"
	"evos-conversion/services"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var dbCurrencyRepository repositories.DBCurrencyRepositoryContract

var dbConversionRateRepository repositories.DBConversionRateRepositoryContract

var currencyService services.CurrencyServiceContract

var conversionRateService services.ConversionRateServiceContract

func initRepositories() {
	dbCurrencyRepository = repositories.NewDBCurrencyRepository(dbConn)
	dbConversionRateRepository = repositories.NewDBConversionRateRepository(dbConn)
}

func initServices() {
	currencyService = services.NewCurrencyService(dbCurrencyRepository)
	conversionRateService = services.NewConversionRateService(dbConversionRateRepository)
}

func serveHTTP() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := "Evos CurrencyService"
		res, _ := json.Marshal(payload)
		w.Write(res)
	}))
	currencyHandler := handlers.NewCurrencyHandler(currencyService)
	conversionRateHandler := handlers.NewConversionRateHandler(conversionRateService)
	router.Mount("/currency", currencyHandler.GetRoutes())
	router.Mount("/conversion_rate", conversionRateHandler.GetRoutes())

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("App running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
