package dlog

type ByKeys struct {
	Keys      []string
	SortOrder []string
}

func (s ByKeys) Len() int {
	return len(s.Keys)
}

func (s ByKeys) Swap(i, j int) {
	s.Keys[i], s.Keys[j] = s.Keys[j], s.Keys[i]
}

func (s ByKeys) Less(i, j int) bool {
	ii := IndexOf(s.SortOrder, s.Keys[i])
	ji := IndexOf(s.SortOrder, s.Keys[j])

	if ii == -1 && ji == -1 {
		return s.Keys[i] < s.Keys[j]
	}
	if ii == -1 {
		ii = len(s.SortOrder)
	}
	if ji == -1 {
		ji = len(s.SortOrder)
	}

	return ii < ji
}

func IndexOf(values []string, value string) int {
	for p, v := range values {
		if v == value {
			return p
		}
	}
	return -1
}
