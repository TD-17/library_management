package repository

import (
	"gorm.io/gorm"

	model "github.com/TD-17/library_management/internal/models"
)

type LibraryRepository struct {
	DB *gorm.DB
}

// NewLibraryRepository initializes a new Library repository
func NewLibraryRepository(db *gorm.DB) *LibraryRepository {
	return &LibraryRepository{DB: db}
}

// Create a new library
func (repo *LibraryRepository) CreateLibrary(library *model.Library) error {
	return repo.DB.Create(library).Error
}

// Get a library by ID
func (repo *LibraryRepository) GetLibraryByID(id int) (*model.Library, error) {
	var library model.Library
	err := repo.DB.First(&library, id).Error
	if err != nil {
		return nil, err
	}
	return &library, nil
}

// Update an existing library
func (repo *LibraryRepository) UpdateLibrary(library *model.Library) error {
	return repo.DB.Save(library).Error
}

// Delete a library by ID
func (repo *LibraryRepository) DeleteLibrary(id int) error {
	return repo.DB.Delete(&model.Library{}, id).Error
}

// Get all libraries
func (repo *LibraryRepository) GetAllLibraries() ([]model.Library, error) {
	var libraries []model.Library
	err := repo.DB.Find(&libraries).Error
	if err != nil {
		return nil, err
	}
	return libraries, nil
}
