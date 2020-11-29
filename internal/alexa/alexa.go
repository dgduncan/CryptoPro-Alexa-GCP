package alexa

// Response test
type Response struct {
	UUID       string `json:"uid"`
	UpdateDate string `json:"updateDate"`
	TitleText  string `json:"titleText"`
	MainText   string `json:"mainText"`
}
