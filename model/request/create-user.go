package request

type User struct {
	FullName   string `json:"fullName" xml:"fullName" form:"fullName" query:"fullName"`
	FirstOrder string `json:"firstOrder" xml:"firstOrder" form:"firstOrder" query:"firstOrder"`
}
