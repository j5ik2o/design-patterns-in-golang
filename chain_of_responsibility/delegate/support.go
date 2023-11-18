package delegate

import "desgin-patterns-in-golang/chain_of_responsibility"

type Support interface {
	Support(trouble *chain_of_responsibility.Trouble)
	GetDelegate() *SupportDelegate
}
