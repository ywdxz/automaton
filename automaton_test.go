package automaton

import (
	"fmt"
	"testing"
)

func TestCheck1(t *testing.T) {

	var tests = []struct {
		src  []string
		dest []byte
	}{
		{
			[]string{"12321", "abc", "ffdsfs"},
			[]byte("11abc22ffdsfs"),
		},
		{
			[]string{"ab", "abc", "abcd"},
			[]byte("abcd"),
		},
		{
			[]string{"ab", "abc", "abcd"},
			[]byte("acbcd"),
		},
	}

	for _, test := range tests {
		auto := NewAutomaton(test.src)
		// auto.Print()
		results := auto.Check(test.dest)

		fmt.Println("-----")
		for _, result := range results {
			fmt.Printf("%d-(%d:%d) - %s \n", result.TokenID, result.StartIndex, result.EndIndex, test.dest[result.StartIndex:result.EndIndex])
		}
	}
}

func BenchmarkTemplateReplace(b *testing.B) {
	text := "{{.Ywdxz}} is {{.Count}} of {{.Material}}"
	auto := NewAutomaton([]string{"{{.Material}}", "{{.Count}}", "{{.Ywdxz}}"})

	Text := text
	for i := 0; i < 10000; i++ {
		Text += text
	}
	src := []byte(Text)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		auto.Check(src)
	}
}
