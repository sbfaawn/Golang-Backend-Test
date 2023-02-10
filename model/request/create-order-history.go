package request

type OrderHistory struct {
	Id          *int   `json:"id;omitempty" xml:"id" form:"id" query:"id"`
	UserId      int    `json:"userId" xml:"userId" form:"userId" query:"userId"`
	OrderItemId int    `json:"orderItemId" xml:"orderItemId" form:"orderItemId" query:"orderItemId"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
}
