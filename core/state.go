package core

type State interface {
	OnEnter(*Game)
	CanHandle(*Game, Event) bool
	Next(*Game, Event) State
}
