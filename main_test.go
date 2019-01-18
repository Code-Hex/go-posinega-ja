package main

import (
	"testing"
)

const (
	negative typ = iota
	positive
)

type typ int

func (t typ) String() string {
	if t == 0 {
		return "negative"
	}
	return "positive"
}

func IsExpected(t typ, score float64) bool {
	if t == positive {
		return score > 0
	}
	return score < 0
}

func TestDecision(t *testing.T) {
	cases := []struct {
		text       string
		expectType typ
	}{
		{
			text:       "拾う神あれば　捨てる神あり",
			expectType: negative,
		},
		{
			text:       "わしは劣化した。もうしょうがない。",
			expectType: negative,
		},
		{
			text:       "この鏡、どの角度から見ても気持ち悪い",
			expectType: negative,
		},
		{
			text:       "『新しい喜びは新しい苦痛をもたらす』モーツァルト",
			expectType: positive,
		},
		{
			text:       "友愛の多くは見せかけ。恋情の多くは愚かさであるにすぎない",
			expectType: negative,
		},
		{
			text:       "勝ちに不思議の勝ちあり、負けに不思議の負けなし。",
			expectType: negative,
		},
		{
			text:       "自分が立っている所を深く掘れ。そこからきっと泉が湧きでる。",
			expectType: positive,
		},
		{
			text:       "人生が終わってしまうことを恐れてはいけません。 人生がいつまでも始まらない事が怖いのです。",
			expectType: negative,
		},
		{
			text:       "今日行ないたい善行があれば、すぐに実行せよ。決して明日に延ばすな",
			expectType: positive,
		},
	}
	for _, c := range cases {
		got := Decision(c.text)
		if !IsExpected(c.expectType, got) {
			t.Errorf("Decision(%s) = %v, want %s", c.text, got, c.expectType)
		}
	}
}
