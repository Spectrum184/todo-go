package storage

import (
	"context"
	"gorm.io/gorm"
	"todo-go/common"
	"todo-go/modules/items/model"
)

func (s *SqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := model.ItemStatusDeleted

	if err := s.db.Table(model.TodoItem{}.
		TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus.String(),
		}).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}

	return nil
}
