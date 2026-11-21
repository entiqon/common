package testkit

import "strings"

// Diff returns a line diff from want to got. Lines only in want are prefixed
// with "- ", lines only in got with "+ ", shared lines with two spaces.
// It returns "" when the inputs are equal.
func Diff(want, got string) string {
	if want == got {
		return ""
	}
	a := strings.Split(want, "\n")
	b := strings.Split(got, "\n")

	// Longest common subsequence table.
	n, m := len(a), len(b)
	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i] == b[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else if lcs[i+1][j] >= lcs[i][j+1] {
				lcs[i][j] = lcs[i+1][j]
			} else {
				lcs[i][j] = lcs[i][j+1]
			}
		}
	}

	var sb strings.Builder
	i, j := 0, 0
	for i < n || j < m {
		switch {
		case i < n && j < m && a[i] == b[j]:
			sb.WriteString("  " + a[i] + "\n")
			i++
			j++
		case j < m && (i == n || lcs[i][j+1] > lcs[i+1][j]):
			sb.WriteString("+ " + b[j] + "\n")
			j++
		default:
			sb.WriteString("- " + a[i] + "\n")
			i++
		}
	}
	return sb.String()
}
