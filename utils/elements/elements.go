package elements

func GetFirstElementIfExists[T any](str ...T) T {
	var res T
	if len(str) > 0 {
		return str[0]
	}
	return res
}

func GetFirstElementOrDefault[T any](def T, vals ...T) T {
	if len(vals) > 0 {
		return vals[0]
	}
	return def
}

func IsElementExists[T any](index int, arr ...T) bool {
	return len(arr) == index+1
}

func IsLastElement[T any](index int, arr ...T) bool {
	lastIndex := len(arr)
	return index == lastIndex
}
