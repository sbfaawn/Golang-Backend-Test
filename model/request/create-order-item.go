package request

type OrderItem struct {
	Id        *int   `json:"id,omitempty" xml:"id" form:"id" query:"id"`
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Price     uint   `json:"price" xml:"price" form:"price" query:"price"`
	ExpiredAt string `json:"expiredAt" xml:"expiredAt" form:"expiredAt" query:"expiredAt"`
}
