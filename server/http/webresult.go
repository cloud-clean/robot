package http

type WebResult struct {
	Code int						`json:"code"`
	Result map[string]interface{}	`json:"result"`
}

func NewResult(code int,res interface{}) *WebResult{
	data := map[string]interface{}{
		"data":res,
	}
	return &WebResult{Code:code,Result:data,}
}

func (res *WebResult)SetData(key string,data interface{}){
	res.Result[key]=data
}
