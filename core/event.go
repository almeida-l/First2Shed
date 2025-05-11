package core

type Event interface{}

// Commands events (external).
type StartGameCommand struct{}

// Internal events (issued by the game core itself).

// NotifyEvent marks an Event that should trigger an external notification.
//
// Events that implement NotifyEvent are used by the game core to signal
// that the external application (e.g., a Telegram bot or UI) should perform
// some notification or update in response to the event.
type NotifyEvent interface {
	// EmitNotification is a marker method with no logic.
	// Its presence identifies events that need external handling.
	EmitNotification()
}

type DealingFinishedEvent struct{}

type InitialCardSetEvent struct{}

func (e *InitialCardSetEvent) EmitNotification() {}

type CardResolvedEvent struct{}
