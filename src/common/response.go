package common

type DataFormat struct {
	Status  uint8       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(status uint8, message string, data interface{}) DataFormat {
	return DataFormat{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ResponseSuccess(data interface{}) DataFormat {
	return Response(1, "success", data)
}

func ResponseFail(data interface{}) DataFormat {
	return Response(0, "fail", data)
}
