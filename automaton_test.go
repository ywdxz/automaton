package automaton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck1(t *testing.T) {

	var tests = []struct {
		src      []string
		dest     []byte
		expected []CheckResult
	}{
		{
			[]string{"12321", "abc", "ffdsfs"},
			[]byte("11abc22ffdsfs"),
			[]CheckResult{{2, 5, 1}, {7, 13, 2}},
		},
		{
			[]string{"ab", "abc", "abcd"},
			[]byte("abcd"),
			[]CheckResult{{0, 2, 0}, {0, 3, 1}, {0, 4, 2}},
		},
		{
			[]string{"abcd", "abc", "ab"},
			[]byte("abcd"),
			[]CheckResult{{0, 2, 2}, {0, 3, 1}, {0, 4, 0}},
		},
		{ //bug?
			[]string{"cd", "bcd", "abcd"},
			[]byte("abcd"),
			[]CheckResult{{0, 4, 2}},
		},
		{ //bug?
			[]string{"abcd", "bcd", "cd"},
			[]byte("abcd"),
			[]CheckResult{{1, 4, 0}},
		},
		{
			[]string{"ab", "abc", "abcd"},
			[]byte("acbcd"),
			nil,
		},
	}

	for _, test := range tests {
		auto := NewAutomaton(test.src)
		results := auto.Check(test.dest)

		if !assert.Equal(t, len(results), len(test.expected)) {
			t.Fatalf("in line with expectations,{%+v}\n", test)
		}

		for i := 0; i < len(results); i++ {
			if !assert.Equal(t, results[i], test.expected[i]) {
				t.Fatalf("in line with expectations,{%+v}\n", test)
			}

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
