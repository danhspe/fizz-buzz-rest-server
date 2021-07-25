package argumentStore

import "github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"

type ArgumentStore = map[arguments.Arguments]int

func New() ArgumentStore {
	return make(ArgumentStore)
}
