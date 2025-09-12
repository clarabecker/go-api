package usecases

import "github.com/clarabecker/estudos-go/Internal/entity"

type CreateProductInputDTO struct {
	Name string 
	Price float64
}

type CreateProductOutputDto struct {
	ID string
	Name string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDTO) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDto{
		ID: product.ID, 
		Name: product.Name, 
		Price: product.Price,
	}, nil 
}