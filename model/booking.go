package model

type Booking struct {
	OrderID  string `json:"orderid,omitempty"`
	SID string `json:"sid,omitempty"`
	Name  string `json:"name"`
	RD1  string `json:"rd1"`
	CreateDate  string `json:"createDate"`
	PostDate  string `json:"postDate"`
	Phone  string `json:"phone"`
	Phone2  string `json:"phone2"`
	Email  string `json:"email"`
	RD3  string `json:"rd3"`
	RD4  string `json:"rd4"`
	Address  string `json:"address"`
	Note  string `json:"note"`
}