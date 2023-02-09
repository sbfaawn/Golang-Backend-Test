package response

type Response struct {
	Message string      `json:"message" xml:"message" form:"message" query:"message"`
	Data    interface{} `json:"data" xml:"data" form:"data" query:"data"`
}
