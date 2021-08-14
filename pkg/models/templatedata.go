package models

//hold data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //dont know return type
	CSRFToken string
	Flash     string //flash message to user
	Warning   string
	Error     string
}
