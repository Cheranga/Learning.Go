package customer

type Customer struct {
	Data    CustomerData
	Support SupportData
}

type CustomerData struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type SupportData struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type GetCustomerByIdResponse struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Avatar    string
	Url       string
	Text      string
}

type GetCustomerByIdRequest struct {
	CustomerId int
}

type GetCustomersByPageIdRequest struct {
	PageId int
}

type GetCustomersResponse struct {
	Total int            `json:"total"`
	Data  []CustomerData `json:"data"`
}
