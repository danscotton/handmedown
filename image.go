package handmedown

type ImageId int

type Image struct {
	ID     ImageId
	Width  int
	Height int
}

type ImageService interface {
	CreateImage(image *Image) (*Image, error)
	FindImageById(id ImageId) (*ImageId, error)
}
