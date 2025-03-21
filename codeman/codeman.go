package codeman

import (
	"fmt"

	"bank/bank"
)

// ExamplePlayer make sure to embed bank.PlayerControls into your player or you won't be able to implement the
// bank.PlayerStrategy interface
type Codeman struct {
	bank.PlayerControls
	RoundLimit int
}

func (c Codeman) Play(game bank.GameInfo, yourInfo bank.PlayerInfo) {
	if game.Round.Score >= c.RoundLimit {
		c.Bank()
	}
}

func (c Codeman) GetName() string {
	return fmt.Sprintf("Bank Once At %d", c.RoundLimit)
}
