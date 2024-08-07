package ransomnote

import "slices"

func canConstruct(ransomNote string, magazine string) bool {
	rn := []byte(ransomNote)
	ma := []byte(magazine)
	slices.Sort(rn)
	slices.Sort(ma)

	c := 0
	for i := 0; i < len(ma); i++ {
		if ma[i] == rn[c] {
			c++
		}
		if c == len(rn) {
			return true
		}
	}
	return false
}
