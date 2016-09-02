package controllers

type Product struct {
	Name string `json:"full_name", db:"product_name"`
	ID   string `json:"-"`
}

func GetProducts() ([]Product, error) {
	p1 := Product{
		Name: "polo",
		ID:   "23354",
	}
	p2 := Product{
		Name: "jeans",
		ID:   "2335dsfsdf4",
	}

	return []Product{p1, p2}, nil
}
