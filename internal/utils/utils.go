package utils

import "encoding/json"

type responceOk struct {
	Status    string
	ReturnMsg string
}

type responceErr struct {
	Status    string
	ReturnMsg string
	ErrMesage string
}

func SendOkResponce(returnMsg string) ([]byte, error) {
	responce := responceOk{
		Status:    "OK",
		ReturnMsg: returnMsg,
	}

	return json.Marshal(responce)
}

func SenErrorMessage(returnMsg string, errReason string) ([]byte, error) {
	responce := responceErr{
		Status:    "ERROR",
		ReturnMsg: returnMsg,
		ErrMesage: errReason,
	}

	return json.Marshal(responce)
}
