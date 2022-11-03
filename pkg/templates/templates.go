package templates

import (
	"html/template"
)

type Model struct {
	Home *template.Template
	Details *template.Template
	FormAdd *template.Template
	FormRename *template.Template
	FormCheckIn *template.Template
	FormDeactivate *template.Template
	FormRemove *template.Template
}