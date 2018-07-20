package http

import (
	"net"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"

	"github.com/danscotton/handmedown"
)

type Server struct {
	ln net.Listener

	// services
	BrandService handmedown.BrandService

	// server config
	Addr string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	s.ln = ln
	go http.Serve(s.ln, s.router())

	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) URL() url.URL {
	if s.ln == nil {
		return url.URL{}
	}

	return url.URL{
		Scheme: "http",
		Host:   s.ln.Addr().String(),
	}
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Mount("/brands", s.brandHandler())
	})

	return r
}

func (s *Server) brandHandler() *brandHandler {
	h := newBrandHandler()
	h.brandService = s.BrandService
	return h
}
