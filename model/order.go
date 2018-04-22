package model

//go:generate stringer -type OrderStatus
//go:generate stringer -type Warehouse

// OrderStatus new order status
type OrderStatus int32

// Warehouse determines where you need to send the order
type Warehouse int32

const ()

// Order 订单信息
type Order struct {
	Number      string      `json:"order_number"`
	ProductID   int64       `json:"product_id"`
	ProductName string      `json:"product_name"`
	SellerSKUID string      `json:"seller_sku_id"`
	SKUName     string      `json:"sku_name"`
	UnitPrice   float64     `json:"unit_price"`
	Quantity    int64       `json:"quantity"`
	Status      OrderStatus `json:"status"`
	DispathDate string      `json:"dispath_date"`
	PaidDate    string      `json:"paid_date"`
	Warehouse   Warehouse   `json:"warehouse"`
	Remark      string      `json:"remark"`
	CancelDate  string      `json:"cancel_date"`
}
