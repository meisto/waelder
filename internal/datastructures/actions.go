// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 01:13:50 AM CET
// Description: -
// ======================================================================
package datastructures

import (
	"fmt"

	art "waelder/internal/asciiart"
)

type Action interface {
   GetTargets() []string
	Display() string
	Apply(c Character) Character
	GetTurn() int
}

type AttackType int64

const (
	Meele   AttackType = iota
	Ranged  AttackType = iota
	Magical AttackType = iota
)

var icon map[AttackType]string = map[AttackType]string{
	Meele:   art.OneLineSword,
	Ranged:  art.OneLineArrow,
	Magical: art.OneLineFire,
}

type Attack struct {
	Turn    int
	Source  string
	Targets []string
	HasHit  bool
	Damage  int
	Range   AttackType
}
func (ma Attack) GetTargets() []string {return ma.Targets}
func (ma Attack) Display() string {
	t := ""
	for _, i := range ma.Targets {
		t += " "
		t += i
	}

	return fmt.Sprintf(
		"%10s %s %s",
		ma.Source,
		icon[ma.Range],
		t,
	)
}
func (a Attack) Apply(c Character) Character {
	if a.HasHit {
		c.Stats.Hp -= a.Damage
	}

	return c
}
func (a Attack) GetTurn() int { return a.Turn }

type Healing struct {
	Turn       int
	Source     string
	Targets    []string
	HasHit     bool
	HpRegained int
}
