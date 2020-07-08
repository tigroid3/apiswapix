package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func CheckStruct(validate *validator.Validate, structure interface{}) (map[string]string, error) {
	var err error
	formattedErrors := map[string]string{}
	err = validate.Struct(structure)
	if err != nil {
		return prepareErrors(err.(validator.ValidationErrors), structure), err
	}

	return formattedErrors, nil
}

func prepareErrors(fErrs []validator.FieldError, structure interface{}) map[string]string {
	mapMessages := getMessages()
	returnedErrs := make(map[string]string)
	mapJsonFields := getMapJsonFieldsByStruct(structure)

	for _, f := range fErrs {
		tag := f.Tag()
		jsonField := mapJsonFields[f.Field()]
		if val, ok := mapMessages[tag]; ok {
			returnedErrs[jsonField] = val
		} else {
			returnedErrs[jsonField] = fmt.Errorf("%s", f).Error()
		}
	}

	return returnedErrs
}

//messages format tagValidation => message
func getMessages() map[string]string {
	messages := map[string]string{
		"exists-category-id": "Category not exists",
		"exists-location-id": "Location not exists",
		"lte":                "Maximum character size exceeded",
	}

	return messages
}

func getMapJsonFieldsByStruct(structure interface{}) map[string]string {
	mapFields := map[string]string{}
	reflectType := reflect.ValueOf(structure)

	for i := 0; i < reflectType.Type().NumField(); i++ {
		typeField := reflectType.Type().Field(i)
		mapFields[typeField.Name] = typeField.Tag.Get("json")
	}

	return mapFields
}
