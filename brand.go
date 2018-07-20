package handmedown

import (
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
)

const (
	ErrBrandAlreadyExists = Error("brand already exists")
)

type Brand struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name" valid:"alphanum,required~name is required"`
}

func (b *Brand) Validate() error {
	if _, err := govalidator.ValidateStruct(b); err != nil {
		return err
	}
	return nil
}

type BrandService interface {
	CreateBrand(brand *Brand) error
	FindBrands() ([]*Brand, error)
}
