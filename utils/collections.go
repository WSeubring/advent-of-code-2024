package utils

func RemoveIndexFromArr[Type any](s []Type, index int) []Type {
	left := append([]Type{}, s[:index]...)
	right := append([]Type{}, s[index+1:]...)
	return append(left, right...)
}

func InsertAtIndex[Type any](s []Type, index int, value Type) []Type {
	left := append([]Type{}, s[:index]...)
	right := append([]Type{}, s[index:]...)
	return append(append(left, value), right...)
}
