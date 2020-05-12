package commons

//NewResponse is what is sended to the requester through API
type NewResponse struct{
	Count int `json:"count"`
	Data interface{} `json:"data"`
	Message APIResponse `json:"response"`
}
