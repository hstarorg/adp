package god

type Request struct {
	body        interface{}
	url         string
	headers     map[string]string
	method      string
	path        string
	query       map[string]interface{}
	querystring string
	ip          string
}

func (r *Request) GetHeader(header string) string {
	return r.headers[header]
}
