package utils

import (
	"encoding/json"
	"fmt"
)

func DebugLog(u interface{}) {
	empJSON, _ := json.MarshalIndent(u, "", "  ")
	fmt.Printf("%s \n", empJSON)
}
