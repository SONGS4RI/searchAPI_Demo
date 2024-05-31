package response

type Response struct {
	Desc   string `json:"desc"  example:"desc"`
	Status int    `json:"status" example:"200"`
	Result []interface{} `json:"result" "example:[]"`
}