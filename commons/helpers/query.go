package helpers

// Query is the result of a request
type Query struct{
	Fields 		string 	`json:"fields"`
	Limit 		int 	`json:"limit"`
	After 		Range 	`json:"after"`
	Before 		Range 	`json:"before"`
	OrderField 	string 	`json:"orderField"`
	Order 		string 	`json:"Order"`
}

// Range is the time limit data
type Range struct{
	Field 	string `json:"field"`
	Date 	string `json:"date"`
}