package graphqltest

import (
	"context"

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
	print("here\n")
	statement, err := db.Prepare("insert into client (fullname, gender, phonenumber, username, hashedpassword) values($1, $2, $3, $4, $5)")
	CheckError(err)
	print("here\n")
	//TODO: Check inputs for uniqueness and appropriate characters.
	// pw := string(input.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),
		bcrypt.DefaultCost)
	CheckError(err)

	hashedInputpw := string(hash)
	_, err = statement.Exec(input.FullName,
		input.Gender, input.PhoneNumber, input.UserName, hashedInputpw)
	CheckError(err)

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
