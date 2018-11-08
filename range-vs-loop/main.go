package main

func main() {
	str := getString(10)
	slice := getSlice(10)

	RangeString(str)
	LoopString(str)
	RangeSlice(slice)
	LoopSlice(slice)
}

// getString returns string of given length
func getString(length int) string {
	char := "a"
	var str string
	for i := 0; i < length; i++ {
		str = str + char
	}
	return str
}

func getSlice(length int) []byte {
	elem := byte(01)
	slice := []byte{}

	for i := 0; i < length; i++ {
		slice = append(slice, elem)
	}
	return slice
}

// RangeString ranges over string
func RangeString(str string) {
	for _, val := range str {
		_ = val
	}
}

// RangeIndexString ranges over string but uses only index
func RangeIndexString(str string) {
	for i := range str {
		_ = str[i]
	}
}

// LoopString loops over string
func LoopString(str string) {
	length := len(str)
	for i := 0; i < length; i++ {
		_ = str[i]
	}
}

// RangeSlice ranges over slice
func RangeSlice(slice []byte) {
	for _, val := range slice {
		_ = val
	}
}

// LoopSlice loops over slice
func LoopSlice(slice []byte) {
	length := len(slice)
	for i := 0; i < length; i++ {
		_ = slice[i]
	}
}
