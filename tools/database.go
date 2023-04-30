package tools

import "fmt"

func BuildUpdateQueryArgs(FieldMap map[string]interface{}, username string) ([]string, []interface{}, int) {
	updateArr := []string{}
    args := []interface{}{}
    argsCount := 1

	for field, value := range FieldMap {
		strValue := value.(string)
		if strValue != "" {
            updateArr = append(updateArr, fmt.Sprintf("%s = $%d", field, argsCount))
            args = append(args, strValue)
            argsCount++
        }
	}

	return updateArr, append(args, username), argsCount
}