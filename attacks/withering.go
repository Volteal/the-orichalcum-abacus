package attacks

import (
	"fmt"
	"strings"

	"github.com/volteal/the-orichalcum-abacus/functions"
)

func WitheringAttack() {
	// Dexterity + [relevant combat Ability] + weapon’s accuracy +any other modifiers (specialty, charm dice, stunt dice)
	var hasSpeciaty, wpa string
	var dexterity, combatSkill, weaponAccuracy, specialty, charmDice, stuntDice, successes, willpower int
	fmt.Println("Calculating Withering Attack")
	fmt.Println("What is your Dexterity?")
	fmt.Scan(&dexterity)

	fmt.Println("What what is the rating the combat skill [Archery, Brawl, Melee, Thrown, Martial Arts]?")
	fmt.Scan(&combatSkill)

	fmt.Println("What is your weapon's accuracy score?")
	fmt.Scan(&weaponAccuracy)

	fmt.Println("Do you have a relevant specialty? [Y]es/[N]o")
	fmt.Scan(&hasSpeciaty)

	if strings.ToLower(hasSpeciaty) == "y" || strings.ToLower(hasSpeciaty) == "yes" {
		specialty = 1
	} else {
		specialty = 0
	}

	fmt.Println("How many dice do you get from charms (including your Excellencies)?")
	fmt.Scan(&charmDice)

	stuntDice, successes = functions.Stunt()

	fmt.Println("Will you be spending Willpower?")
	fmt.Scan(&wpa)
	if strings.ToLower(wpa) == "y" || strings.ToLower(wpa) == "yes" {
		willpower = 1
	} else {
		willpower = 0
	}

	dicePool := dexterity + combatSkill + weaponAccuracy + specialty + charmDice + stuntDice

	fmt.Println("=====================================================================")
	fmt.Println("Your dice pool for this attack is:", dicePool)
	fmt.Println("Your automatic successes are:", successes+willpower)
}
