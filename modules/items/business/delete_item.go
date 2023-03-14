package business

import (
	"context"
	"todo-go/common"
	"todo-go/modules/items/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type DeleteItemBusiness struct {
	store DeleteItemStorage
}

func NewDeleteItemBusiness(store DeleteItemStorage) *DeleteItemBusiness {
	return &DeleteItemBusiness{store: store}
}

func (business DeleteItemBusiness) DeleteItemById(ctx context.Context, id int) error {

	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrModelDeleted
	}

	if err := business.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}

		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
