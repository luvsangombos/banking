package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Luka", "Ulaanbaatar", "11560", "1994-06-18", "active"},
		{"2", "Doncic", "Dallas", "52557", "2000-02-01", "active"},
	}
	return CustomerRepositoryStub{customers: customers}

}
