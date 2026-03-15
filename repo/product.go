package repo

import "errors"

type ProductRepo interface {
	EmptyProduct() Product

	List() ([]*Product, error)
	Store(Product) (*Product, error)
	Get(int) (*Product, error)
	Update(int, Product) (*Product, error)
	Delete(int) error
}

type Product struct {
	ID          int     `json:"id"` // tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProduct(repo)

	return repo

}

func (r *productRepo) EmptyProduct() Product {
	return Product{}
}

func (r *productRepo) Store(p Product) (*Product, error) {
	r.productList = append(r.productList, &p)

	return &p, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, errors.New("No product found")
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {
	for i := 0; i < len(r.productList); i++ {
		if r.productList[i].ID == id {
			r.productList[i] = &p
			return r.productList[i], nil
		}

	}

	return nil, errors.New("No product found")
}

func (r *productRepo) Delete(id int) error {
	tmpProductList := make([]*Product, 0)
	isDeleted := false

	for i := 0; i < len(r.productList); i++ {
		if r.productList[i].ID != id {
			tmpProductList = append(tmpProductList, r.productList[i])
		} else {
			isDeleted = true
		}
	}

	if !isDeleted {
		return errors.New("No product found")
	}

	r.productList = tmpProductList
	return nil
}

func generateInitialProduct(r *productRepo) {
	r.Store(Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is yellow",
		Price:       300,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-orange-640x480-orange.jpg",
	})

	r.Store(Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red",
		Price:       400,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-apple-640x480-apple.jpg",
	})

	r.Store(Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       60,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-banana-640x480-banana.png",
	})

	r.Store(Product{
		ID:          4,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       600,
		ImgUrl:      "https://shegrowsveg.com/wp-content/uploads/2024/11/Pomegranate-2-2048x2048-1.jpg",
	})

	r.Store(Product{
		ID:          5,
		Title:       "Lemon",
		Description: "Lemon is yellow",
		Price:       200,
		ImgUrl:      "https://cdn.britannica.com/84/188484-050-F27B0049/lemons-tree.jpg?w=300",
	})
}
