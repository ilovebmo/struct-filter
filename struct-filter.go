// structfilter provides StructFilter function for filtering structs based on a map
package structfilter

import (
	"reflect"
)

// StructFilter returns a slice of any struct s after being filtered with the filters map
func StructFilter(filters map[string]any, s []any) []any {
	filtered := []any{}

	for _, element := range s {
		passed := false
		iter := reflect.ValueOf(element)
		typeElement := iter.Type()
		for i := 0; i < iter.NumField(); i++ {
			if iter.Field(i).String() != filters[typeElement.Field(i).Name] && filters[typeElement.Field(i).Name] != nil {
				if passed {
					filtered = filtered[0 : len(filtered)-1]
				}
				break
			}
			if !passed {
				filtered = append(filtered, element)
				passed = true
			}
		}
	}
	return filtered
}
