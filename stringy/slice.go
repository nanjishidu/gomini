package stringy

//https://github.com/nsqio/nsq/blob/master/internal/stringy/slice.go
func Add(s []string, a string) []string {
	for _, existing := range s {
		if a == existing {
			return s
		}
	}
	return append(s, a)

}

func Union(s []string, a []string) []string {
	for _, entry := range a {
		found := false
		for _, existing := range s {
			if entry == existing {
				found = true
				break
			}
		}
		if !found {
			s = append(s, entry)
		}
	}
	return s
}

func Uniq(s []string) (r []string) {
	for _, entry := range s {
		found := false
		for _, existing := range r {
			if existing == entry {
				found = true
				break
			}
		}
		if !found {
			r = append(r, entry)
		}
	}
	return
}
func InSlice(s []string, a string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			return true
		}
	}
	return false
}
func AddInt64(s []int64, a int64) []int64 {
	for _, existing := range s {
		if a == existing {
			return s
		}
	}
	return append(s, a)
}
