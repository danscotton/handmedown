package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/danscotton/handmedown"
	"github.com/danscotton/handmedown/http"
	"github.com/danscotton/handmedown/mongo"
)

func main() {
	m := NewMain()

	// run app
	if err := m.Run(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	// Shutdown on SIGINT (CTRL-C).
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Fprintln(os.Stdout, "received interrupt, shutting down...")
}

type Main struct {
	Config Config
}

func (m *Main) Run() error {

	// set up database
	db := mongo.NewDB()
	db.Url = m.Config.Database.Url
	if err := db.Open(); err != nil {
		return err
	}
	// defer db.Close()
	fmt.Fprintf(os.Stdout, "database connected\n")

	// set up brand service
	brandService, err := mongo.NewBrandService(db)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "brand service running\n")

	// set up http server
	server := http.NewServer()
	server.Addr = m.Config.HTTP.Addr
	server.BrandService = brandService
	if err := server.Open(); err != nil {
		return err
	}
	u := server.URL()
	fmt.Fprintf(os.Stdout, "http server running on %s\n", u.String())

	// create example brand data
	brands := []string{"Gap", "Next", "Primark", "Tu"}
	for _, brand := range brands {
		if err := brandService.CreateBrand(&handmedown.Brand{Name: brand}); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}

	return nil
}

type Config struct {
	Database struct {
		Url string
	}
	HTTP struct {
		Addr string
	}
}

func NewMain() *Main {
	return &Main{
		Config: DefaultConfig(),
	}
}

func DefaultConfig() Config {
	var c Config
	c.Database.Url = "localhost"
	c.HTTP.Addr = "localhost:3010"
	return c
}
