package formatting

import (
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
)

func FormatOutput(output string, useBox bool) string {
	msg := strings.ToUpper(output)

	if !useBox {
		return msg
	}

	Box := box.New(box.Config{Px: 1, Py: 1, Type: "Single", Color: "Cyan"})
	return Box.String("Verkada Codelab", msg)
}
