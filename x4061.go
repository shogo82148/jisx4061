package jisx4061

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

type class int

func Print() {
	var runes = []rune{}
	for r := range table {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		if table[runes[i]].class != table[runes[j]].class {
			return table[runes[i]].class < table[runes[j]].class
		}
		if table[runes[i]].order != table[runes[j]].order {
			return table[runes[i]].order < table[runes[j]].order
		}
		if table[runes[i]].voiced != table[runes[j]].voiced {
			return table[runes[i]].voiced < table[runes[j]].voiced
		}
		if table[runes[i]].symbolType != table[runes[j]].symbolType {
			return table[runes[i]].symbolType < table[runes[j]].symbolType
		}
		if table[runes[i]].kanaType != table[runes[j]].kanaType {
			return table[runes[i]].kanaType < table[runes[j]].kanaType
		}
		if table[runes[i]].diacriticalMark != table[runes[j]].diacriticalMark {
			return table[runes[i]].diacriticalMark < table[runes[j]].diacriticalMark
		}
		if table[runes[i]].letterCase != table[runes[j]].letterCase {
			return table[runes[i]].letterCase < table[runes[j]].letterCase
		}
		return runes[i] < runes[j]
	})

	fmt.Println("文字\t文字クラス\t番号\tダイアクリティカルマーク\t大小\t清濁\t記号種別\t仮名種別")
	for _, r := range runes {
		attr := table[r]
		fmt.Printf("%c\t", r)
		switch attr.class {
		case classSpace:
			fmt.Printf("スペース")
		case classDescriptor:
			fmt.Printf("記述記号")
		case classBracket:
			fmt.Printf("括弧記号")
		case classScience:
			fmt.Printf("学術記号")
		case classGeneral:
			fmt.Printf("一般記号")
		case classUnit:
			fmt.Printf("単位記号")
		case classNumber:
			fmt.Printf("アラビア数字")
		case classSymbol:
			fmt.Printf("欧字記号")
		case classAlphabet:
			fmt.Printf("ラテンアルファベット")
		case classKana:
			fmt.Printf("仮名")
		case classKanji:
			fmt.Printf("漢字")
		case classGeta:
			fmt.Printf("げた記号")
		}

		fmt.Printf("\t%d", attr.order)

		fmt.Printf("\t")
		if attr.letterCase != letterCaseNone {
			switch attr.diacriticalMark {
			case diacriticalMarkNone:
				fmt.Print("ダイアクリティカルマークなし")
			case diacriticalMarkMacron:
				fmt.Print("マクロン付き")
			case diacriticalMarkCircumflexAccent:
				fmt.Print("サーカムフレックスアクセント付き")
			}
		}

		fmt.Printf("\t")
		switch attr.letterCase {
		case letterCaseLower:
			fmt.Print("大文字")
		case letterCaseUpper:
			fmt.Print("小文字")
		}

		fmt.Printf("\t")
		switch attr.voiced {
		case voicedUnvoiced:
			fmt.Print("清音")
		case voicedVoiced:
			fmt.Print("濁音")
		case voicedSemivoiced:
			fmt.Print("半濁音")
		}

		fmt.Printf("\t")
		switch attr.symbolType {
		case symbolTypeUpper:
			fmt.Print("大文字")
		case symbolTypeLower:
			fmt.Print("小文字")
		case symbolTypeRepeat:
			fmt.Print("繰返し記号")
		case symbolTypeLongVowel:
			fmt.Print("長音記号")
		}

		fmt.Print("\t")
		switch attr.kanaType {
		case kanaTypeHigagana:
			fmt.Print("平仮名")
		case kanaTypeKatakana:
			fmt.Print("片仮名")
		}

		fmt.Println()
	}
}

const (
	classSpace      class = iota + 1 // スペース
	classDescriptor                  // 記述記号
	classBracket                     // 括弧
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
	symbolTypeRepeat    // 繰りし記号
	symbolTypeUpper     // 大文字
)

type kanaType int // 仮名種別

const (
	kanaTypeNone = iota
	kanaTypeHigagana
	kanaTypeKatakana
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
	voiced          voiced
	symbolType      symbolType
	kanaType        kanaType
	diacriticalMark diacriticalMark
	letterCase      letterCase
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

func getAttr(s string, offset int) (attr attr, n int) {
	for offset+n < len(s) {
		var ok bool
		r, m := utf8.DecodeRuneInString(s[offset+n:])
		n += m
		switch r {
		case 'ー':
			attr = table[r]
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n-m])
			if v, ok := vowelTable[last]; ok {
				attr.order = table[v].order
			}
			return
		case 'ゝ', 'ゞ', 'ヽ', 'ヾ':
			attr = table[r]
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n-m])
			if last == 'ゝ' || last == 'ゞ' || last == 'ヽ' || last == 'ヾ' || last == 'ー' {
				return
			}
			attr0, ok := table[last]
			if !ok {
				return
			}
			attr.order = attr0.order
			return
		}
		attr, ok = table[r]
		if ok {
			return
		}
	}
	return
}

func Less(a, b string) bool {
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
			return attrA.class < attrB.class
		}
		if attrA.order != attrB.order {
			return attrA.order < attrB.order
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.voiced != attrB.voiced {
			return attrA.voiced < attrB.voiced
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.symbolType != attrB.symbolType {
			return attrA.symbolType < attrB.symbolType
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.kanaType != attrB.kanaType {
			return attrA.kanaType < attrB.kanaType
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.diacriticalMark != attrB.diacriticalMark {
			return attrA.diacriticalMark < attrB.diacriticalMark
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}

	i, j = 0, 0
	for i < len(a) && j < len(b) {
		attrA, n := getAttr(a, i)
		i += n
		attrB, n := getAttr(b, j)
		j += n

		if attrA.letterCase != attrB.letterCase {
			return attrA.letterCase < attrB.letterCase
		}
	}
	if i >= len(a) && j < len(b) {
		return true
	}
	if i < len(a) && j >= len(b) {
		return false
	}
	return false
}
