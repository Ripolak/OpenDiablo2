package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Inventory struct {
	frame   *d2ui.Sprite
	panel   *d2ui.Sprite
	grid    *ItemGrid
	originX int
	originY int
	isOpen  bool
}

func NewInventory() *Inventory {
	originX := 400
	originY := 0
	return &Inventory{
		grid:    NewItemGrid(10, 4, originX+19, originY+320),
		originX: originX,
		originY: originY,
	}
}

func (g *Inventory) IsOpen() bool {
	return g.isOpen
}

func (g *Inventory) Toggle() {
	g.isOpen = !g.isOpen
}

func (g *Inventory) Open() {
	g.isOpen = true
}

func (g *Inventory) Close() {
	g.isOpen = false
}

func (g *Inventory) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.Frame, d2resource.PaletteSky)
	g.frame, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	g.panel, _ = d2ui.LoadSprite(animation)

	items := []InventoryItem{
		d2inventory.GetWeaponItemByCode("wnd"),
		d2inventory.GetWeaponItemByCode("sst"),
		d2inventory.GetWeaponItemByCode("jav"),
		d2inventory.GetArmorItemByCode("buc"),
		d2inventory.GetWeaponItemByCode("clb"),
		// TODO: Load the player's actual items
	}
	g.grid.Add(items...)
	g.grid.ChangeEquippedSlot("LeftHand", d2inventory.GetWeaponItemByCode("wnd"))
	g.grid.ChangeEquippedSlot("RightHand", d2inventory.GetWeaponItemByCode("wnd"))
	g.grid.ChangeEquippedSlot("Helmet", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("Chest", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("Shoes", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("Belt", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("Gloves", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("LeftRing", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("RightRing", d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot("Amulet", d2inventory.GetArmorItemByCode("buc"))
}

func (g *Inventory) Render(target d2render.Surface) {
	if !g.isOpen {
		return
	}

	x, y := g.originX, g.originY

	// Frame
	// Top left
	g.frame.SetCurrentFrame(5)
	w, h := g.frame.GetCurrentFrameSize()
	g.frame.SetPosition(x, y+h)
	g.frame.Render(target)
	x += w

	// Top right
	g.frame.SetCurrentFrame(6)
	w, h = g.frame.GetCurrentFrameSize()
	g.frame.SetPosition(x, y+h)
	g.frame.Render(target)
	x += w
	y += h

	// Right
	g.frame.SetCurrentFrame(7)
	w, h = g.frame.GetCurrentFrameSize()
	g.frame.SetPosition(x-w, y+h)
	g.frame.Render(target)
	y += h

	// Bottom right
	g.frame.SetCurrentFrame(8)
	w, h = g.frame.GetCurrentFrameSize()
	g.frame.SetPosition(x-w, y+h)
	g.frame.Render(target)
	x -= w

	// Bottom left
	g.frame.SetCurrentFrame(9)
	w, h = g.frame.GetCurrentFrameSize()
	g.frame.SetPosition(x-w, y+h)
	g.frame.Render(target)

	x, y = g.originX, g.originY
	y += 64

	// Panel
	// Top left
	g.panel.SetCurrentFrame(4)
	w, h = g.panel.GetCurrentFrameSize()
	g.panel.SetPosition(x, y+h)
	g.panel.Render(target)
	x += w

	// Top right
	g.panel.SetCurrentFrame(5)
	w, h = g.panel.GetCurrentFrameSize()
	g.panel.SetPosition(x, y+h)
	g.panel.Render(target)
	y += h

	// Bottom right
	g.panel.SetCurrentFrame(7)
	w, h = g.panel.GetCurrentFrameSize()
	g.panel.SetPosition(x, y+h)
	g.panel.Render(target)

	// Bottom left
	g.panel.SetCurrentFrame(6)
	w, h = g.panel.GetCurrentFrameSize()
	g.panel.SetPosition(x-w, y+h)
	g.panel.Render(target)

	g.grid.Render(target)
}
