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

// ConversionRateHandler serve HTTP routes for conversion rate
type ConversionRateHandler struct {
	conversionRateService services.ConversionRateServiceContract
}

// NewConversionRateHandler returns new ConversionRateHandler
func NewConversionRateHandler(conversionRateService services.ConversionRateServiceContract) *ConversionRateHandler {
	return &ConversionRateHandler{
		conversionRateService: conversionRateService,
	}
}

// GetRoutes mount routes
func (handler *ConversionRateHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handler.Get)
	router.Post("/", handler.Store)
	router.Post("/convert", handler.Convert)

	return router
}

// Get :nodoc
func (handler *ConversionRateHandler) Get(writer http.ResponseWriter, request *http.Request) {
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

	conversionRateList, err := handler.conversionRateService.Get(limit, offset)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	response := responses.NewConversionRateListResponse(conversionRateList, limit, offset)
	render.Render(writer, request, response)
}

// Store :nodoc
func (handler *ConversionRateHandler) Store(writer http.ResponseWriter, request *http.Request) {
	var requestData requests.CreateConversionRateRequest
	if err := render.Bind(request, &requestData); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(err.Error()))
		return
	}
	conversionRate := models.ConversionRate{}
	conversionRate.CurrencyFromID = requestData.CurrencyFromID
	conversionRate.CurrencyToID = requestData.CurrencyToID
	conversionRate.Rate = requestData.Rate
	conversionRateData := models.ConversionRate{}
	conversionRateData, err := handler.conversionRateService.Create(conversionRate)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	conversionRateResponse := models.ConversionRate{}
	conversionRateResponse.ID = conversionRateData.ID
	conversionRateResponse.CurrencyFromID = conversionRateData.CurrencyFromID
	conversionRateResponse.CurrencyToID = conversionRateData.CurrencyToID
	conversionRateResponse.Rate = conversionRateData.Rate
	response := responses.NewConversionRateAcceptedResponse(conversionRateResponse)
	render.Status(request, http.StatusAccepted)
	render.Render(writer, request, response)
}

// Convert :nodoc
func (handler *ConversionRateHandler) Convert(writer http.ResponseWriter, request *http.Request) {
	var requestData requests.ConvertConversionRateRequest
	if err := render.Bind(request, &requestData); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(err.Error()))
		return
	}
	result, err := handler.conversionRateService.Convert(requestData.CurrencyFromID, requestData.CurrencyToID, requestData.Amount)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	response := responses.NewConvertResponse(result)
	render.Status(request, http.StatusAccepted)
	render.Render(writer, request, response)
}
