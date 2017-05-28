package main

func StrMapKeys (m map[string]string) ([]string){
	i := 0
	keys := make([]string, len(m))

	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
