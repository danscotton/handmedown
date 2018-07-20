package http

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/danscotton/handmedown"
	"github.com/go-chi/chi"
)

type brandHandler struct {
	router chi.Router

	brandService handmedown.BrandService
}

func newBrandHandler() *brandHandler {
	h := &brandHandler{router: chi.NewRouter()}
	h.router.Post("/", h.handlePostBrand)
	h.router.Get("/", h.handleGetBrands)
	return h
}

func (h *brandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *brandHandler) handlePostBrand(w http.ResponseWriter, r *http.Request) {
	brand := handmedown.Brand{}

	// decode body
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		InvalidRequestError(w)
		return
	}

	// validate
	if err := brand.Validate(); err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate id
	brand.ID = bson.NewObjectId()

	// create brand
	if err := h.brandService.CreateBrand(&brand); err != nil {
		switch err {
		case handmedown.ErrBrandAlreadyExists:
			Error(w, err.Error(), http.StatusConflict)
			return

		default:
			Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// respond
	Respond(w).With(http.StatusCreated, &brand)
}

func (h *brandHandler) handleGetBrands(w http.ResponseWriter, r *http.Request) {
	brands, err := h.brandService.FindBrands()
	if err != nil {
		Error(w, "unable to retrieve brands", http.StatusInternalServerError)
		return
	}

	Respond(w).With(http.StatusOK, brandsResponse{brands})
}

type brandsResponse struct {
	Brands []*handmedown.Brand `json:"brands"`
}
