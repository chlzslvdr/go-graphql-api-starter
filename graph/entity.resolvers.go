package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/go-graphl-server-lookup-api/graph/generated"
	models1 "github.com/go-graphl-server-lookup-api/models"
)

func (r *entityResolver) FindLookupsByID(ctx context.Context, id string) (*models1.Lookups, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated1.EntityResolver implementation.
func (r *Resolver) Entity() generated1.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
