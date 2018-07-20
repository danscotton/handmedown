package http

import (
	"encoding/json"
	"net/http"

	"github.com/danscotton/handmedown"
	"github.com/go-chi/chi"
)

type brandHandler struct {
	router chi.Router

	brandService handmedown.BrandService
}

func newBrandHandler() *brandHandler {
	h := &brandHandler{router: chi.NewRouter()}
	h.router.Get("/", h.handleGetBrands)
	return h
}

func (h *brandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
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
