package util

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v any) {
	b, _ := json.MarshalIndent(v, "", " ")
	fmt.Println(string(b))
}
