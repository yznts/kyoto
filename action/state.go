package action

import (
	"errors"

	"go.kyoto.codes/v3/component"
)

// State allows to extract component state from action.
// It's an important part of the action workflow that
// might be easily missed and have to be mentioned explicitly in the docs.
func State(ctx *component.Context, state component.State) error {
	// Prepare context
	context(ctx)
	// Pass, if action not recognized.
	// It means we are not in the action workflow now.
	if ctx.Store.Get("Action") == nil {
		return errors.New("action not recognized")
	}
	// Extract action
	action := ctx.Store.Get("Action").(Action)
	// Pass, if component not recognized
	if action.Component == "" {
		return errors.New("component not recognized")
	}
	// Unmarshal state
	state.Unmarshal(action.State)
	// Return
	return nil
}
