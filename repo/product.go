package repo

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	EmptyProduct() Product

	List() ([]*Product, error)
	Store(Product) (*Product, error)
	Get(int) (*Product, error)
	Update(int, Product) (*Product, error)
	Delete(int) error
	Exists(int) (bool, error)
}

type Product struct {
	ID          int        `json:"id" db:"id"` // tag
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Price       float64    `json:"price" db:"price"`
	ImgUrl      string     `json:"img_url" db:"img_url"`
	CreatedAt   *time.Time `json:"-" db:"created_at"`
	UpdatedAt   *time.Time `json:"-" db:"updated_at"`
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}

	// generateInitialProduct(repo)

	return repo

}

func (r *productRepo) EmptyProduct() Product {
	return Product{}
}

func (r *productRepo) Exists(id int) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 
				1
			FROM products 
			WHERE
				id=$1
		)
	`
	err := r.db.Get(&exists, query, id)

	return exists, err
}

func (r *productRepo) Store(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			img_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) 
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgUrl,
	).Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product

	query := `
		SELECT 
			*
		FROM products
		ORDER BY
			id
	`

	err := r.db.Select(&products, query)
	return products, err
}

func (r *productRepo) Get(id int) (*Product, error) {
	var product Product

	query := `
		SELECT 
			*
		FROM products
		WHERE
			id=$1
	`

	err := r.db.Get(&product, query, id)
	return &product, err
}

func (r *productRepo) Update(id int, p Product) (*Product, error) {

	exists, err := r.Exists(id)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("No product found")
	}

	query := `
		UPDATE products
		SET 
			title=$1,
			description=$2,
			price=$3,
			img_url=$4,
			updated_at=NOW()
		WHERE 
			id=$5
	`
	_, err = r.db.Exec(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgUrl,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &p, err
}

func (r *productRepo) Delete(id int) error {
	exists, err := r.Exists(id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("No product found")
	}

	query := `
		DELETE FROM products
		WHERE
			id=$1
	`

	_, err = r.db.Exec(query, id)
	return err
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
