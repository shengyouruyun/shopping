package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/shengyouruyun/hackernews/graph/generated"
	"github.com/shengyouruyun/hackernews/graph/model"
)


func (r *mutationResolver) CreateShopItem(ctx context.Context, input model.NewShopItem) (*model.ShopItem, error) {
	var shopitem model.ShopItem
	shopitem.ID = "1"
	shopitem.Name = input.Name
	shopitem.Description = input.Description
	shopitem.ImageURL = input.ImageURL
	shopitem.Price = input.Price
	return &shopitem, nil
}


func (r *mutationResolver) RemoveShopItem(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	var order model.Order
	//var newshopItems []model.NewShopItem
	var shopItems []*model.ShopItem

	order.Name = input.Name

	order.Address = input.Address
	order.Phone = input.Phone

	for _, newshopitem := range input.OrderedItems {

		shopItems = append(shopItems, &model.ShopItem{
			Name:        newshopitem.Name,
			Description: newshopitem.Description,
			Price:       newshopitem.Price,
			ImageURL:    newshopitem.ImageURL,
		})

		order.OrderedItems = shopItems

	}

	return &order, nil
}

func (r *mutationResolver) RemoveOrder(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ShopItems(ctx context.Context) ([]*model.ShopItem, error) {
	var shopItems []*model.ShopItem
	//var shopitem model.ShopItem
	shopitem := model.ShopItem{
		ID: "1", Name: "book",
		Description: "book for working with time",
		Price:       23.90,
		ImageURL:    "http://wwww.bing.com"}

	shopItems = append(shopItems, &shopitem)

	return shopItems, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }




















