package models

type Field struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
}

type ContentType struct {
	ContentName string  `json:"content_name"`
	ContentDesc string  `json:"content_desc"`
	Fields      []Field `json:"fields"`
}
