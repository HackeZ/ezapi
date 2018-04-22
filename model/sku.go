package model

// SKU 最小库存管理单元
type SKU struct {
	ID           string  `json:"sku_id"`
	Name         string  `json:"name"`
	ProductID    int64   `json:"product_id"`
	SellerSKUID  string  `json:"seller_sku_id"`
	RetailPrice  float64 `json:"retail_price"`
	SellingPrice float64 `json:"selling_price"`
	Stock        int32   `json:"stock"`
	Weight       float64 `json:"weight"`
	IsOnSale     bool    `json:"is_onsale"`
	Length       int32   `json:"length"`
	Width        int32   `json:"width"`
	Height       int32   `json:"height"`
}

// SKUEdition SKU编辑属性
type SKUEdition struct {
	Props        []SKUProperty       `json:"props"`
	CustomProps  []SKUCustomProperty `json:"customs_props"`
	SellerSKUID  string              `json:"seller_sku_id"`
	RetailPrice  float64             `json:"retail_price"`
	SellingPrice float64             `json:"selling_price"`
	Stock        int32               `json:"stock"`
	Weight       float64             `json:"weight"`
	IsOnSale     bool                `json:"is_onsale"`
	Length       int32               `json:"length"`
	Width        int32               `json:"width"`
	Height       int32               `json:"height"`
}

// SKUProperty SKU属性
type SKUProperty struct {
	PID int32 `json:"pid"`
	VID int32 `json:"vid"`
}

// SKUCustomProperty SKU自定义属性
type SKUCustomProperty struct {
	PID   int32  `json:"pid"`
	Value string `json:"value"`
}

// SKUPropertyImage SKU属性图片
type SKUPropertyImage struct {
	PID   int32  `json:"pid"`
	VID   int32  `json:"vid"`
	Image string `json:"image"`
}

// SKUStockOut 订单SKU缺货信息
type SKUStockOut struct {
	SKUID    int64  `json:"sku_id"`
	Quantity int32  `json:"quantity"`
	Reason   string `json:"reason"`
}
