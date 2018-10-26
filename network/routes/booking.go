package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Booking struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func BookingRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Put("/", CreateBooking)
	return router
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Booking successfully"
	render.JSON(w, r, response)
}
