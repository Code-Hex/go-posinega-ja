package main

func IsStopWord(s string) bool {
	_, ok := stopWords[s]
	return ok
}

var stopWords = map[string]struct{}{
	"あそこ":  struct{}{},
	"あたり":  struct{}{},
	"あちら":  struct{}{},
	"あっち":  struct{}{},
	"あと":   struct{}{},
	"あな":   struct{}{},
	"あなた":  struct{}{},
	"あれ":   struct{}{},
	"いくつ":  struct{}{},
	"いつ":   struct{}{},
	"いま":   struct{}{},
	"いや":   struct{}{},
	"いろいろ": struct{}{},
	"うち":   struct{}{},
	"おおまか": struct{}{},
	"おまえ":  struct{}{},
	"おれ":   struct{}{},
	"がい":   struct{}{},
	"かく":   struct{}{},
	"かたち":  struct{}{},
	"かやの":  struct{}{},
	"から":   struct{}{},
	"がら":   struct{}{},
	"きた":   struct{}{},
	"くせ":   struct{}{},
	"ここ":   struct{}{},
	"こっち":  struct{}{},
	"こと":   struct{}{},
	"ごと":   struct{}{},
	"こちら":  struct{}{},
	"ごっちゃ": struct{}{},
	"これ":   struct{}{},
	"これら":  struct{}{},
	"ごろ":   struct{}{},
	"さまざま": struct{}{},
	"さらい":  struct{}{},
	"さん":   struct{}{},
	"しかた":  struct{}{},
	"しよう":  struct{}{},
	"すか":   struct{}{},
	"ずつ":   struct{}{},
	"すね":   struct{}{},
	"すべて":  struct{}{},
	"ぜんぶ":  struct{}{},
	"そう":   struct{}{},
	"そこ":   struct{}{},
	"そちら":  struct{}{},
	"そっち":  struct{}{},
	"そで":   struct{}{},
	"それ":   struct{}{},
	"それぞれ": struct{}{},
	"それなり": struct{}{},
	"たくさん": struct{}{},
	"たち":   struct{}{},
	"たび":   struct{}{},
	"ため":   struct{}{},
	"だめ":   struct{}{},
	"ちゃ":   struct{}{},
	"ちゃん":  struct{}{},
	"てん":   struct{}{},
	"とおり":  struct{}{},
	"とき":   struct{}{},
	"どこ":   struct{}{},
	"どこか":  struct{}{},
	"ところ":  struct{}{},
	"どちら":  struct{}{},
	"どっか":  struct{}{},
	"どっち":  struct{}{},
	"どれ":   struct{}{},
	"なか":   struct{}{},
	"なかば":  struct{}{},
	"なに":   struct{}{},
	"など":   struct{}{},
	"なん":   struct{}{},
	"はじめ":  struct{}{},
	"はず":   struct{}{},
	"はるか":  struct{}{},
	"ひと":   struct{}{},
	"ひとつ":  struct{}{},
	"ふく":   struct{}{},
	"ぶり":   struct{}{},
	"べつ":   struct{}{},
	"へん":   struct{}{},
	"ぺん":   struct{}{},
	"ほう":   struct{}{},
	"ほか":   struct{}{},
	"まさ":   struct{}{},
	"まし":   struct{}{},
	"まとも":  struct{}{},
	"まま":   struct{}{},
	"みたい":  struct{}{},
	"みつ":   struct{}{},
	"みなさん": struct{}{},
	"みんな":  struct{}{},
	"もと":   struct{}{},
	"もの":   struct{}{},
	"もん":   struct{}{},
	"やつ":   struct{}{},
	"よう":   struct{}{},
	"よそ":   struct{}{},
	"わけ":   struct{}{},
	"わたし":  struct{}{},
	"ハイ":   struct{}{},
	"上":    struct{}{},
	"中":    struct{}{},
	"下":    struct{}{},
	"字":    struct{}{},
	"年":    struct{}{},
	"月":    struct{}{},
	"日":    struct{}{},
	"時":    struct{}{},
	"分":    struct{}{},
	"秒":    struct{}{},
	"週":    struct{}{},
	"火":    struct{}{},
	"水":    struct{}{},
	"木":    struct{}{},
	"金":    struct{}{},
	"土":    struct{}{},
	"国":    struct{}{},
	"都":    struct{}{},
	"道":    struct{}{},
	"府":    struct{}{},
	"県":    struct{}{},
	"市":    struct{}{},
	"区":    struct{}{},
	"町":    struct{}{},
	"村":    struct{}{},
	"各":    struct{}{},
	"第":    struct{}{},
	"方":    struct{}{},
	"何":    struct{}{},
	"的":    struct{}{},
	"度":    struct{}{},
	"文":    struct{}{},
	"者":    struct{}{},
	"性":    struct{}{},
	"体":    struct{}{},
	"人":    struct{}{},
	"他":    struct{}{},
	"今":    struct{}{},
	"部":    struct{}{},
	"課":    struct{}{},
	"係":    struct{}{},
	"外":    struct{}{},
	"類":    struct{}{},
	"達":    struct{}{},
	"気":    struct{}{},
	"室":    struct{}{},
	"口":    struct{}{},
	"誰":    struct{}{},
	"用":    struct{}{},
	"界":    struct{}{},
	"会":    struct{}{},
	"首":    struct{}{},
	"男":    struct{}{},
	"女":    struct{}{},
	"別":    struct{}{},
	"話":    struct{}{},
	"私":    struct{}{},
	"屋":    struct{}{},
	"店":    struct{}{},
	"家":    struct{}{},
	"場":    struct{}{},
	"等":    struct{}{},
	"見":    struct{}{},
	"際":    struct{}{},
	"観":    struct{}{},
	"段":    struct{}{},
	"略":    struct{}{},
	"例":    struct{}{},
	"系":    struct{}{},
	"論":    struct{}{},
	"形":    struct{}{},
	"間":    struct{}{},
	"地":    struct{}{},
	"員":    struct{}{},
	"線":    struct{}{},
	"点":    struct{}{},
	"書":    struct{}{},
	"品":    struct{}{},
	"力":    struct{}{},
	"法":    struct{}{},
	"感":    struct{}{},
	"作":    struct{}{},
	"元":    struct{}{},
	"手":    struct{}{},
	"数":    struct{}{},
	"彼":    struct{}{},
	"彼女":   struct{}{},
	"子":    struct{}{},
	"内":    struct{}{},
	"楽":    struct{}{},
	"喜":    struct{}{},
	"怒":    struct{}{},
	"哀":    struct{}{},
	"輪":    struct{}{},
	"頃":    struct{}{},
	"化":    struct{}{},
	"境":    struct{}{},
	"俺":    struct{}{},
	"奴":    struct{}{},
	"高":    struct{}{},
	"校":    struct{}{},
	"婦":    struct{}{},
	"伸":    struct{}{},
	"紀":    struct{}{},
	"誌":    struct{}{},
	"レ":    struct{}{},
	"行":    struct{}{},
	"列":    struct{}{},
	"事":    struct{}{},
	"士":    struct{}{},
	"台":    struct{}{},
	"集":    struct{}{},
	"様":    struct{}{},
	"所":    struct{}{},
	"歴":    struct{}{},
	"器":    struct{}{},
	"名":    struct{}{},
	"情":    struct{}{},
	"連":    struct{}{},
	"毎":    struct{}{},
	"式":    struct{}{},
	"簿":    struct{}{},
	"回":    struct{}{},
	"匹":    struct{}{},
	"個":    struct{}{},
	"席":    struct{}{},
	"束":    struct{}{},
	"歳":    struct{}{},
	"目":    struct{}{},
	"通":    struct{}{},
	"面":    struct{}{},
	"円":    struct{}{},
	"玉":    struct{}{},
	"枚":    struct{}{},
	"前":    struct{}{},
	"後":    struct{}{},
	"左":    struct{}{},
	"右":    struct{}{},
	"次":    struct{}{},
	"先":    struct{}{},
	"春":    struct{}{},
	"夏":    struct{}{},
	"秋":    struct{}{},
	"冬":    struct{}{},
	"一":    struct{}{},
	"二":    struct{}{},
	"三":    struct{}{},
	"四":    struct{}{},
	"五":    struct{}{},
	"六":    struct{}{},
	"七":    struct{}{},
	"八":    struct{}{},
	"九":    struct{}{},
	"十":    struct{}{},
	"百":    struct{}{},
	"千":    struct{}{},
	"万":    struct{}{},
	"億":    struct{}{},
	"兆":    struct{}{},
	"下記":   struct{}{},
	"上記":   struct{}{},
	"時間":   struct{}{},
	"今回":   struct{}{},
	"前回":   struct{}{},
	"場合":   struct{}{},
	"一つ":   struct{}{},
	"年生":   struct{}{},
	"自分":   struct{}{},
	"ヶ所":   struct{}{},
	"ヵ所":   struct{}{},
	"カ所":   struct{}{},
	"箇所":   struct{}{},
	"ヶ月":   struct{}{},
	"ヵ月":   struct{}{},
	"カ月":   struct{}{},
	"箇月":   struct{}{},
	"名前":   struct{}{},
	"本当":   struct{}{},
	"確か":   struct{}{},
	"時点":   struct{}{},
	"全部":   struct{}{},
	"関係":   struct{}{},
	"近く":   struct{}{},
	"方法":   struct{}{},
	"我々":   struct{}{},
	"違い":   struct{}{},
	"多く":   struct{}{},
	"扱い":   struct{}{},
	"新た":   struct{}{},
	"その後":  struct{}{},
	"半ば":   struct{}{},
	"結局":   struct{}{},
	"様々":   struct{}{},
	"以前":   struct{}{},
	"以後":   struct{}{},
	"以降":   struct{}{},
	"未満":   struct{}{},
	"以上":   struct{}{},
	"以下":   struct{}{},
	"幾つ":   struct{}{},
	"毎日":   struct{}{},
	"自体":   struct{}{},
	"向こう":  struct{}{},
	"何人":   struct{}{},
	"手段":   struct{}{},
	"同じ":   struct{}{},
	"感じ":   struct{}{},
}
