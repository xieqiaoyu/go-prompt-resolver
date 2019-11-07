package main

//ParamCmd cmd accept params
type ParamCmd func(...string)

//Resolve implement Resolver
func (c ParamCmd) Resolve(t ...string) {
	c(t...)
}
