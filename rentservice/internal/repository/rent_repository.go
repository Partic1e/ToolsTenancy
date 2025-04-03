package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rentservice/internal/core/entity"
	"time"
)

type RentRepository struct {
	db *sql.DB
}

func NewRentRepository(db *sql.DB) *RentRepository {
	return &RentRepository{db: db}
}

func (r *RentRepository) GetRentsByLandlord(landlordID int64) ([]*entity.Rent, error) {
	query := `SELECT id, status, cost, date_start, date_end, ad_id, landlord_id, renter_id, held_id
			  FROM rents WHERE landlord_id = $1`

	rows, err := r.db.Query(query, landlordID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аренд по арендодателю: %v", err)
	}
	defer rows.Close()

	var rents []*entity.Rent
	for rows.Next() {
		rent := &entity.Rent{}
		err := rows.Scan(&rent.ID, &rent.Status, &rent.Cost, &rent.DateStart, &rent.DateEnd, &rent.AdID, &rent.LandlordID, &rent.RenterID, &rent.HeldID)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		rents = append(rents, rent)
	}

	return rents, nil
}

func (r *RentRepository) GetRentsByRenter(renterID int64) ([]*entity.Rent, error) {
	query := `SELECT id, status, cost, date_start, date_end, ad_id, landlord_id, renter_id, held_id 
			  FROM rents WHERE renter_id = $1`

	rows, err := r.db.Query(query, renterID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аренд по арендателю: %v", err)
	}
	defer rows.Close()

	var rents []*entity.Rent
	for rows.Next() {
		rent := &entity.Rent{}
		err := rows.Scan(&rent.ID, &rent.Status, &rent.Cost, &rent.DateStart, &rent.DateEnd, &rent.AdID, &rent.LandlordID, &rent.RenterID, &rent.HeldID)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		rents = append(rents, rent)
	}

	return rents, nil
}

func (r *RentRepository) GetRentedDates(adID int64) ([]time.Time, error) {
	query := `SELECT date_start, date_end
			  FROM rents
       		  WHERE ad_id = $1 AND date_end >= CURRENT_DATE;`

	rows, err := r.db.Query(query, adID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentedDates []time.Time

	for rows.Next() {
		var startDate, endDate time.Time
		if err := rows.Scan(&startDate, &endDate); err != nil {
			return nil, err
		}

		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			if d.After(time.Now()) {
				rentedDates = append(rentedDates, d)
			}
		}
	}

	return rentedDates, nil
}

func (r *RentRepository) CreateRent(rent entity.Rent) error {
	query := `
		INSERT INTO rents (status, cost, date_start, date_end, ad_id, landlord_id, renter_id, held_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(query,
		rent.Status,
		rent.Cost,
		rent.DateStart,
		rent.DateEnd,
		rent.AdID,
		rent.LandlordID,
		rent.RenterID,
		rent.HeldID,
	)
	if err != nil {
		return fmt.Errorf("ошибка при создании аренды: %v", err)
	}

	return nil
}

func (r *RentRepository) CompleteRent(rentID int64) error {
	_, err := r.db.Exec("UPDATE Rents SET status = 'complete' WHERE id = ?", rentID)
	if err != nil {
		return errors.New("failed to update rent status to 'complete'")
	}
	return nil
}
