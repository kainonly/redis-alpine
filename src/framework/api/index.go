package api

type IndexController struct{}

func (c *IndexController) Get() interface{} {
	return map[string]string{"message": "this api"}
}
