package spree

type Address struct {
	Id               int
	FirstName        string
	LastName         string
	Address1         string
	Address2         string
	City             string
	ZipCode          string
	Phone            string
	Company          string
	AlternativePhone string `json:"alternative_phone"`
	Country          country
	State            state
}

type state struct {
	Id        int
	Name      string
	Abbr      string
	CountryId int `json:"country_id"`
}

type country struct {
	Id      int
	IsoName string `json:"iso_name"`
	Iso     string
	Iso3    string
	Name    string
	NumCode int
}
