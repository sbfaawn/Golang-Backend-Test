package request

type User struct {
	Id         *int   `json:"id,omitempty" xml:"id" form:"id" query:"id"`
	FullName   string `json:"fullName" xml:"fullName" form:"fullName" query:"fullName"`
	FirstOrder string `json:"firstOrder" xml:"firstOrder" form:"firstOrder" query:"firstOrder"`
}
