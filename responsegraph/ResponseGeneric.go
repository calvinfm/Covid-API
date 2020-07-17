package responsegraph

type ResponseGenericGet struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseGeneric struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
