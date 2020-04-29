package client


type Client struct {
	ClientID    string  
	UserName    string  
	Password    string  
	FirstName   string  
	LastName    string 
	PhoneNumber string  
	Gender      *string 
}

func (cli *Client) getClientUsername() {
	
}