package response

type Response struct {
	Desc   string `json:"desc"  example:"null"`
	Status int    `json:"status" example:"200"`
	// Result `json:"result,omitempty" example:"null"`
	Result []interface{} `json:"result,omitempty" example:"null"`
}