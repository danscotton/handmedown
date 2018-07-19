package main

import (
	"fmt"
	"os"

	"github.com/danscotton/handmedown"
	"github.com/danscotton/handmedown/mongo"
)

func main() {
	m := NewMain()

	if err := m.Run(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

type Main struct {
	Config Config
}

func (m *Main) Run() error {

	db := mongo.NewDB()
	db.Url = m.Config.Database.Url
	if err := db.Open(); err != nil {
		return err
	}
	defer db.Close()
	fmt.Fprintf(os.Stdout, "database connected\n")

	brandService, err := mongo.NewBrandService(db)
	if err != nil {
		return err
	}

	if err := brandService.CreateBrand(&handmedown.Brand{Name: "Next"}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	brands, err := brandService.FindBrands()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%#v\n", brands)

	return nil
}

type Config struct {
	Database struct {
		Url string
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
	return c
}
