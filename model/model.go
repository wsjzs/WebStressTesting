package model

type Target struct {
	Url       string
	Interval  int
	Goroutine int
}

//receive from html
type Mama struct {
	Purpose   string `json:"purpose"`
	Url       string `json:"url"`
	Interval  int    `json:"interval"`
	Goroutine int    `json:"goroutine"`
}

//send to html
type Papa struct {
	State       int    `json:"state"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func BuildPapa(state int,description,data string)Papa{
	result:=Papa{
		State:state,
		Description:description,
		Data:data,
	}
	return result
}