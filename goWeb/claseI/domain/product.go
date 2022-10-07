package domain

type Product struct{
	Id int `json:"Id"`
    Name string `json:"Name"`
	Colour string `json:"Colour"`
	Price float32 `json:"Price"`
    Stock int `json:"Stock"`
	Code string `json:"Code"`
	Published bool `json:"Published"`
	CreationDate string `json:"creationDate"`
}