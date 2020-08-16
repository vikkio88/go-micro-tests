package models

import "github.com/gookit/validate"

// Page model to handle page objects
type Page struct {
	Title string `json:"title" validate:"required|minLen:7"`
}

// IsValid validating Page creation request
func (p *Page) IsValid() (bool, *validate.Validation) {
	v := validate.Struct(p)

	return v.Validate(), v
}
