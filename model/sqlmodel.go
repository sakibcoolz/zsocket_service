package model

type Request struct {
	SQL  string `json:"sql"`
	Type string `json:"type"`
}

type Response struct {
	Data string `json:"data"`
	Msg  string `json:"msg"`
}
