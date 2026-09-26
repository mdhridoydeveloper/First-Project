package database

var productlist []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productlist)
	productlist = append(productlist, p)
	return p
}

func List() []Product {
	return productlist
}

func Get(productID int) *Product {
	for _, product := range productlist {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}
func Update(product Product) {
	for idx, p := range productlist {
		if p.ID == product.ID {
			productlist[idx] = product
		}
	}

}
func Delete(productID int) {
	var temList []Product

	for _, p := range productlist {
		if p.ID != productID {
			temList = append(temList, p)
		}
	}

	productlist = temList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "",
		Description: "",
		Price:       0,
		ImageUrl:    "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2018/08/bananas-1354785_1920.jpg",
	}
	productlist = append(productlist, prd1)
}
