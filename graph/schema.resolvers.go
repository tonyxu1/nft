package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"nft/db"
	"nft/etherscanapi"
	"nft/graph/generated"
	"nft/graph/model"
)

// GetEvents is the resolver for the getEvents field.
func (r *queryResolver) GetEvents(ctx context.Context, address *string, fromBlock *string, toBlock *string) (*model.Event, error) {
	token, err := db.GetTokenInfo()
	if err != nil {
		return nil, err
	}
	e := etherscanapi.GetEvents(token.LastReadBlock)
	return &e, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
