package brush

import (
	"github.com/df-mc/dragonfly/server/player"
	"sync"
)

// handlers is a map holding handlers of players.
var handlers sync.Map

// LookupHandler finds the Handler of a specific player.Player, assuming it is currently online.
func LookupHandler(p *player.Player) (*Handler, bool) {
	v, _ := handlers.Load(p.Name())
	h, ok := v.(*Handler)
	return h, ok
}

// Handler implements the brushing of players. It enables activation of brushes and stores the data needed to
// undo/redo those actions.
type Handler struct {
	undo []func()
}

// NewHandler creates a new Handler for the *player.Player passed.
func NewHandler(p *player.Player) *Handler {
	h := &Handler{}
	handlers.Store(p.Name(), h)
	return h
}

// UndoLatest undoes the latest brush action. If no action was left to undo, false is returned.
func (h *Handler) UndoLatest() bool {
	if len(h.undo) == 0 {
		return false
	}
	offset := len(h.undo) - 1
	h.undo[offset]()
	h.undo = h.undo[:offset]
	return true
}

// HandleItemUse activates the brush on a player's item if present.
func (h *Handler) HandleItemUse(ctx *player.Context) {
	held, _ := ctx.Val().HeldItems()
	if b, ok := find(held); ok {
		ctx.Cancel()
		go b.Use(ctx.Val().H())
	}
}

// HandleQuit deletes the Handler from the handlers map.
func (h *Handler) HandleQuit(p *player.Player) {
	handlers.Delete(p.Name())
}
