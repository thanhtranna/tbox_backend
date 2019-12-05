package common

type DataFormat struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(message string, data interface{}) DataFormat {
	return DataFormat{
		Message: message,
		Data:    data,
	}
}

func ResponseSuccess(data interface{}) DataFormat {
	return Response("success", data)
}

func ResponseFail(data interface{}) DataFormat {
	return Response("fail", data)
}
