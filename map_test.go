package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"p0001", "name":"Apple Mac Book Pro", "price": 20000000}`

	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

//Untuk mengatasi Struct dalam JSON yang data-nya dynamic, digunakan map[string][interface{}]
//Atribut akan menjadi key di map, dan value menjadi value di map
//Karena value-nya interface{}, maka harus ada konversi manual
//Tipe data Map tidak mendukung JSON Tag