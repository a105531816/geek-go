package apis

type Person struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Age    int     `json:"age"`
	Weight int     `json:"weight"`
	Height float64 `json:"height"`
}