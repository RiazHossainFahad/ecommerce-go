package db

type Product struct {
	ID          int     `json:"id"` // tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var productList []Product

func init() {
	StoreProduct(Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is yellow",
		Price:       300,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-orange-640x480-orange.jpg",
	})

	StoreProduct(Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red",
		Price:       400,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-apple-640x480-apple.jpg",
	})

	StoreProduct(Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       60,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-banana-640x480-banana.png",
	})

	StoreProduct(Product{
		ID:          4,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       600,
		ImgUrl:      "https://shegrowsveg.com/wp-content/uploads/2024/11/Pomegranate-2-2048x2048-1.jpg",
	})

	StoreProduct(Product{
		ID:          5,
		Title:       "Lemon",
		Description: "Lemon is yellow",
		Price:       200,
		ImgUrl:      "https://cdn.britannica.com/84/188484-050-F27B0049/lemons-tree.jpg?w=300",
	})
}

func GetProductList() []Product {
	return productList
}

func StoreProduct(product Product) {
	productList = append(productList, product)
}

func GetProduct(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}

	return nil
}

func UpdateProduct(id int, product Product) (bool, string) {
	for i := 0; i < len(productList); i++ {
		if productList[i].ID == id {
			productList[i] = product
			return true, "Successfully updated"
		}

	}

	return false, "Not Found"
}

func DeleteProduct(id int) (bool, string) {
	tmpProductList := make([]Product, 0)
	isDeleted := false

	for i := 0; i < len(productList); i++ {
		if productList[i].ID != id {
			tmpProductList = append(tmpProductList, productList[i])
		} else {
			isDeleted = true
		}
	}

	if !isDeleted {
		return false, "Not Found"
	}

	productList = tmpProductList
	return true, "Successfully Delete"
}
