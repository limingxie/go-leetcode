package main

func IsValid(s string) bool {
	if len(s) == 0 || s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}

	var ss []byte
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			ss = append(ss, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(ss) == 0 {
				return false
			}

			if (ss[len(ss)-1] == '(' && s[i] == ')') || (ss[len(ss)-1] == '[' && s[i] == ']') || (ss[len(ss)-1] == '{' && s[i] == '}') {
				if len(ss) == 1 {
					ss = nil
				} else {
					ss = ss[:len(ss)-1]
				}
			} else {
				return false
			}
		}
	}

	return len(ss) == 0
}
