package repository

import "store/pkg/domain"

func AddProduct(info domain.Store) error {
	return DB.Create(&info).Error
}
