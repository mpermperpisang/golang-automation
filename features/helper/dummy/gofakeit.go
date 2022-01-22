package dummy

import "github.com/brianvoe/gofakeit"

func GoFakeItAddress() *gofakeit.AddressInfo {
	return gofakeit.Address()
}

func GoFakeItCity() string {
	return gofakeit.City()
}

func GoFakeItColor() string {
	return gofakeit.Color()
}

func GoFakeItCompany() string {
	return gofakeit.Company()
}

func GoFakeItCountry() string {
	return gofakeit.Country()
}

func GoFakeItContact() *gofakeit.ContactInfo {
	return gofakeit.Contact()
}

func GoFakeItEmail() string {
	return gofakeit.Email()
}

func GoFakeItFirstName() string {
	return gofakeit.FirstName()
}

func GoFakeItGender() string {
	return gofakeit.Gender()
}

func GoFakeItLastName() string {
	return gofakeit.LastName()
}

func GoFakeItPhone() string {
	return gofakeit.Phone()
}

func GoFakeItStreet() string {
	return gofakeit.Street()
}
