// Package jisx4061 implements [JIS X 4061] Japanese character string collation order.
// It is commonly referred to as "辞書順(the dictionary order)", "50音順(the syllabic order), "あいうえお順(the a-i-u-e-o order)".
//
// The collation method is simple collation (単純照合), where comparisons are made according to basic collation rules (基本照合規則).
// The Latin alphabet is processed, including macronised (マクロン付き文字) and circumflexed characters (サーカムフレックス付き文字).
// The extended Kanji character class (拡張漢字クラス) is used for the Kanji character class.
//
// [JIS X 4061]: https://ja.wikipedia.org/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E6%96%87%E5%AD%97%E5%88%97%E7%85%A7%E5%90%88%E9%A0%86%E7%95%AA
package jisx4061

import (
	"unicode/utf8"
)

//go:generate go run gen/main.go

type class int

const (
	classSpace      class = iota + 1 // スペース
	classDescriptor                  // 記述記号
	classBracket                     // 括弧記号
	classScience                     // 学術記号
	classGeneral                     // 一般記号
	classUnit                        // 単位記号
	classNumber                      // アラビア数字
	classSymbol                      // 欧字記号
	classAlphabet                    // ラテンアルファベット
	classKana                        // 仮名
	classKanji                       // 漢字
	classGeta                        // げた記号
)

type voiced int // 清濁

const (
	voicedNone       = iota
	voicedUnvoiced   // 清音
	voicedVoiced     // 濁音
	voicedSemivoiced // 半濁音
)

type symbolType int // 記号種別

const (
	symbolTypeNone      = iota
	symbolTypeLongVowel // 長音
	symbolTypeLower     // 小文字
	symbolTypeRepeat    // 繰返し記号
	symbolTypeUpper     // 大文字
)

type kanaType int // 仮名種別

const (
	kanaTypeNone     = iota
	kanaTypeHiragana // 平仮名
	kanaTypeKatakana // 片仮名
)

// ダイアクリティカルマーク
type diacriticalMark int

const (
	diacriticalMarkNone             diacriticalMark = iota // ダイアクリティカルマークなし
	diacriticalMarkMacron                                  // マクロン
	diacriticalMarkCircumflexAccent                        // サーカムフレックスアクセント
)

type letterCase int

const (
	letterCaseNone  letterCase = iota
	letterCaseLower            // 小文字
	letterCaseUpper            // 大文字
)

type attr struct {
	class           class
	order           int
	diacriticalMark diacriticalMark
	letterCase      letterCase
	voiced          voiced
	symbolType      symbolType
	kanaType        kanaType
}

var vowelTable = map[rune]rune{
	'あ': 'あ',
	'か': 'あ',
	'さ': 'あ',
	'た': 'あ',
	'な': 'あ',
	'は': 'あ',
	'ま': 'あ',
	'や': 'あ',
	'ら': 'あ',
	'わ': 'あ',
	'が': 'あ',
	'ざ': 'あ',
	'だ': 'あ',
	'ば': 'あ',
	'ぱ': 'あ',
	'ぁ': 'あ',
	'ゃ': 'あ',
	'ア': 'あ',
	'カ': 'あ',
	'サ': 'あ',
	'タ': 'あ',
	'ナ': 'あ',
	'ハ': 'あ',
	'マ': 'あ',
	'ヤ': 'あ',
	'ラ': 'あ',
	'ワ': 'あ',
	'ガ': 'あ',
	'ザ': 'あ',
	'ダ': 'あ',
	'バ': 'あ',
	'パ': 'あ',
	'ァ': 'あ',
	'ャ': 'あ',

	'い': 'い',
	'き': 'い',
	'し': 'い',
	'ち': 'い',
	'に': 'い',
	'ひ': 'い',
	'み': 'い',
	'り': 'い',
	'ゐ': 'い',
	'ぎ': 'い',
	'じ': 'い',
	'ぢ': 'い',
	'び': 'い',
	'ぴ': 'い',
	'ぃ': 'い',
	'イ': 'い',
	'キ': 'い',
	'シ': 'い',
	'チ': 'い',
	'ニ': 'い',
	'ヒ': 'い',
	'ミ': 'い',
	'リ': 'い',
	'ヰ': 'い',
	'ギ': 'い',
	'ジ': 'い',
	'ヂ': 'い',
	'ビ': 'い',
	'ピ': 'い',
	'ィ': 'い',

	'う': 'う',
	'く': 'う',
	'す': 'う',
	'つ': 'う',
	'ぬ': 'う',
	'ふ': 'う',
	'む': 'う',
	'ゆ': 'う',
	'る': 'う',
	'ぐ': 'う',
	'ず': 'う',
	'づ': 'う',
	'ぶ': 'う',
	'ぷ': 'う',
	'ぅ': 'う',
	'ゅ': 'う',
	'ウ': 'う',
	'ク': 'う',
	'ス': 'う',
	'ツ': 'う',
	'ヌ': 'う',
	'フ': 'う',
	'ム': 'う',
	'ユ': 'う',
	'ル': 'う',
	'グ': 'う',
	'ズ': 'う',
	'ヅ': 'う',
	'ブ': 'う',
	'プ': 'う',
	'ゥ': 'う',
	'ヴ': 'う',
	'ュ': 'う',

	'え': 'え',
	'け': 'え',
	'せ': 'え',
	'て': 'え',
	'ね': 'え',
	'へ': 'え',
	'め': 'え',
	'れ': 'え',
	'ゑ': 'え',
	'げ': 'え',
	'ぜ': 'え',
	'で': 'え',
	'べ': 'え',
	'ぺ': 'え',
	'ぇ': 'え',
	'エ': 'え',
	'ケ': 'え',
	'セ': 'え',
	'テ': 'え',
	'ネ': 'え',
	'ヘ': 'え',
	'メ': 'え',
	'レ': 'え',
	'ヱ': 'え',
	'ゲ': 'え',
	'ゼ': 'え',
	'デ': 'え',
	'ベ': 'え',
	'ペ': 'え',
	'ェ': 'え',

	'お': 'お',
	'こ': 'お',
	'そ': 'お',
	'と': 'お',
	'の': 'お',
	'ほ': 'お',
	'も': 'お',
	'よ': 'お',
	'ろ': 'お',
	'を': 'お',
	'ご': 'お',
	'ぞ': 'お',
	'ど': 'お',
	'ぼ': 'お',
	'ぽ': 'お',
	'ぉ': 'お',
	'ょ': 'お',
	'オ': 'お',
	'コ': 'お',
	'ソ': 'お',
	'ト': 'お',
	'ノ': 'お',
	'ホ': 'お',
	'モ': 'お',
	'ヨ': 'お',
	'ロ': 'お',
	'ヲ': 'お',
	'ゴ': 'お',
	'ゾ': 'お',
	'ド': 'お',
	'ボ': 'お',
	'ポ': 'お',
	'ォ': 'お',
	'ョ': 'お',

	'ん': 'ん',
	'ン': 'ん',
}

