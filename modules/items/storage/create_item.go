package storage

import (
	"context"
	"todo-go/common"
	"todo-go/modules/items/model"
)

func (s *SqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
