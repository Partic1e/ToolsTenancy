package repository

import (
	"adsservice/internal/core/entity"
	"database/sql"
	"fmt"
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

func (r *AdRepository) UpdateAd(ad *entity.Ad) error {
	query := `UPDATE ads SET name = $1, description = $2, cost_per_day = $3, deposit = $4, 
              photo_path = $5, landlord_id = $6, category_id = $7 WHERE id = $8`

	_, err := r.db.Exec(query,
		ad.Name,
		ad.Description,
		ad.CostPerDay,
		ad.Deposit,
		ad.PhotoPath,
		ad.LandlordId,
		ad.CategoryId,
		ad.ID)
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

func (r *AdRepository) GetAdsByCategory(categoryID int64) ([]*entity.Ad, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, cost_per_day, deposit, photo_path, landlord_id, category_id
		FROM ads WHERE category_id = $1`, categoryID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении объявлений: %v", err)
	}
	defer rows.Close()

	var ads []*entity.Ad
	for rows.Next() {
		var ad entity.Ad
		err := rows.Scan(&ad.ID,
			&ad.Name,
			&ad.Description,
			&ad.CostPerDay,
			&ad.Deposit,
			&ad.PhotoPath,
			&ad.LandlordId,
			&ad.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании объявлений: %v", err)
		}
		ads = append(ads, &ad)
	}
	return ads, nil
}

func (r *AdRepository) GetAdsByLandlord(landlordID int64) ([]*entity.Ad, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, cost_per_day, deposit, photo_path, landlord_id, category_id
		FROM ads WHERE landlord_id = $1`, landlordID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении объявлений: %v", err)
	}
	defer rows.Close()

	var ads []*entity.Ad
	for rows.Next() {
		var ad entity.Ad
		err := rows.Scan(&ad.ID, &ad.Name, &ad.Description, &ad.CostPerDay, &ad.Deposit, &ad.PhotoPath, &ad.LandlordId, &ad.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании объявлений: %v", err)
		}
		ads = append(ads, &ad)
	}
	return ads, nil
}
