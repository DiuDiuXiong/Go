package main

/**
1. map[<key_type>]<value_type> : 1D map
2. map[<key_type1>]map[<key_type2>]<value_type>: 2D map
3. create (see 3 ways under create)
4. iterate (see iterate section)
	- Note map is orderless (hashmap), so:
		for k, v := range m {
			...
		}
		k, v pair will arise in different order for multiple runs
5. get val by: map[key_val],
	- if key not exist, will be Zero value
	- v, exist := map[key_val], exist will be bool indicate if value exist
6. delete(map, key) for delete value
7. map allowed key type, it uses hash, thus key must allow equality comparison
- primitive data type except slice/map/function
- struct that doesn't contain slice, map, function (check during compilation)
8. see example code usage of map, find longest substring without repeating characters
e.g.
	- abcabcbb -> abc
	- abcdabcabcabcaa -> abcd
	- bbbbb -> b
9. rune is like super char in golang
10. strings. has many functions
*/

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) { // []rune(s) so all transferred to type that can consume many characters
		// avoid default 0 value, check exist
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i

	}
	return maxLength
}

/*
func main() {
	// create
	m := map[string]string{
		"name":    "wuji",
		"course":  "golang",
		"site":    "google",
		"quality": "not bad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil
	fmt.Println(m, m2, m3)

	// iterate
	for k, v := range m {
		fmt.Println(k, v)
	}

	// getting values, check if exist
	courseName, exist1 := m["course"]
	fmt.Println(courseName, exist1)
	notExist, exist2 := m["not exist"]
	fmt.Println(notExist, exist2)

	// delete
	n, ok := m["name"]
	fmt.Println(n, ok)
	delete(m, "name")
	n, ok = m["name"]
	fmt.Println(n, ok)

	fmt.Println(lengthOfNonRepeatingSubStr("abcd"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdabcabcabcaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("哈哈哈哈哈嘿黑哈"))

}*/
