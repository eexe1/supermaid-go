package routes

import (
	"github.com/eexe1/supermaid-go/database"
	networkError "github.com/eexe1/supermaid-go/network/error"
	"github.com/eexe1/supermaid-go/network/reqmodel"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func BookingRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Put("/", CreateBooking)
	return router
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {

	data := &reqmodel.BookingRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, networkError.ErrInvalidRequest(err))
		return
	}

	booking := data.Booking
	database.InsertBooking(booking)
	//dbNewArticle(article)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, reqmodel.NewBookingResponse(booking))

	//response := make(map[string]string)
	//response["message"] = "Created Booking successfully"
	//render.JSON(w, r, response)
}
