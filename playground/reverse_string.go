package playground

// ReverseString takes a string and returns it reversed.
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"

func ReverseString(s string) string {
	result := ""
	for i := len(s) -1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}