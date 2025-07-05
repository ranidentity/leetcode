package services

//mississippi	"mis*is*ip*."
//ab a**
// "." any cahracter , "*" multiple
func IsMatch2(s string, p string) bool {
	if p == ".*" {
		return true
	}
	var buf []string
	var part string
	i := 0
	for len(p) > i {
		if p[i] == '.' && p[i+1] != '*' && i < len(p)-1 {
			buf = append(buf, string(p[i]))
		} else if p[i] != '.' && p[i+1] != '*' && i < len(p)-1 {
			part += string(p[i])
		} else {
			if part != "" {
				buf = append(buf, part) // clear off part
				part = ""
			}
			// check combo
			if p[i+1] == '*' && i < len(p)-1 {
				buf = append(buf, string(p[i]+'*'))
				i++
			}
		}
		i++
	}
	flag := true
	for _, i := range buf {
		for _, j := range i {
		}
		if s[:len(i)] == i {
			continue
		} // mis s*

	}
	return flag
}

func IsMatch(s string, p string) bool {
	if p == ".*" || p == "*." {
		return true
	}
	mode := 0
	var key rune
	for index, r := range p {
		if index == 0 {
			key = r
		}
		if r == '.' {
			mode = 1
		}
		if r == '*' {
			mode = 2
		}
	}

	if mode == 0 && s != p {
		return false
	}
	for _, i := range s {
		if i != key && mode == 1 {
			return false
		}
	}
	return true
}
