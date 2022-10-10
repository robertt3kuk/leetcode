func reverseWords(s string) string {
   str := ""
	s = " " + s
	res := ""
	for _, c := range s {
		if string(c) != " " {
			str = string(c) + str
		} else {
			res = res + str
			str = string(c)
		}
	}
	res = res + str

	return res[:len(res)-1] 
}

    