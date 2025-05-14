# First2Shed

A UNO-inspired game engine implemented in Go. Designed for integration with chat platforms like Telegram for multiplayer gaming experiences.

## Project Structure

The project is being developed in two main phases:

1. Core Game Engine (Current Phase)
   - Complete card game implementation in Go
   - Finite state machine for game flow control
   - Event-driven architecture
   - CLI interface for testing

2. Telegram Bot Integration (Planned)
   - Interactive game sessions in Telegram group chats
   - Multi-player support through chat commands
   - User-friendly card selection interface
   - Multiple concurrent games across different chats
   - Players stats persistence

## Quick Start

### Testing with CLI Version

1. Clone the repository:
```bash
git clone https://github.com/almeida-l/First2Shed.git
cd First2Shed
```

2. Build the project:
```bash
go build -o cli main.go
```

3. Run the CLI version:
```bash
./cli
```

The CLI version automatically starts a 2-player game for testing. Available commands:
- Enter card code to play a card (e.g., "R1" for Red 1, "BT" for Blue Draw Two)
- Enter "1" to draw a card
- Enter "2" to pass after drawing
- Enter "3" to choose a color when playing a Wild card

Card Codes:
- Colors: R (Red), B (Blue), G (Green), Y (Yellow), W (Wild)
- Values: 0-9 (Numbers), S (Skip), R (Reverse), T (Draw Two), W (Wild), F (Wild Draw Four)

Example card codes: "R1" (Red 1), "BS" (Blue Skip), "WF" (Wild Draw Four)

No external dependencies are required - the project uses only Go standard library.

## Development Status

### Core Game Engine Features

| Feature | Status |
|---------|--------|
| Game flow with turn management | ✓ |
| Card matching rules | ✓ |
| 108-card deck | ✓ |
| Initial 7-card deal | ✓ |
| Draw and play/pass mechanics | ✓ |
| Two-player special rules | ✓ |
| Win condition | ✓ |

### Card Effects

| Effect | Status | Implementation |
|--------|--------|----------------|
| Skip | ✓ | Skips next player's turn |
| Reverse | ✓ | Changes play direction |
| Draw Two | ✓ | Next player draws 2 and loses turn |
| Wild | ✓ | Player chooses next color |
| Wild Draw Four | ✓ | Draw 4 + color choice + skip |

### Game Management

| Feature | Status | Notes |
|---------|--------|-------|
| Player joining/lobby system | ✓ | Complete |
| Hand management | ✓ | Add/remove/sort cards |
| Draw/discard pile management | ✓ | Auto-reshuffling |
| Valid play validation | ✓ | Complete |
| Turn direction management | ✓ | Complete |

### Planned Features

| Feature | Status | Phase |
|---------|--------|-------|
| Challenge rules for Wild Draw Four | Pending | Core Engine |
| Scoring system | Pending | Core Engine |
| Multiple round support | Pending | Core Engine |
| "Say UNO" mechanic | Pending | Core Engine |
| Penalties for not saying UNO | Pending | Core Engine |
| Time limits for player actions | Pending | Core Engine |
| Telegram Bot Integration | Planned | Phase 2 |
| Group Chat Support | Planned | Phase 2 |
| Concurrent Games Management | Planned | Phase 2 |
| Player Statistics System | Planned | Phase 2 |
| Interactive Card Selection | Planned | Phase 2 |

### Unsupported Features

The following features are intentionally not implemented:
- Allow Draw Two Stacking rule
- "Draw until you can play" rule

## Architecture

### Finite State Machine (FSM)

The engine processes events through a well-defined FSM with the following states:

#### StateLobby
- Waiting for players to join
- Handles global player join events
- Transitions to `StateDealing` when game starts (requires minimum 2 players)

#### StateDealing
- Initializes the draw pile with a complete 108-card deck
- Shuffles the deck
- Deals 7 cards to each player
- Transitions to `StateSettingInitialCard`

#### StateSettingInitialCard
- Flips the first card from the draw pile
- If the card is wild, it is returned to the deck and reshuffled until a non-wild card is drawn
- Transitions to `StateResolvingCard`

#### StateResolvingCard
- Centralized logic to apply card effects:
  - Skip: Advances turn to skip next player
  - Reverse: Changes turn direction and advances turn
  - Draw Two: Makes next player draw 2 cards and skips their turn
  - Wild: Waits for player to choose color (transitions to `StateAwaitingColorChoice`)
  - Wild Draw Four: Makes next player draw 4 cards, skips their turn, waits for color choice
- Also handles normal cards for turn advancement
- Transitions to:
  - `StateGameOver` if the current player has emptied their hand
  - `StatePlayerTurn` otherwise

#### StateAwaitingColorChoice
- Waits for the current player to choose a color after playing a wild card
- Only accepts valid colors (Red, Blue, Green, Yellow)
- Transitions back to `StateResolvingCard` when color is set

#### StatePlayerTurn
- Validates and handles player actions:
  - Play a card (must match color or value, or be wild)
  - Draw a card (only once per turn)
  - Pass (only allowed after drawing)
- Transitions to:
  - `StateResolvingCard` after a card is played
  - Stays in `StatePlayerTurn` after drawing
  - Changes to next player's turn when passing

#### StateGameOver
- Terminal state when a player wins by emptying their hand
- No further transitions possible
