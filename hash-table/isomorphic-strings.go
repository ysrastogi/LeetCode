func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m1 := make(map[byte]byte)
    m2 := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        c1, c2 := s[i], t[i]

        if v, ok := m1[c1]; ok && v != c2 {
            return false
        }
        if v, ok := m2[c2]; ok && v != c1 {
            return false
        }

        m1[c1] = c2
        m2[c2] = c1
    }

    return true
}
