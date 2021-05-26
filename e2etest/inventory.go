package e2etest

import (
	env "istio.io/proxy/test/envoye2e/env"
)

var ExtensionE2ETests *env.TestInventory

func init() {
	ExtensionE2ETests = &env.TestInventory{
		Tests: []string{
			"Nop/TestDoNothing",
			"StaticHeader/AddHeader",
		},
	}
}
