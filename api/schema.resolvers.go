package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/generated"
	"api/jwtoken"
	"api/model"
	"api/pdf"
	"context"
	"fmt"
	"strings"

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

	token, err := jwtoken.GenerateToken(input.UserName)
	if err != nil {
		return nil, err
	}
	res := &model.Response{Token: token}

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

		return &model.Response{Token: gentoken}, nil
	}

	return &model.Response{Token: ""}, fmt.Errorf("Wrong username or password!")
}

func (r *queryResolver) RefreshToken(ctx context.Context, input model.Oldtoken) (*model.Response, error) {
	username, err := jwtoken.VerifyToken(input.Token)
	if err != nil {
		return nil, err
	}
	token, err := jwtoken.GenerateToken(username)
	if err != nil {
		return nil, err
	}
	return &model.Response{Token: token}, nil
}

func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
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

func (r *queryResolver) Allshops(ctx context.Context) ([]*model.Shop, error) {
	rows, err := DB.Query("select * from shop_ratings")

	if dbError(err) {
		return nil, err
	}

	shops := []*model.Shop{}

	defer rows.Close()
	for rows.Next() {
		shop := &model.Shop{}
		err := rows.Scan(&shop.ShopID, &shop.StreetAddr, &shop.State,
			&shop.AreaCode, &shop.City, &shop.Country, &shop.ShopName,
			&shop.Latitude, &shop.Longitude, &shop.Rating)

		if dbError(err) {
			return nil, err
		}

		shops = append(shops, shop)
	}

	return shops, nil
}

func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	rows, err := DB.Query("select * from service")

	if dbError(err) {
		return nil, err
	}

	services := []*model.Service{}

	defer rows.Close()

	for rows.Next() {
		service := &model.Service{}
		err := rows.Scan(&service.ServiceID, &service.ServiceName,
			&service.ServiceDescription, &service.Price,
			&service.CustomDuration)

		if dbError(err) {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (r *queryResolver) WeeklyAppt(ctx context.Context, input model.Shopidentifier) ([]*model.AppointmentWeek, error) {
	querystring := `select a.apptid, a.barberid, a.apptdate, a.starttime, a.endtime 
						from appointment a 
						join 
						barber b 
						on a.barberid=b.barberid 
						where shopid = $1 
						and apptdate >= NOW()
						and apptdate <= CURRENT_DATE + integer '7'
						and a.clientcancelled = false`
	stmt, err := DB.Prepare(querystring)
	if dbError(err) {
		return nil, err
	}
	rows, err := stmt.Query(input.ShopID)

	if dbError(err) {
		return nil, err
	}

	apptweeks := []*model.AppointmentWeek{}
	defer rows.Close()
	for rows.Next() {
		apptweek := &model.AppointmentWeek{}
		rows.Scan(&apptweek.ApptID, &apptweek.BarberID, &apptweek.ApptDate,
			&apptweek.StartTime, &apptweek.EndTime)

		apptweek.ApptDate = strings.Split(apptweek.ApptDate, "T")[0]
		apptweek.StartTime = strings.Split(apptweek.StartTime, "T")[1]
		apptweek.StartTime = apptweek.StartTime[:len(apptweek.StartTime)-1]
		apptweek.EndTime = strings.Split(apptweek.EndTime, "T")[1]
		apptweek.EndTime = apptweek.EndTime[:len(apptweek.EndTime)-1]
		apptweeks = append(apptweeks, apptweek)
	}
	return apptweeks, nil
}

func (r *queryResolver) BarbersAtShop(ctx context.Context, input model.Shopidentifier) ([]*model.AllBarbersAtShop, error) {
	querystring := `select barberid, firstname, lastname 
					from barber 
					where shopid=$1`
	stmt, err := DB.Prepare(querystring)
	if dbError(err) {
		return nil, err
	}
	rows, err := stmt.Query(input.ShopID)

	if dbError(err) {
		return nil, err
	}

	bsps := []*model.AllBarbersAtShop{}
	defer rows.Close()
	for rows.Next() {
		bsp := &model.AllBarbersAtShop{}
		rows.Scan(&bsp.BarberID, &bsp.FirstName, &bsp.LastName)
		bsps = append(bsps, bsp)
	}
	return bsps, nil
}

func (r *queryResolver) Receipt(ctx context.Context, input model.Receiptinput) ([]*model.ReceiptData, error) {
	querystring := ` with appt_dets as (
		select a.apptid, a.clientid, a.barberid, a.paymenttype,
		 a.apptdate, a.starttime, a.endtime, s.servicename, 
		 inc.price from appointment a natural join includes 
		 inc join service s on inc.serviceid=s.serviceid where a.apptid = $1 and a.clientid = $2
	 )
	 select ad.*, s.shopname, s.streetaddr, s.city, s.state, 
	 b.firstname as barberfirst, b.lastname as barberlast,
	  c.firstname, c.lastname from 
	  shop s join barber b on s.shopid=b.shopid 
	  join appt_dets ad on b.barberid=ad.barberid 
	  join client c on c.clientid=ad.clientid;`
	stmt, err := DB.Prepare(querystring)
	if dbError(err) {
		return nil, err
	}
	rows, err := stmt.Query(input.ApptID, input.ClientID)

	if dbError(err) {
		return nil, err
	}

	bsps := []*model.ReceiptData{}
	defer rows.Close()
	for rows.Next() {
		bsp := &model.ReceiptData{}
		rows.Scan(&bsp.ApptID, &bsp.ClientID, &bsp.BarberID, &bsp.Paymenttype,
			&bsp.ApptDate, &bsp.StartTime, &bsp.EndTime, &bsp.ServiceName,
			&bsp.Price, &bsp.ShopName, &bsp.Shopstreetaddr, &bsp.ShopCity,
			&bsp.ShopState, &bsp.Barberfirstname, &bsp.Barberlastname,
			&bsp.Clientfirstname, &bsp.Clientlastname)

		bsp.ApptDate = strings.Split(bsp.ApptDate, "T")[0]
		bsp.StartTime = strings.Split(bsp.StartTime, "T")[1]
		bsp.StartTime = bsp.StartTime[:len(bsp.StartTime)-1]
		bsp.EndTime = strings.Split(bsp.EndTime, "T")[1]
		bsp.EndTime = bsp.EndTime[:len(bsp.EndTime)-1]
		bsps = append(bsps, bsp)
	}
	pdf.Receipt(bsps)

	return bsps, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
