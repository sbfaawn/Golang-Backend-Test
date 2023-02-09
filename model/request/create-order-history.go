package request

type OrderHistory struct {
	UserId      int    `json:"userId" xml:"userId" form:"userId" query:"userId"`
	OrderItemId int    `json:"orderItemId" xml:"orderItemId" form:"orderItemId" query:"orderItemId"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
}
