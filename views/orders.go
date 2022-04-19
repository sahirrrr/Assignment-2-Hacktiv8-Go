package views

type OrdersCreateView struct {
	OrderedAt    string             `json:"ordered_at"`
	CustomerName string             `json:"customer_name"`
	Items        []ItemsCreateViews `json:"items"`
}

type ItemsCreateViews struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
