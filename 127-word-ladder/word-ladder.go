type node struct {
    word  string
    level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

    wordSet := make(map[string]struct{})
    for _, w := range wordList {
        wordSet[w] = struct{}{}
    }

    if _, ok := wordSet[endWord]; !ok {
        return 0
    }

    q := []node{{beginWord, 1}}

    for len(q) > 0 {

        curr := q[0]
        q = q[1:]

        wordBytes := []byte(curr.word)

        for i := 0; i < len(wordBytes); i++ {

            original := wordBytes[i]

            for c := byte('a'); c <= byte('z'); c++ {

                wordBytes[i] = c
                newWord := string(wordBytes)

                if newWord == endWord {
                    return curr.level + 1
                }

                if _, exists := wordSet[newWord]; exists {
                    delete(wordSet, newWord)
                    q = append(q, node{newWord, curr.level + 1})
                }
            }

            wordBytes[i] = original
        }
    }

    return 0
}