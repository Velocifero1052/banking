package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Rakhmon", "Tashkent", "1000008", "01.05.1993", "not married"},
		{"2", "Murad", "Tashkent", "1000008", "26.02.1996", "married"},
		{"3", "Dono", "Tashkent", "1000008", "14.02.1969", "not married"},
	}
	return CustomerRepositoryStub{customers: customers}
}
