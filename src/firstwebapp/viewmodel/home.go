package viewmodel

type Home struct {
	Title string
}

func NewHome() Home {
	result := Home{
		Title: "Home Page",
	}
	return result
}
