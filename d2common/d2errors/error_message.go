package d2errors

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"image/color"
	"strings"
)

type ErrorMessage struct {
	ErrorText string
}

func (e *ErrorMessage) DisplayError(screen d2render.Surface) {
	var spriteLen = 263
	var x = 280
	var y = 175
	e.renderSprite(x, y, screen)
	e.renderLabel(x, y, spriteLen, screen)
	e.renderButton(x, spriteLen)
}

func (e *ErrorMessage) renderLabel(spriteX int, spriteY int, spriteLen int, screen d2render.Surface) {
	var label = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	lines := d2common.SplitIntoLinesWithMaxWidth("An error has occurred: "+e.ErrorText, 30)
	label.SetText(strings.Join(lines, "\n"))
	label.Alignment = d2ui.LabelAlignCenter
	var x, y = getLabelCoords(spriteX, spriteY, spriteLen)
	label.SetPosition(x, y)
	screen.DrawRect(800, 600, color.RGBA{A: 128})
	label.Render(screen)
}

func getLabelCoords(spriteX int, spriteY int, spriteLen int) (int, int) {
	var x = spriteX + (spriteLen / 2)
	var y = spriteY + 10
	return x, y
}

func (e *ErrorMessage) renderSprite(x int, y int, screen d2render.Surface) {
	var animation, _ = d2asset.LoadAnimation(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	var okCancelBox, _ = d2ui.LoadSprite(animation)
	okCancelBox.SetPosition(x, y)
	okCancelBox.RenderSegmented(screen, 2, 1, 0)
}

func (e *ErrorMessage) renderButton(spriteX int, spriteLen int) {
	var button = d2ui.CreateButton(d2ui.ButtonTypeOkCancel, "Ok")
	button.SetPosition(spriteX+(spriteLen/2)-47, 308)
	// TODO: Move the X calculation elsewhere. The 47 is because its half of 94, which is the length of the button.
	button.SetVisible(true)
	button.OnActivated(func() {})
	d2ui.AddWidget(&button)
}
