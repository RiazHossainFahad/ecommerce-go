package repo

import (
	"ecommerce/domain"
	"ecommerce/product"
	"errors"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
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

func (r *productRepo) EmptyProduct() domain.Product {
	return domain.Product{}
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

func (r *productRepo) Unique(title string, id int) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM products 
			WHERE title = $1
	`

	args := []interface{}{title}

	// If id > 0, exclude that record from the check
	if id > 0 {
		query += ` AND id != $2`
		args = append(args, id)
	}

	query += `)`

	err := r.db.Get(&exists, query, args...)

	return exists, err
}

func (r *productRepo) Store(p domain.Product) (*domain.Product, error) {
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

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {
	var products []*domain.Product

	query := `
		SELECT 
			*
		FROM products
		LIMIT $1
		OFFSET $2
	`

	offset := (page - 1) * limit
	err := r.db.Select(&products, query, limit, offset)
	return products, err
}

func (r *productRepo) Count() (int, error) {
	var count int

	query := `
		SELECT 
			COUNT(id)
		FROM products
	`

	err := r.db.Get(&count, query)
	return count, err
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var product domain.Product

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

func (r *productRepo) Update(id int, p domain.Product) (*domain.Product, error) {

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
	r.Store(domain.Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is yellow",
		Price:       300,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-orange-640x480-orange.jpg",
	})

	r.Store(domain.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red",
		Price:       400,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-apple-640x480-apple.jpg",
	})

	r.Store(domain.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       60,
		ImgUrl:      "https://www.quanta.org/thumbs/thumb-banana-640x480-banana.png",
	})

	r.Store(domain.Product{
		ID:          4,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       600,
		ImgUrl:      "https://shegrowsveg.com/wp-content/uploads/2024/11/Pomegranate-2-2048x2048-1.jpg",
	})

	r.Store(domain.Product{
		ID:          5,
		Title:       "Lemon",
		Description: "Lemon is yellow",
		Price:       200,
		ImgUrl:      "https://cdn.britannica.com/84/188484-050-F27B0049/lemons-tree.jpg?w=300",
	})
}
