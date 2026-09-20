package database

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "",
		Description: "",
		Price:       0,
		ImgUrl:      "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2018/08/bananas-1354785_1920.jpg",
	}
	Productlist = append(Productlist, prd1)
}
