package directive

import (
	"api/auth"
	"api/generated"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

func AddAuth(c *generated.Config) {

	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		username := auth.ForContext(ctx)
		if username != "" {
			return next(ctx)
		} else {
			return nil, errors.New("Unauthorised")
		}
	}
}
