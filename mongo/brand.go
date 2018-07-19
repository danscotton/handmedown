package mongo

import (
	"github.com/danscotton/handmedown"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type BrandService struct {
	db *DB
}

func NewBrandService(db *DB) (*BrandService, error) {
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	if err := db.session.DB("handmedown").C("brand").EnsureIndex(index); err != nil {
		return nil, err
	}

	return &BrandService{db: db}, nil
}

func (bs *BrandService) CreateBrand(brand *handmedown.Brand) error {
	if err := bs.db.session.DB("handmedown").C("brand").Insert(&brand); err != nil {
		if mgo.IsDup(err) {
			return handmedown.ErrBrandAlreadyExists
		}
		return err
	}

	return nil
}

func (bs *BrandService) FindBrands() ([]*handmedown.Brand, error) {
	var brands []*handmedown.Brand

	if err := bs.db.session.DB("handmedown").C("brand").Find(bson.M{}).All(&brands); err != nil {
		return nil, err
	}

	return brands, nil
}
