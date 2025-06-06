package core

type Event interface{}

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

// GlobalEvent marks an Event that can be handled in any game state.
//
// Events that implement GlobalEvent are intercepted and processed by
// the game engine itself rather than being delegated to the current FSM state.
type GlobalEvent interface {
	MarkAsGlobal()
}

// Commands events (external).

type PlayerJoinCommand struct {
	ID int
}

func (PlayerJoinCommand) MarkAsGlobal() {}

type StartGameCommand struct{}

type PlayCardCommand struct {
	Card   Card
	Player *Player
}

type DrawCardCommand struct {
	Player *Player
}

func (DrawCardCommand) EmitNotification() {}

type PassCommand struct {
	Player *Player
}

func (PassCommand) EmitNotification() {}

type SetWildColorCommand struct {
	Player *Player
	Color  Color
}

func (SetWildColorCommand) EmitNotification() {}

// Internal events (issued by the game core itself).

type DealingFinishedEvent struct{}

type InitialCardSetEvent struct{}

func (InitialCardSetEvent) EmitNotification() {}

type CardResolvedEvent struct{}

type WildCardPlayedEvent struct{}

func (WildCardPlayedEvent) EmitNotification() {}

type SetWinner struct {
	Player *Player
}

func (SetWinner) EmitNotification() {}
