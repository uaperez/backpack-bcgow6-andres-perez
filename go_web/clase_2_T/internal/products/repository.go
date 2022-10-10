package products

type Product struct {
	Id    int
	Name  string
	Type  string
	Price float32
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, productType string, count int, price float64) (Product, error)
	LastId() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return []Product{}, nil
}

func (r *repository) Store(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}

func (r *repository) LastId() (int, error) {
	return 1, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return Product{}, nil
}
