package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/library/config"
	"github.com/library/modules/address/domain"
	"github.com/library/modules/address/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

type MySQLRepo struct {
	db *gorm.DB
}

func NewMySQLRepo() *MySQLRepo {
	db := getDBInstance()
	return &MySQLRepo{db: db}
}

func (r *MySQLRepo) Create(entity *domain.Entity) (*domain.Entity, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *MySQLRepo) Delete(id string) error {
	if err := r.db.Delete(&domain.Entity{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepo) Get(id string) (*domain.Entity, error) {
	var entity domain.Entity
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *MySQLRepo) List(filter dto.FilterRequestDTO) ([]*domain.Entity, error) {
	var entityList []*domain.Entity

	const defaultPage = 1
	const defaultPageSize = 10

	if filter.Page < 1 {
		filter.Page = defaultPage
	}
	if filter.PageSize < 1 {
		filter.PageSize = defaultPageSize
	}

	offset := (filter.Page - 1) * filter.PageSize

	query := r.db.Limit(filter.PageSize).Offset(offset)
	if filter.Owner != "" {
		query = query.Where("owner = ?", filter.Owner)
	}

	if err := query.Find(&entityList).Error; err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	return entityList, nil
}

func (r *MySQLRepo) Update(entity *domain.Entity) (*domain.Entity, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func getDBInstance() *gorm.DB {
	once.Do(func() {
		cfg := config.GetConfig()

		user := cfg.DBUser
		password := cfg.DBPassword
		host := cfg.DBHost
		port := cfg.DBPort
		dbName := cfg.DBName

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)

		var err error
		dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		if err = dbInstance.AutoMigrate(&domain.Entity{}); err != nil {
			log.Fatalf("Error running auto migration: %v", err)
		}
	})

	return dbInstance
}
