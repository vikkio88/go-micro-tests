package models

import (
	"fmt"

	"github.com/gookit/validate"
)

// Page model to handle page objects
type Page struct {
	ID    string `json:"id" validate:"optional"`
	Title string `json:"title" validate:"required|minLen:7"`
	Lists []list `json:"lists"`
}

type list struct {
	ID    string `json:"id" validate:"optional"`
	Title string `json:"title" validate:"required|minLen:2"`
}

// IsValid validating Page creation request
func GetAllPages() []Page {
	pages := make([]Page, 3)

	for i := 0; i < 10; i++ {
		page := Page{ID: fmt.Sprintf("someid%d", i), Title: fmt.Sprintf("sometitle%d", i)}
		lists := make([]list, 5)
		for j := 0; j < 5; j++ {
			lists[j] = list{ID: fmt.Sprintf("%d-listid-%d", i, j), Title: fmt.Sprintf("listtitle-%d", j)}
		}
		page.Lists = lists
		pages = append(pages, page)
	}

	return pages
}
func (p *Page) IsValid() (bool, []validate.Errors) {
	vPage := validate.Struct(p)
	isValid := vPage.Validate()
	errors := make([]validate.Errors, 0)
	if !isValid {
		errors = append(errors, vPage.Errors)
	}

	for _, listItem := range p.Lists {
		validator := validate.Struct(list(listItem))
		isCurrentValid := validator.Validate()
		isValid = isValid && isCurrentValid
		if !isCurrentValid {
			errors = append(errors, validator.Errors)
		}
	}

	return isValid, errors
}
