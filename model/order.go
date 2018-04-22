package model

import (
	"github.com/hackez/ezapi/model/orderstatus"
	"github.com/hackez/ezapi/model/warehouse"
)

// Order 订单信息
type Order struct {
	Number      string                  `json:"order_number"`
	ProductID   int64                   `json:"product_id"`
	ProductName string                  `json:"product_name"`
	SellerSKUID string                  `json:"seller_sku_id"`
	SKUName     string                  `json:"sku_name"`
	UnitPrice   float64                 `json:"unit_price"`
	Quantity    int64                   `json:"quantity"`
	Status      orderstatus.OrderStatus `json:"status"`
	DispathDate string                  `json:"dispath_date"`
	PaidDate    string                  `json:"paid_date"`
	Warehouse   warehouse.Warehouse     `json:"warehouse"`
	Remark      string                  `json:"remark"`
	CancelDate  string                  `json:"cancel_date"`
}
