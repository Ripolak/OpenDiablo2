package d2gamescreen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
)

type GuiTestMain struct {
	renderer d2interface.Renderer
}

func CreateGuiTestMain(renderer d2interface.Renderer) *GuiTestMain {
	return &GuiTestMain{
		renderer: renderer,
	}
}

func (g *GuiTestMain) OnLoad(loading d2screen.LoadingState) {
	layout := d2gui.CreateLayout(g.renderer, d2gui.PositionTypeHorizontal)

	loading.Progress(0.3)
	//
	layoutLeft := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutLeft.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	_, _ = layoutLeft.AddLabel("FontStyle16Units", d2gui.FontStyle16Units)
	layoutLeft.AddSpacerStatic(0, 100)
	_, _ = layoutLeft.AddLabel("FontStyle30Units", d2gui.FontStyle30Units)
	_, _ = layoutLeft.AddLabel("FontStyle42Units", d2gui.FontStyle42Units)
	_, _ = layoutLeft.AddLabel("FontStyleFormal10Static", d2gui.FontStyleFormal10Static)
	_, _ = layoutLeft.AddLabel("FontStyleFormal11Units", d2gui.FontStyleFormal11Units)
	_, _ = layoutLeft.AddLabel("FontStyleFormal12Static", d2gui.FontStyleFormal12Static)
	loading.Progress(0.6)

	layout.AddSpacerDynamic()

	layoutRight := layout.AddLayout(d2gui.PositionTypeVertical)
	layoutRight.SetHorizontalAlign(d2gui.HorizontalAlignRight)
	_, _ = layoutRight.AddButton("Medium", d2gui.ButtonStyleMedium)
	_, _ = layoutRight.AddButton("Narrow", d2gui.ButtonStyleNarrow)
	_, _ = layoutRight.AddButton("OkCancel", d2gui.ButtonStyleOkCancel)
	_, _ = layoutRight.AddButton("Short", d2gui.ButtonStyleShort)
	_, _ = layoutRight.AddButton("Wide", d2gui.ButtonStyleWide)
	loading.Progress(0.9)

	layout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	d2gui.SetLayout(layout)
}

func (g *GuiTestMain) Render(_ d2interface.Surface) error {
	return nil
}

func (g *GuiTestMain) Advance(_ float64) error {
	return nil
}
