package pkg

import "github.com/gin-gonic/gin"

type BodJson struct {
	Code    int `json:"code"`
	Data    any `json:"data,omitempty"`
	Message any `json:"pesan,omitempty"`
}

func Responses(code int, data *BodJson) *BodJson {
	res := &BodJson{}

	res.Code = code

	if data.Data != nil {
		res.Data = data.Data
	}
	if data.Message != nil {
		res.Message = data.Message
	}
	return res
}

func (r *BodJson) Send(c *gin.Context) {
	c.JSON(r.Code, r)
	c.Abort()
}
