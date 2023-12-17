package app

import (
	"context"
	"errors"

	"template-manager/dto"
	"template-manager/entity"
)

func (a *App) CreateApiKey(ctx context.Context, req dto.CreateApiKeyRequest) error {
	var key = entity.Key{
		AccountID: req.AccountID,
		Name:      req.ApiKeyName,
	}

	if err := key.GenerateKey(); err != nil {
		return err
	}
	if err := a.db.Model(&key).Create(&key).Error; err != nil {
		a.logger.ErrorContext(ctx, "failed to create account %+v", err)
		return err
	}

	return nil
}

func (a *App) FindApiKeys(ctx context.Context, accountID string) ([]entity.Key, error) {
	var keys []entity.Key
	if err := a.db.Model(&entity.Key{}).Where("id = ?", accountID).Find(&keys).Error; err != nil {
		return nil, errors.New("account does not exist")
	}
	return keys, nil
}

func (a *App) DeleteApiKey(ctx context.Context, req dto.DeleteApiKeyRequest) error {
	var key = entity.Key{
		ID:        req.ApiKeyID,
		AccountID: req.AccountID,
	}
	if err := a.db.Model(&key).Where(&key).Delete(&key).Error; err != nil {
		return errors.New("account does not exist")
	}
	return nil
}
