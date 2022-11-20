package jisx4061

import "sort"

// StringSlice implements [sort.Interface].
type StringSlice []string

// Len implements [sort.Interface].
func (s StringSlice) Len() int { return len(s) }

// Swap implements [sort.Interface].
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less implements [sort.Interface].
func (s StringSlice) Less(i, j int) bool { return Less(s[i], s[j]) }

// Sort sorts s.
func Sort(s []string) {
	sort.Sort(StringSlice(s))
}

// Stable sorts s.
func Stable(s []string) {
	sort.Stable(StringSlice(s))
}

// IsSorted reports whether data is sorted.
func IsSorted(data []string) bool {
	return sort.IsSorted(StringSlice(data))
}
