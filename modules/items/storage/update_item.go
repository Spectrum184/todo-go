package storage

import (
	"context"
	"todo-go/common"
	"todo-go/modules/items/model"
)

func (s *SqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {

		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
