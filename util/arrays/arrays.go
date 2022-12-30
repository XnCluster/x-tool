package arrays

func CopyArr(src []interface{}, srcPos int, dest []interface{}, destPos int, length int) {
	for i := srcPos; i < length; i++ {
		dest[destPos] = src[srcPos]
		srcPos++
		destPos++
	}
}
