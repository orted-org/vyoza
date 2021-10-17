package db

import (
	"errors"
	"fmt"
	"strings"
)

func CreateDynamicUpdateQuery(incomingMap map[string]interface{}, allowedFields map[string]string, tableName, closing string) (string, error) {
	if len(incomingMap) == 0 {
		return "", errors.New("no field found for update")
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("UPDATE %s ", tableName))
	i := 0
	for k, v := range incomingMap {
		dataType, err := isAllowedType(k, v, allowedFields)
		if err != nil {
			return "", err
		}
		if dataType != "" {
			if i != 0 {
				sb.WriteString(", ")
			} else {
				sb.WriteString("SET ")
			}
			if dataType == "string" {
				sb.WriteString(fmt.Sprintf("%s = '%v'", k, v))
			} else {
				sb.WriteString(fmt.Sprintf("%s = %v", k, v))
			}
		}
		i++
	}

	sb.WriteString(" " + closing)
	return sb.String(), nil
}
func isAllowedType(field string, data interface{}, allowedFields map[string]string) (string, error) {
	for k, v := range allowedFields {
		if k == field {
			dataType := fmt.Sprintf("%T", data)
			if dataType == v {
				return dataType, nil
			} else {
				if v == "custom" {
					return "custom", nil
				}
				return "", fmt.Errorf("datatype mismatch for %v, expected %v, got %v", k, v, dataType)
			}
		}
	}
	return "", nil
}
