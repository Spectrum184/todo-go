package business

import (
	"context"
	"todo-go/common"
	"todo-go/modules/items/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type GetItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *GetItemBusiness {
	return &GetItemBusiness{store: store}
}

func (business GetItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
