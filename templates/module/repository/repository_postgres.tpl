package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/library/config"
	"github.com/library/modules/{{ toSnake .Title }}/domain"
	"github.com/library/modules/{{ toSnake .Title }}/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	pgDBInstance *gorm.DB
	pgOnce       sync.Once
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo() *PostgresRepo {
	db := getPGDBInstance()
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Create(entity *domain.Entity) (*domain.Entity, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *PostgresRepo) Delete(id string) error {
	if err := r.db.Delete(&domain.Entity{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepo) Get(id string) (*domain.Entity, error) {
	var entity domain.Entity
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *PostgresRepo) List(filter dto.FilterRequestDTO) ([]*domain.Entity, error) {
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

func (r *PostgresRepo) Update(entity *domain.Entity) (*domain.Entity, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// Função específica para Postgres
func getPGDBInstance() *gorm.DB {
	pgOnce.Do(func() {
		cfg := config.GetConfig()

		user := cfg.DBUser
		password := cfg.DBPassword
		host := cfg.DBHost
		port := cfg.DBPort
		dbName := cfg.DBName

		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s",
			host, port, user, password, dbName,
		)

		var err error
		pgDBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		if err = pgDBInstance.AutoMigrate(&domain.Entity{}); err != nil {
			log.Fatalf("Error running auto migration: %v", err)
		}
	})

	return pgDBInstance
}
