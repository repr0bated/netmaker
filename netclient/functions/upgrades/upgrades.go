package upgrades

func init() {
	addUpgrades([]UpgradeInfo{
		upgrade0147,
	})
}

// Upgrades - holds all upgrade funcs
var Upgrades = []UpgradeInfo{}

// addUpgrades - Adds upgrades to make to client
func addUpgrades(upgrades []UpgradeInfo) {
	Upgrades = append(Upgrades, upgrades...)
}

// ReleaseUpgrades - releases upgrade funcs from memory
func ReleaseUpgrades() {
	Upgrades = nil
}
