package formatting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.kevinchen-verkada/helloworld/formatting"
)

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
