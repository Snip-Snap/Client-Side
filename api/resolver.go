package graphqltest

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
	input NewClient) (*Response, error) {
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
	client := &Client{string(0), input.UserName,
		hashinputpass, input.FullName, input.Gender, input.PhoneNumber}
	fmt.Printf("%v+", sum)
	fmt.Printf("%v", client)
	res := &Response{Error: "Okay"}

	return res, nil

}
func (r *mutationResolver) SignUpBarber(ctx context.Context, input NewBarber) (*Response, error) {

	stmt, err := db.Prepare("insert into barber (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")

	if err != nil {
		return nil, err
	}

	hashpw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(input.FullName, input.Gender, input.PhoneNumber, input.UserName, string(hashpw))
	if err != nil {
		return nil, err
	}
	res := &Response{Error: "Okay"}

	return res, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Response(ctx context.Context) (*Response, error) {
	res := &Response{Error: "nothing here"}
	return res, nil

}
