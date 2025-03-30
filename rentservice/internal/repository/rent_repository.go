package repository

import (
	"database/sql"
	"fmt"
	"rentservice/internal/core/entity"
)

type RentRepository struct {
	db *sql.DB
}

func NewRentRepository(db *sql.DB) *RentRepository {
	return &RentRepository{db: db}
}

func (r *RentRepository) GetRentsByLandlord(landlordID int64) ([]*entity.Rent, error) {
	query := `SELECT id, status, cost, date_start, date_end, ad_id, landlord_id, renter_id 
			  FROM rents WHERE landlord_id = $1`

	rows, err := r.db.Query(query, landlordID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аренд по арендодателю: %v", err)
	}
	defer rows.Close()

	var rents []*entity.Rent
	for rows.Next() {
		rent := &entity.Rent{}
		err := rows.Scan(&rent.ID, &rent.Status, &rent.Cost, &rent.DateStart, &rent.DateEnd, &rent.AdID, &rent.LandlordID, &rent.RenterID)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		rents = append(rents, rent)
	}

	return rents, nil
}
