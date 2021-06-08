package requests

import (
	"errors"
	"net/http"
)

// CreateCurrencyRequest holds data for currency creation
type CreateCurrencyRequest struct {
	Name string `json:"name"`
}

// Bind :nodoc:
func (req CreateCurrencyRequest) Bind(request *http.Request) error {
	if req.Name == "" {
		return errors.New("`name` is required")
	}
	return nil
}
