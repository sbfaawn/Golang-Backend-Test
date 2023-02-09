package request

type OrderItem struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Price     uint   `json:"price" xml:"price" form:"price" query:"price"`
	ExpiredAt string `json:"expiredAt" xml:"expiredAt" form:"expiredAt" query:"expiredAt"`
}
