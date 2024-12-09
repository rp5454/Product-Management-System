package models

type Product struct {
	ID                 int      `json:"id,omitempty"`
	ProductName        string   `json:"name"`
	ProductDescription string   `json:"description"`
	ProductPrice       float64  `json:"price"`
	ProductImages      []string `json:"image_url"`
}
