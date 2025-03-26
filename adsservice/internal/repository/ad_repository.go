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
