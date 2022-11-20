package jisx4061

import "unicode/utf8"

type class int

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
	letterCaseNone  = iota
	letterCaseLower // 小文字
	letterCaseUpper // 大文字
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

	'ん': 'ん',
	'ン': 'ん',
}

func getAttr(s string, offset int) (attr attr, n int) {
	for n < len(s) {
		var ok bool
		r, m := utf8.DecodeRuneInString(s[offset+n:])
		switch r {
		case 'ー':
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n])
			n += m
			if v, ok := vowelTable[last]; ok {
				attr = table[v]
			}
			return
		case 'ゝ', 'ゞ', 'ヽ', 'ヾ':
			last, _ := utf8.DecodeLastRuneInString(s[:offset+n])
			n += m
			if last == 'ゝ' || last == 'ゞ' || last == 'ヽ' || last == 'ヾ' || last == 'ー' {
				return
			}
			attr0, ok := table[last]
			if !ok {
				return
			}
			attr = attr0
			return
		}
		n += m
		attr, ok = table[r]
		if ok {
			return
		}
	}
	return
}

func Less(a, b string) bool {
	var i, j int
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
		if attrA.symbolType != attrB.symbolType {
			return attrA.symbolType < attrB.symbolType
		}
		if attrA.kanaType != attrB.kanaType {
			return attrA.kanaType < attrB.kanaType
		}
		if attrA.diacriticalMark != attrB.diacriticalMark {
			return attrA.diacriticalMark < attrB.diacriticalMark
		}
		if attrA.letterCase != attrB.letterCase {
			return attrA.letterCase < attrB.letterCase
		}
	}

	return i >= len(a) && j < len(b)
}
