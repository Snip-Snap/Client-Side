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

func (r *mutationResolver) SignupClient(ctx context.Context,
	input NewClient) (*Client, error) {
	statement, err := db.Prepare("insert into client, (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")
	CheckError(err)

	//hash password
	//check inputs: phone# (unique)
	pw := string(input.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	CheckError(err)

	hashinputpass := string(hash)
	sum, err := statement.Exec(input.FullName,
		input.Gender, input.PhoneNumber, input.UserName, hashinputpass)
	CheckError(err)
	// idVal, err := sum.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }
	// string(0) is placeholder
	client := &Client{string(0), input.UserName,
		hashinputpass, input.FullName, input.Gender, input.PhoneNumber}
	fmt.Printf("%v+", sum)

	return client, nil

}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Clients(ctx context.Context) ([]*Client, error) {
	rows, err := db.Query("Select * from client")
	CheckError(err)

	clients := []*Client{}

	for rows.Next() {
		// if outside for loop, it would just rewrite to last client
		client := &Client{}
		err := rows.Scan(&client.ClientID, &client.FullName, &client.Gender,
			&client.PhoneNumber, &client.UserName, &client.Password)
		CheckError(err)
		fmt.Printf("%+v", client)
		clients = append(clients, client)
	}

	return clients, nil
}
