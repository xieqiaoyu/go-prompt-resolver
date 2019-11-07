package main

//Cmd a cmd action
type Cmd func()

//Resolve implement Resolver
func (c Cmd) Resolve(...string) {
	c()
}
