package reqmodel

import (
	"errors"
	"github.com/eexe1/supermaid-go/model"
	"net/http"
)
/**
	model for receiving requests
 */
type BookingRequest struct {
	*model.Booking
}

func (a *BookingRequest) Bind(r *http.Request) error {
	// a.Booking is nil if no Booking fields are sent in the request.
	// Return an error to avoid a nil pointer dereference.
	if a.Booking == nil {
		return errors.New("missing required booking fields")
	}
	return nil
}

type BookingResponse struct {
	*model.Booking

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}
// Create a response for a booking
func NewBookingResponse(booking *model.Booking) *BookingResponse {
	resp := &BookingResponse{Booking: booking}
	return resp
}

// Render booking response
func (rd *BookingResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}