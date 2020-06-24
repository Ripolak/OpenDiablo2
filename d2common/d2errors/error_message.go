package d2errors

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type ErrorMessage struct {
	ErrorText string
}

func (e *ErrorMessage) DisplayError(target d2render.Surface) {
	var label = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	label.SetText(e.ErrorText)
	label.Alignment = d2ui.LabelAlignCenter
	label.Render(target)
}
