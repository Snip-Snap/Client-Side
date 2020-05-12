// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Client struct {
	ClientID    string  `json:"clientID"`
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"phoneNumber"`
	Gender      *string `json:"gender"`
}

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type NewClient struct {
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Gender      *string `json:"gender"`
	PhoneNumber string  `json:"phoneNumber"`
}

type Oldtoken struct {
	Token string `json:"token"`
}

type Response struct {
	Token string `json:"token"`
}

type Shop struct {
	ShopID     string  `json:"shopID"`
	ShopName   string  `json:"ShopName"`
	StreetAddr string  `json:"StreetAddr"`
	City       string  `json:"City"`
	State      string  `json:"State"`
	AreaCode   string  `json:"AreaCode"`
	Country    string  `json:"Country"`
	Latitude   string  `json:"Latitude"`
	Longitude  string  `json:"Longitude"`
	Rating     float64 `json:"Rating"`
}
