package db

type Product struct {
	ID          int     `json:"id"` // tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var ProductList []Product

func init() {
	ProductList = append(ProductList, Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is yellow",
		Price:       300,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-orange-640x480-orange.jpg",
	})

	ProductList = append(ProductList, Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red",
		Price:       400,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-apple-640x480-apple.jpg",
	})

	ProductList = append(ProductList, Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       60,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-banana-640x480-banana.png",
	})

	ProductList = append(ProductList, Product{
		ID:          4,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       600,
		ImgUrl:      "https://shegrowsveg.com/wp-content/uploads/2024/11/Pomegranate-2-2048x2048-1.jpg",
	})

	ProductList = append(ProductList, Product{
		ID:          5,
		Title:       "Lemon",
		Description: "Lemon is yellow",
		Price:       200,
		ImgUrl:      "https://cdn.britannica.com/84/188484-050-F27B0049/lemons-tree.jpg?w=300",
	})
}
