package repository

import "adsservice/internal/core/entity"

type AdRepositoryInterface interface {
	CreateAd(ad *entity.Ad) (*entity.Ad, error)
	UpdateAd(ad *entity.Ad) (*entity.Ad, error)
	DeleteAd(id int64, landlordId int64) error
}
