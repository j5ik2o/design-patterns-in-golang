package chain_of_responsibility

type Support interface {
	SupportBase
	Resolve(trouble *Trouble) bool
	Support(trouble *Trouble)
}
