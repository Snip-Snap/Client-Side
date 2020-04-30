package directive

import (
	"api/auth"
	"api/generated"
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func AddAuth(c *generated.Config) {

	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		username := auth.ForContext(ctx)
		fmt.Println(ctx)
		if username != "" {
			return next(ctx)
		} else {
			return nil, errors.New("Unauthorised")
		}
	}
}
