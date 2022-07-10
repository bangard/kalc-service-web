package repository

import (
	"context"
	"dobledcloud.com/consumers/models"
)

type Repository interface {
	GetFilesPublishedByEmission(ctx context.Context, key int) ([]*models.Publishes, error)
	GetEmissionByKey(ctx context.Context, key string) (*models.Emission, error)
	GetSecretForEmission(ctx context.Context, emission_id int, emission_client string) bool
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// GetFilesPublishedByEmission Return all files published by an emission
func GetFilesPublishedByEmission(ctx context.Context, key int) ([]*models.Publishes, error) {
	return implementation.GetFilesPublishedByEmission(ctx, key)
}

// GetEmissionByKey Return emission id from emission key
func GetEmissionByKey(ctx context.Context, key string) (*models.Emission, error) {
	return implementation.GetEmissionByKey(ctx, key)
}

// GetSecretForEmission Return key id from emission key
func GetSecretForEmission(ctx context.Context, emission_id int, emission_client string) bool {
	return implementation.GetSecretForEmission(ctx, emission_id, emission_client)
}

func Close() error {
	return implementation.Close()
}
