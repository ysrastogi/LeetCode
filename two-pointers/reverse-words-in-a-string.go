func splitWords(sentence string) []string {
	words := []string{}
	word := ""

	for _, ch := range sentence {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}

func reverseWords(s string) string {
    words := splitWords(s)
    res := ""
    for i := len(words) - 1; i >= 0; i-- {
	res += words[i]
    if i != 0{
        res += " "
    }
    
    }
    return res
    
}