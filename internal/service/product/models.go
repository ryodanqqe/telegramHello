package product

type Product struct {
	Title string `json:"title"`
}

var Products = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}
