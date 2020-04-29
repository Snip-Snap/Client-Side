package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/auth"
	"api/generated"
	"api/jwtoken"
	"api/model"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) SignupClient(ctx context.Context, input model.NewClient) (*model.Response, error) {
	statement, err := DB.Prepare("insert into client (firstname, lastname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5, $6)")
	CheckError(err)
	//TODO: Check inputs for uniqueness and appropriate characters.
	// pw := string(input.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),
		bcrypt.DefaultCost)
	CheckError(err)

	hashedInputpw := string(hash)
	_, err = statement.Exec(input.FirstName, input.LastName,
		input.Gender, input.PhoneNumber, input.UserName, hashedInputpw)
	CheckError(err)

	res := &model.Response{Error: "Okay"}

	return res, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Response, error) {
	stmt, err := DB.Prepare("select hashedpassword from client where username=$1")
	if dbError(err) {
		return nil, err
	}
	var cli_pass string
	err = stmt.QueryRow(input.UserName).Scan(&cli_pass)
	if dbError(err) {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(cli_pass), []byte(input.Password))

	if !dbError(err) {
		gentoken, err := jwtoken.GenerateToken(input.UserName)
		if err != nil {
			return nil, err
		}

		return &model.Response{Token: gentoken, Error: "Okay"}, nil
	}

	return &model.Response{Token: "", Error: "Wrong username or password"}, nil

}

func (r *queryResolver) Response(ctx context.Context) (*model.Response, error) {
	res := &model.Response{Token: "", Error: "nothing here"}
	return res, nil
}

func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
	user := auth.ForContext(ctx)

	if user == "" {
		return nil, fmt.Errorf("access denied")
	}

	rows, err := DB.Query("select * from client")

	if dbError(err) {
		return nil, err
	}

	clients := []*model.Client{}

	defer rows.Close()
	for rows.Next() {
		client := &model.Client{}
		err := rows.Scan(&client.ClientID, &client.UserName, &client.Password,
			&client.FirstName, &client.LastName, &client.PhoneNumber,
			&client.Gender)

		if dbError(err) {
			return nil, err
		}

		clients = append(clients, client)
	}
	return clients, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
