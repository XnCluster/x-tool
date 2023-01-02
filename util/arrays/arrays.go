package arrays

func CopyArr(src []interface{}, srcPos int, dest []interface{}, destPos int, length int) {
	for i := srcPos; i < length+srcPos; i++ {
		dest[destPos] = src[i]
		destPos++
	}
}
