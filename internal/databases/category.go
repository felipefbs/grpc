package databases

import (
	"database/sql"

	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *sql.DB
}

type CategoryModel struct {
	ID, Name, Description string
}

func NewCategory(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) Create(name, description string) (*CategoryModel, error) {
	id := uuid.NewString()

	_, err := repo.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description,
	)
	if err != nil {
		return nil, err
	}

	return &CategoryModel{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (repo *CategoryRepository) FindAll() ([]CategoryModel, error) {
	rows, err := repo.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []CategoryModel{}
	for rows.Next() {
		var id, name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, CategoryModel{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (repo *CategoryRepository) FindByCourseID(courseID string) (*CategoryModel, error) {
	var id, name, description string
	err := repo.db.QueryRow("SELECT ca.id, ca.name, ca.description FROM categories ca join courses co on co.category_id = ca.id where co.id = $1;", courseID).Scan(&id, &name, &description)
	if err != nil {
		return nil, err
	}

	return &CategoryModel{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
