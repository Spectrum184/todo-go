package business

import (
	"context"
	"todo-go/common"
	"todo-go/modules/items/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type ListItemBusiness struct {
	store ListItemStorage
}

func NewListItemBusiness(store ListItemStorage) *ListItemBusiness {
	return &ListItemBusiness{store: store}
}

func (business ListItemBusiness) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {

	data, err := business.store.ListItem(ctx, filter, paging)

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(model.EntityName, err)
		}

		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
