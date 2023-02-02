// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 01 Feb 2023 03:06:24 AM CET
// Description: -
// ======================================================================
package datastructures

import (
   "fmt"

	art "waelder/internal/asciiart"
   "waelder/internal/renderer"
)

/** Struct representing an attack **/
type Attack struct {
	Round   int
	Source  string
	Targets []string
	HasHit  bool
	Damage  int
	Range   AttackType
}

/** Enum representing differen types of Attacks **/
type AttackType int64
const (
	Meele    AttackType = iota
	Ranged   AttackType = iota
	Magical  AttackType = iota
   Healing  AttackType = iota
)

func ToAttackType(i int) AttackType {
   switch i {
      case 0: return Meele
      case 1: return Ranged
      case 2: return Magical
      case 3: return Healing
   }

   return Meele
}

var attackIcon map[AttackType]string = map[AttackType]string{
	Meele:   art.OneLineSword,
	Ranged:  art.OneLineArrow,
	Magical: art.OneLineFire,
}

func (ma Attack) GetTargets() []string {return ma.Targets}
func (ma Attack) Display(d Data) renderer.Renderable {
   ch := d.GetCharacter(ma.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode(" "),
      renderer.GenerateNode(attackIcon[ma.Range], "bold"),
      renderer.GenerateNoRenderNode(" "),
   }
	for _, i := range ma.Targets {
      ch := d.GetCharacter(i)
		t = append(t, renderer.GenerateNode(ch.Name, ch.GetStyle()))
	}

   t = append(t, renderer.GenerateNoRenderNode(" "))
   if !ma.HasHit {
      t = append(t, renderer.GenerateNode("(missed)", "italic"))
   } else {
      t = append(
         t,
         renderer.GenerateNode(fmt.Sprintf("(%d dmg)",ma.Damage), "italic"),
      )
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (a Attack) Apply(ch Character) Character {
   if a.HasHit {
      ch.Stats.Hp -= a.Damage
   }

   return ch
}
func (a Attack) Undo(ch Character) Character{
   

   if a.HasHit {
      ch.Stats.Hp += a.Damage
   }

   return ch
}
func (a Attack) GetRoundNo() int { return a.Round }
