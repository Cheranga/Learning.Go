package customer

type Customer struct {
	Data    CustomerData
	Support SupportData
}

type CustomerData struct {
	Id        int    `json: "id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json: "avatar"`
}

type SupportData struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}
