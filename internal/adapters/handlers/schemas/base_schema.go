package schemas

type (
	BaseResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	BaseDataResponse struct {
		BaseResponse
		Data interface{} `json:"data,omitempty"`
	}
)

func NewBaseResponse(success bool, message string) BaseResponse {
	return BaseResponse{
		Success: success,
		Message: message,
	}
}

func NewDataResponse(success bool, message string, data interface{}) BaseDataResponse {
	return BaseDataResponse{
		BaseResponse: NewBaseResponse(success, message),
		Data:         data,
	}
}
