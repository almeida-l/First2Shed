# First2Shed Game

This project implements a game engine inspired by UNO in Go with a well-defined finite state machine (FSM) for managing game flow.
The goal of this project is to be integrated with Telegram bots so it can be played in group chats.

## Status
This project is currently in the **initial stages of development** and is **not yet runnable**.  

---

## Finite State Machine (FSM)

The engine is driven by events injected through the game `Process` method, which drives an FSM with the following states:

### `StateLobby`
- Waiting for players to join.
- Transitions to `StateDealing` when the game starts.

### `StateDealing`
- Initializes the draw pile.
- Deals 7 cards to each player.
- Transitions to `StateSettingInitialCard`.

### `StateSettingInitialCard`
- Flips the first card from the draw pile.
- If the card is invalid (e.g., Wild Draw Four), it is replaced.
- Transitions to `StateResolvingCard`.

### `StateResolvingCard`
- Centralized logic to apply card effects:
  - Skip
  - Reverse
  - Draw Two
  - Wild (waits for player to choose color)
  - Wild Draw Four
- Also called for normal cards to centralize resolution and turn advancement.
- Transitions to:
  - `StateGameOver` if the last player has emptied their hand.
  - `StatePlayerTurn` otherwise.

### `StatePlayerTurn`
- Waits for current player to:
  - Play a card (must match color or value, or be wild)
  - Or draw and play it or pass.
- Transitions to `StateResolvingCard` after a play.
- Transitions back to `SatePlayerTurn` when player pass (after changing turn)

### `StateGameOver`
- Ends the game
