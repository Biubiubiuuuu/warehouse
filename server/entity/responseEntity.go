package entity

type ResponseData struct {
	Status  bool                   `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

// response struct
// {"status":true,"data":{},"message":""}
// {"status":true,"data":nil,"message":""}
// {"status":false,"data":nil,"message":""}
func ResponseJson(status bool, data map[string]interface{}, message string) (responseData ResponseData) {
	responseData.Status = status
	responseData.Data = data
	responseData.Message = message
	return responseData
}
