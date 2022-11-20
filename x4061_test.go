package jisx4061

import (
	"testing"
)

func TestLess(t *testing.T) {
	tests := [][]string{
		// JIS X 4061-1996 参考2 適合性試験データ
		{
			"シャーレ", "シャイ", "シヤィ", "シャレ",
			"ちょこ", "ちよこ", "チョコレート",
			"てーた", "テータ", "テェタ", "てえた", "でーた", "データ", "デェタ", "でえた",
		},
	}
	for _, tt := range tests {
		for i, a := range tt {
			for j, b := range tt {
				got := Less(a, b)
				want := i < j
				if got != want {
					t.Errorf("want %s < %s is %t, but not", a, b, want)
				}
			}
		}
	}
}
