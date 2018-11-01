package controller

type ExampleController struct{}

func (c *ExampleController) PostPing() interface{} {
	return map[string]map[string]string{"message": {
		"name": "kain",
	}}
}
