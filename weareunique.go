package piscine

func WeAreUnique(str1 , str2 string) int {
	count:=0
	if str1=="" && str2 == ""{
		return -1
	}
	for i:= range str1{
		if str1[i] != str2[i]{
			count++
		}
	}
	return count
}