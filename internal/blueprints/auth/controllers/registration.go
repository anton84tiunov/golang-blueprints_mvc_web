package auth

import (
	"encoding/json"
	"fmt"

	models "../../../models"
)

func RegDataParse(data map[string]interface{}) {
	// convert map to json
	jsonString, _ := json.Marshal(data)
	fmt.Println("string(jsonString)", string(jsonString))

	// convert json to struct
	s := models.User_json{}
	err_json := json.Unmarshal(jsonString, &s)
	if err_json != nil {
		fmt.Println(err_json)
	}
	fmt.Println(s)
	fmt.Println(s.Birthday)
}
