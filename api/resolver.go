package graphqltest

//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignupClient(ctx context.Context, input NewClient) (*Client, error) {
	//input.FullName
	statement, err := db.Prepare("insert into client (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")
	if err != nil {
		return nil, err
	}
	//hash password
	//check inputs: phone# (unique),
	pw := string(input.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashinputpass := string(hash)
	sum, err := statement.Exec(input.FullName, input.Gender, input.PhoneNumber, input.UserName, hashinputpass)
	if err != nil {
		return nil, err
	}
	// idVal, err := sum.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }
	client := &Client{string(0), input.UserName, hashinputpass, input.FullName, input.Gender, input.PhoneNumber}
	fmt.Printf("%v+", sum)

	return client, nil

}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Clients(ctx context.Context) ([]*Client, error) {
	rows, err := db.Query("Select * from client")
	if err != nil {
		return nil, err
	}

	clients := []*Client{}

	for rows.Next() {
		// if outside for loop, it would just rewrite to last client
		client := &Client{}
		err := rows.Scan(&client.ClientID, &client.FullName, &client.Gender,
			&client.PhoneNumber, &client.UserName, &client.Password)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v", client)
		clients = append(clients, client)
	}

	return clients, nil
}
