package main

import (
	"encoding/json"
	"fmt"
)

// We can use a map if we are not sure what we expect

// if we use a map[string]interface{} as receiver for json.Unmarshal.
/*
1. Whatever name of json will be interpreted as string for map key
2. If the value is another json, it will be nested map, otherwise will try to resolve
3. However, looking above it is very hard to actually get values. So we define struct instead.
4. We can use a struct to receive, as long as data type for value & json tag for key can match
5. We don't have to match all of them, partial match is also fine
*/

const randomJsonString = `{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    },
	"count": 1
}`

func useMapToAcceptJson() {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(randomJsonString), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
	fmt.Printf("%s\n", m["glossary"].(map[string]interface{})["GlossDiv"].(map[string]interface{})["title"])
}

const otherJSONString = `{
	"data": [
		{ "name":"n1", "id": "1"},
		{ "name":"n2", "id": "2"},
		{ "name":"n3", "id": "3"}
	]
}`

func resolveByStruct() {
	m := struct {
		Data []struct {
			Name string `json:"name"`
		} `json:"data"`
	}{}

	json.Unmarshal([]byte(otherJSONString), &m)
	fmt.Println(m.Data[0].Name, m.Data[1].Name)
}

/*
func main() {
	// useMapToAcceptJson()
	resolveByStruct()
}
*/
