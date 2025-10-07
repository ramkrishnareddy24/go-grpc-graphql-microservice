package main

import (
	"context"

	"github.com/99designs/gqlgen/integration/server"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Account(ctx context.Context, pagination *PaginationInput) (*Account, error) {
	
}

func (r *queryResolver) Product(ctx context.Context, query *string,id *string) ([]*Product, error) {
	
}

