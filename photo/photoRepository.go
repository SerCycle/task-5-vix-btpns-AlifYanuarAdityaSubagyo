package photo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Photo, error)
	FindByID(ID int) (Photo, error)
	Create(user Photo) (Photo, error)
	Update(user Photo) (Photo, error)
	Delete(user Photo) (Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Photo, error) {
	var photos []Photo

	err := r.db.Find(&photos).Error

	return photos, err
}

func (r *repository) FindByID(ID int) (Photo, error) {
	var photo Photo

	err := r.db.Find(&photo, ID).Error

	return photo, err
}

func (r *repository) Create(photo Photo) (Photo, error) {
	err := r.db.Create(&photo).Error

	return photo, err
}

func (r *repository) Update(photo Photo) (Photo, error) {
	err := r.db.Save(&photo).Error

	return photo, err
}

func (r *repository) Delete(photo Photo) (Photo, error) {
	err := r.db.Delete(&photo).Error

	return photo, err
}
