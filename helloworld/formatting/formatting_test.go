package formatting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.kevinchen-verkada/helloworld/formatting"
)

func TestFormattingUppercase(t *testing.T) {
	expected := "HELLO THERE"
	actual := formatting.FormatOutput("Hello there", false /* useBox */)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestFormattingUppercaseWithBox(t *testing.T) {
	expected := `┌─────────────────┐
│                 │
│ Verkada Codelab │
│                 │
│ HELLO THERE     │
│                 │
└─────────────────┘
`
	actual := formatting.FormatOutput("Hello there", true /* useBox */)
	assert.Equal(t, expected, actual, "strings should be equal")

}
