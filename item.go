package handmedown

import "time"

type ItemId int

type Item struct {
	ID          ItemId
	Image       *Image
	Description string
	Gender      Gender
	Brand       *Brand
	Owner       *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemService interface {
	CreateItem(item *Item) (*Item, error)
	FindItemByID(id ItemId) (*Item, error)
	FindItemsByGender(gender Gender) ([]*Item, error)
	FindItemsByBrand(brand *Brand) ([]*Item, error)
	FindItemsByOwner(owner *User) ([]*Item, error)
}
