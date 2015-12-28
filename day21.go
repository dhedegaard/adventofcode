package main

import (
	"fmt"
	"time"
)

// A player or boss.
type entity struct {
	hitpoints    int
	damage       int
	armor        int
	boughtArmor  bool
	boughtRings  int
	boughtWeapon bool
}

func (e entity) String() string {
	return fmt.Sprint("[Entity: hitpoints:", e.hitpoints,
		" damage:", e.damage,
		" armor:", e.armor, "]")
}

func newEntity(hitpoints int, damage int, armor int) entity {
	return entity{
		hitpoints: hitpoints,
		damage:    damage,
		armor:     armor,
	}
}

// An item, with a given type.
type item struct {
	name     string
	cost     int
	damage   int
	armor    int
	itemtype string
}

func (i item) String() string {
	return fmt.Sprint(
		"[Item: name:", i.name, " type:", i.itemtype, " cost:", i.cost,
		" damage:", i.damage, " armor:", i.armor, "]")
}

const (
	weaponType = "weapon"
	armorType  = "armor"
	ringType   = "ring"
)

func newWeapon(name string, cost int, damage int) item {
	return item{
		name:     name,
		cost:     cost,
		damage:   damage,
		itemtype: weaponType,
	}
}

func newArmor(name string, cost int, armor int) item {
	return item{
		name:     name,
		cost:     cost,
		armor:    armor,
		itemtype: armorType,
	}
}
func newRing(name string, cost int, damage int, armor int) item {
	return item{
		name:     name,
		cost:     cost,
		damage:   damage,
		armor:    armor,
		itemtype: ringType,
	}
}

// Tries to equip the entity with an item, if the item cannot be equipped false
// is returned.
// Otherwise the item is equipped, modifying the stats of the entity.
func (e *entity) equip(i item) bool {
	// Validate that we can equip the item, otherwise return false.
	if (i.itemtype == weaponType && e.boughtWeapon) ||
		(i.itemtype == armorType && e.boughtArmor) ||
		(i.itemtype == ringType && e.boughtRings >= 2) {
		return false
	}

	// Modify the entity with the stats from the item.
	e.damage += i.damage
	e.armor += i.armor
	switch i.itemtype {
	case weaponType:
		e.boughtWeapon = true
		break
	case armorType:
		e.boughtArmor = true
		break
	case ringType:
		e.boughtRings++
		break
	default:
		panic(fmt.Sprint("Unknown itemtype:", i.itemtype))
	}
	return true
}

// A given entity hits a target. Ie from.hit(to)
func (e *entity) hit(target *entity) {
	damage := e.damage - target.armor
	if damage < 1 {
		damage = 1
	}
	target.hitpoints -= damage
}

// Checks to see if a given entity has enough hitpoints left.
func (e *entity) isDead() bool {
	return e.hitpoints <= 0
}

// Checks to see if a given player will win against a given boss opponent.
func willPlayerWin(player entity, boss entity) bool {
	for {
		// Player hits boss, if boss is not dead yet.
		player.hit(&boss)
		if boss.isDead() {
			return true
		}
		// Boss hits player, if player is not dead yet.
		boss.hit(&player)
		if player.isDead() {
			return false
		}
	}
}

func main() {
	// Assemble allItems for later.
	for _, weap := range weapons {
		allItems = append(allItems, &weap)
	}
	for _, armor := range armors {
		allItems = append(allItems, &armor)
	}
	for _, ring := range rings {
		allItems = append(allItems, &ring)
	}

	// Create initial state.
	player := newEntity(100, 0, 0)
	boss := input

	// Brute force all solutions, since it's fast enough :)
	before := time.Now()
	minCostAndWin := 999999
	maxCostAndLose := 0
	for _, weap := range weapons {
		currWeap := player
		currWeap.equip(weap)
		for _, armor := range armors {
			currArmor := currWeap
			currArmor.equip(armor)
			for _, ring1 := range rings {
				currRing1 := currArmor
				currRing1.equip(ring1)
				for _, ring2 := range rings {
					// You can only buy the same item once, rings are the only
					// items to override.
					if ring2 == ring1 {
						continue
					}
					currRing2 := currRing1
					currRing2.equip(ring2)
					playerWon := willPlayerWin(currRing2, boss)
					itemcost := weap.cost + armor.cost + ring1.cost + ring2.cost
					if playerWon && itemcost < minCostAndWin {
						minCostAndWin = itemcost
					} else if !playerWon && itemcost > maxCostAndLose {
						maxCostAndLose = itemcost
					}
				}
			}
		}
	}
	fmt.Println("minCostAndWin(part1):", minCostAndWin,
		"maxCostAndLose(part2):", maxCostAndLose,
		"took:", time.Now().Sub(before))
}

var input = newEntity(100, 8, 2)

// The items available to us.
var allItems = make([]*item, 0)
var weapons = []item{
	newWeapon("Dagger", 8, 4),
	newWeapon("Shortsword", 10, 5),
	newWeapon("Warhammer", 25, 6),
	newWeapon("Longsword", 40, 7),
	newWeapon("Greataxe", 74, 8),
}
var armors = []item{
	newArmor("nil armor", 0, 0),
	newArmor("Leather", 13, 1),
	newArmor("Chainmail", 31, 2),
	newArmor("Splintmail", 53, 3),
	newArmor("Bandedmail", 75, 4),
	newArmor("Platemail", 102, 5),
}
var rings = []item{
	newRing("nil ring", 0, 0, 0),
	newRing("Damage +1", 25, 1, 0),
	newRing("Damage +2", 50, 2, 0),
	newRing("Damage +3", 100, 3, 0),
	newRing("Defense +1", 20, 0, 1),
	newRing("Defense +2", 40, 0, 2),
	newRing("Defense +3", 80, 0, 3),
}
