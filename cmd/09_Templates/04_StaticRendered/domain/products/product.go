package products

type Product struct {
	SKU          string `json:"sku"`
	CODE         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	HTMLTemplate string `json:"html"`

	OnSaleUntil uint64 `json:"onsaleunixtime"`

	Quantity       uint8 `json:"quantity"`
	PriceCents     uint8 `json:"pricecents"`
	SalePriceCents uint8 `json:"salepricecents"`
	SaleDiscount   uint8 `json:"salediscount"`
}
