package handlers

import (
	"evos-conversion/handlers/requests"
	"evos-conversion/handlers/responses"
	"evos-conversion/models"
	"evos-conversion/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CurrencyHandler serve HTTP routes for currency
type CurrencyHandler struct {
	currencyService services.CurrencyServiceContract
}

// NewCurrencyHandler returns new CurrencyHandler
func NewCurrencyHandler(currencyService services.CurrencyServiceContract) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: currencyService,
	}
}

// GetRoutes mount routes
func (handler *CurrencyHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handler.Get)
	router.Post("/", handler.Store)

	return router
}

// Get :nodoc
func (handler *CurrencyHandler) Get(writer http.ResponseWriter, request *http.Request) {
	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid `limit` query"))
		return
	}
	if limit == 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(request.URL.Query().Get("offset"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid `offset` query"))
		return
	}

	currencyList, err := handler.currencyService.Get(limit, offset)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	response := responses.NewCurrencyListResponse(currencyList, limit, offset)
	render.Render(writer, request, response)
}

// Store :nodoc
func (handler *CurrencyHandler) Store(writer http.ResponseWriter, request *http.Request) {
	var requestData requests.CreateCurrencyRequest
	if err := render.Bind(request, &requestData); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(err.Error()))
		return
	}
	currency := models.Currency{}
	currency.Name = requestData.Name
	currencyData := models.Currency{}
	currencyData, err := handler.currencyService.Create(currency)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	currencyResponse := responses.CreateCurrencyResponse{}
	currencyResponse.ID = currencyData.ID
	currencyResponse.Name = currencyData.Name
	response := responses.NewCurrencyAcceptedResponse(currencyResponse)
	render.Status(request, http.StatusAccepted)
	render.Render(writer, request, response)
}
