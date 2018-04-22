package model

// Product 商品信息
type Product struct {
	ProductID         int64              `json:"product_id"`
	NameCN            string             `json:"name"`
	NameEN            string             `json:"name_en"`
	IsSensitive       bool               `json:"is_sensitive"`
	SKUIDs            []int64            `json:"sku_ids"`
	IsOnSale          bool               `json:"is_onsale"`
	CreateDate        string             `json:"create_date"`
	Images            []string           `json:"images"`
	Description       string             `json:"description"`
	CID               string             `json:"cid"`
	Spec              map[string]string  `json:"spec"`
	SKUPropertyImages []SKUPropertyImage `json:"sku_prop_image"`
}

// ProductDetail 商品详情
type ProductDetail struct {
	Product
	SKUs []SKU `json:"skus"`
}
