package playground

// ReverseString takes a string and returns it reversed.
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"

func ReverseString(s string) string {
	result := ""
	for i := len(s)-1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func ReverseString2(s string) string {
	var result string
	for _, i := range s {
		result = string(i) + result
	}
	return result
}

func ReverseString3(s string) string {
	// Convert the string into a slice of runes (Unicode code points)
    runes := []rune(s)
    
    // Initialize two pointers:
    // i starts at the beginning (0)
    // j starts at the end (len(runes)-1)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        // Swap the runes at positions i and j
        runes[i], runes[j] = runes[j], runes[i]
    }
    
    // Convert the slice of runes back to a string
    return string(runes)
}