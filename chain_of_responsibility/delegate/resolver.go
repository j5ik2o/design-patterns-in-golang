package delegate

import "desgin-patterns-in-golang/chain_of_responsibility"

type Resolver func(trouble *chain_of_responsibility.Trouble) bool
