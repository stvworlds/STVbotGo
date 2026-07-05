package ticketsData

import "strings"

func RenderTemplate(tmpl string, values map[string]string) string {
	result := tmpl
	for key, value := range values {
		result = strings.ReplaceAll(result, "{"+key+"}", value)
	}
	return result
}
