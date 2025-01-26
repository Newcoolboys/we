package we

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	_ "github.com/df-mc/we/act"
	"github.com/df-mc/we/brush"
	"github.com/df-mc/we/palette"
	"github.com/go-gl/mathgl/mgl64"
)

type Handler struct {
	player.NopHandler
	ph *palette.Handler
	bh *brush.Handler
}

func NewHandler(p *player.Player) *Handler {
	return &Handler{ph: palette.NewHandler(p), bh: brush.NewHandler(p)}
}

func (h *Handler) HandleItemUse(ctx *player.Context) {
	h.bh.HandleItemUse(ctx)
}

func (h *Handler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, vec mgl64.Vec3) {
	h.ph.HandleItemUseOnBlock(ctx, pos, face, vec)
}

func (h *Handler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	h.ph.HandleBlockBreak(ctx, pos, drops, xp)
}

func (h *Handler) HandleQuit(p *player.Player) {
	h.bh.HandleQuit(p)
	h.ph.HandleQuit(p)
}
