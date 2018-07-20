package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate
	if _, err := govalidator.ValidateStruct(&brand); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create brand
	if err := h.brandService.CreateBrand(&brand); err != nil {
		switch err {
		case handmedown.ErrBrandAlreadyExists:
			http.Error(w, err.Error(), http.StatusConflict)
			return

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// respond
	response, _ := json.Marshal(&brand)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *brandHandler) handleGetBrands(w http.ResponseWriter, r *http.Request) {
	brands, err := h.brandService.FindBrands()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(struct {
		Brands []*handmedown.Brand `json:"brands"`
	}{
		Brands: brands,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
