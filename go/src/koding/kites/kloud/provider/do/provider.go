package do

import "koding/kites/kloud/stack/provider"

var doProvider = &provider.Provider{
	Name:         "digitalocean",
	ResourceName: "droplet",
	Machine:      newMachine,
	Stack:        newStack,
	Schema:       newSchema(),
}

func init() {
	provider.Register(doProvider)
}
