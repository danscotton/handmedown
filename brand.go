package handmedown

const (
	ErrBrandAlreadyExists = Error("brand already exists")
)

type Brand struct {
	Name string `bson:"name" json:"name"`
}

type BrandService interface {
	CreateBrand(brand *Brand) error
	FindBrands() ([]*Brand, error)
}
