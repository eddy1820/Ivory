package domain

type Response struct {
	Method string `form:"method" json:"method"`
	Path   string `form:"path" json:"path"`
	Data   string `form:"data" json:"data"`
}
