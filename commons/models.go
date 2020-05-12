package commons

//NewResponse is what is sended to the requester through API
type NewResponse struct{
	Count int
	Data interface{}
	Message APIResponse
}