func getAttr(s string, offset int) (attr1 attr, n int) {
	for offset+n < len(s) {
		var ok bool
		r, m := utf8.DecodeRuneInString(s[offset+n:])
		n += m
		switch r {
		case 'ー':
			attr1 = table[r]
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n-m])
			if v, ok := vowelTable[last]; ok {
				attr1.order = table[v].order
			}
			return
		case 'ゝ', 'ゞ', 'ヽ', 'ヾ':
			attr1 = table[r]
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n-m])
			if last == 'ゝ' || last == 'ゞ' || last == 'ヽ' || last == 'ヾ' || last == 'ー' {
				return
			}
			attr0, ok := table[last]
			if !ok {
				return
			}
			attr1.order = attr0.order
			return
		}
		attr1, ok = table[r]
		if ok {
			return
		}

		// handle CJK Unified Ideographs
		if 0x4e00 <= r && r <= 0x10000 {
			attr1 = attr{
				class: classKanji,
				order: int(r),
			}
			return
		}
	}
	return
}

// Less compares the strings a and b according to JIS X 4061.
// if a < b it returns -1, if a > b it returns 1, and if a == b it returns 0.
func Compare(a, b string) int {
	var i, j int
	// log.Printf("checking %s < %s", a, b)
	// for i < len(a) && j < len(b) {
	// 	attrA, n := getAttr(a, i)
	// 	i += n
	// 	attrB, n := getAttr(b, j)
	// 	j += n

	// 	log.Printf("%#v", attrA)
	// 	log.Printf("%#v", attrB)
	// }

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.class != attrB.class {
			return compare(attrA.class, attrB.class)
		}
		if attrA.order != attrB.order {
			return compare(attrA.order, attrB.order)
		}
	}
	if i >= len(a) && j < len(b) {
		return -1
	}
	if i < len(a) && j >= len(b) {
		return 1
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.voiced != attrB.voiced {
			return compare(attrA.voiced, attrB.voiced)
		}
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.symbolType != attrB.symbolType {
			return compare(attrA.symbolType, attrB.symbolType)
		}
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.kanaType != attrB.kanaType {
			return compare(attrA.kanaType, attrB.kanaType)
		}
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.diacriticalMark != attrB.diacriticalMark {
			return compare(attrA.diacriticalMark, attrB.diacriticalMark)
		}
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.letterCase != attrB.letterCase {
			return compare(attrA.letterCase, attrB.letterCase)
		}
	}
	return 0
}

func compare[T ~int](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Less compares the strings a and b according to JIS X 4061 and returns the result a < b.
func Less(a, b string) bool {
	return Compare(a, b) < 0
}
