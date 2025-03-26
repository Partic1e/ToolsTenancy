package repository

import (
	"adsservice/internal/core/entity"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

type AdRepository struct {
	db *sql.DB
}

func NewAdRepository(db *sql.DB) *AdRepository {
	return &AdRepository{db: db}
}

func (r *AdRepository) CreateAd(ad *entity.Ad) (*entity.Ad, error) {
	var id int64
	err := r.db.QueryRow(`
		INSERT INTO ads (name, description, cost_per_day, deposit, photo_path, landlord_id, category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		ad.Name, ad.Description, ad.CostPerDay, ad.Deposit, ad.PhotoPath, ad.LandlordId, ad.CategoryId).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("ошибка при создании объявления: %v", err)
	}

	ad.ID = id
	return ad, nil
}

func (r *AdRepository) DeleteAd(name string, landlordId int64) error {
	result, err := r.db.Exec(`DELETE FROM ads WHERE name = $1 AND landlord_id = $2`, name, landlordId)
	if err != nil {
		return fmt.Errorf("ошибка при удалении объявления: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при проверке удаления объявления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("объявление не найдено")
	}

	return nil
}

func (r *AdRepository) UpdateAd(name, description string, costPerDay, deposit decimal.Decimal, photoPath string, id, landlordId, categoryId int64) error {
	query := `UPDATE ads SET name = $1, description = $2, cost_per_day = $3, deposit = $4, photo_path = $5, landlord_id =$6, category_id = $7 WHERE id = $8`

	_, err := r.db.Exec(query, name, description, costPerDay, deposit, photoPath, landlordId, categoryId, id)
	if err != nil {
		return fmt.Errorf("не удалось обновить объявление в БД: %v", err)
	}

	return nil
}

func (r *AdRepository) GetAllCategories() ([]entity.Category, error) {
	rows, err := r.db.Query("SELECT id, name FROM category")
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении категорий: %v", err)
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
