package helpers

type Response struct{
	Status	Status	`json:"status"`
	Data	any		`json:"data"`
}

type Status struct{
	Message	string	`json:"message"`
	Code	int		`json:"code"`
}

func APIResponse[D any](message string, code int, data D) Response{
	status := Status{
		Message: message,
		Code: code,
	}

	jsonResponse := Response{
		Status : status,
		Data : data,
	}

	return jsonResponse
}

type Response2 struct {
	Status Status2 `json:"status"`
	Meta   Meta2   `json:"meta"`
	Data   any     `json:"data"`
}

type Meta2 struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Status2 struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func APIResponse2[D any](message string, code int, page int, limit int, total int, data D) Response2 {
	status := Status2{
		Message: message,
		Code:    code,
	}

	meta := Meta2{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	jsonResponse := Response2{
		Status: status,
		Meta:   meta,
		Data:   data,
	}

	return jsonResponse
}