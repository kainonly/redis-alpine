package controller

type IndexController struct{}

func (c *IndexController) PostPing() interface{} {
	return map[string]map[string]string{"message": {
		"name": "kain",
	}}
}
