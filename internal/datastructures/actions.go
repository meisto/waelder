// ======================================================================
// Author: meisto
// Creation Date: Mon 16 Jan 2023 01:13:50 AM CET
// Description: -
// ======================================================================
package datastructures

import (
	art "waelder/internal/asciiart"
   "waelder/internal/renderer"
)

type Action interface {
   GetTargets() []string
	Display(Data) renderer.Renderable
	Apply(Character) Character
   Undo(Character) Character
	GetRoundNo() int
}

type AttackType int64

const (
	Meele    AttackType = iota
	Ranged   AttackType = iota
	Magical  AttackType = iota
   Healing  AttackType = iota
)

func GetAttackType(i int) AttackType {
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

type Attack struct {
	Round   int
	Source  string
	Targets []string
	HasHit  bool
	Damage  int
	Range   AttackType
}
func (ma Attack) GetTargets() []string {return ma.Targets}
func (ma Attack) Display(d Data) renderer.Renderable {
	t := []renderer.Renderable{
      d.GetCharacter(ma.Source).GetStyledName(),
      renderer.GenerateNoRenderNode(" "),
      renderer.GenerateNode(attackIcon[ma.Range], "bold"),
      renderer.GenerateNoRenderNode(" "),
   }
	for _, i := range ma.Targets {
		t = append(t, d.GetCharacter(i).GetStyledName())
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

type Heal struct {
	Round      int
	Source     string
	Targets    []string
	HasHit     bool
	HpRegained int
}

type Dash struct {
   Round int
   Source string
}
func (d Dash) GetTargets() []string { return []string{} }
func (d Dash) Display( data Data) renderer.Renderable {
   return renderer.GenerateNoRenderNode("NODATA")
}
func (d Dash) Apply(ch Character) Character {return ch}
func (d Dash) Undo(ch Character) Character {return ch}
func (d Dash) GetRoundNo() int { return d.Round }

type Disengage struct {
   Round int
   Source string
}
func (d Disengage) GetTargets() []string { return []string{} }
func (d Disengage) Display( data Data) renderer.Renderable {
   return renderer.GenerateNoRenderNode("NODATA")
}
func (d Disengage) Apply(ch Character) Character {return ch}
func (d Disengage) Undo(ch Character) Character {return ch}
func (d Disengage) GetRoundNo() int { return d.Round }

type Dodge struct {
   Round    int
   Source   string
}
func (d Dodge) GetTargets() []string { return []string{} }
func (d Dodge) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(d.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("prepared to dodge."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Dodge) Apply(ch Character) Character {return ch}
func (d Dodge) Undo(ch Character) Character {return ch}
func (d Dodge) GetRoundNo() int { return d.Round }

type Help struct {
   Round    int
   Source   string
}
func (d Help) GetTargets() []string { return []string{} }
func (h Help) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(h.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("hid themselve."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Help) Apply(ch Character) Character {return ch}
func (d Help) Undo(ch Character) Character {return ch}
func (d Help) GetRoundNo() int { return d.Round }

type Hide struct {
   Round    int
   Source   string
}
func (d Hide) GetTargets() []string { return []string{} }
func (h Hide) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(h.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("helped someone."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Hide) Apply(ch Character) Character {return ch}
func (d Hide) Undo(ch Character) Character {return ch}
func (d Hide) GetRoundNo() int { return d.Round }

type Ready struct {
   Round    int
   Source   string
}
func (r Ready) GetTargets() []string { return []string{} }
func (r Ready) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(r.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("prepared themselve."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Ready) Apply(ch Character) Character {return ch}
func (d Ready) Undo(ch Character) Character {return ch}
func (d Ready) GetRoundNo() int { return d.Round }

type Search struct {
   Round    int
   Source   string
}
func (d Search) GetTargets() []string { return []string{} }
func (s Search) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(s.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("used an object."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Search) Apply(ch Character) Character {return ch}
func (d Search) Undo(ch Character) Character {return ch}
func (d Search) GetRoundNo() int { return d.Round }

type UseObject struct {
   Round    int
   Source   string
}
func (d UseObject) GetTargets() []string { return []string{} }
func (uo UseObject) Display( data Data) renderer.Renderable {
	t := []renderer.Renderable{
      data.GetCharacter(uo.Source).GetStyledName(),
      renderer.GenerateNoRenderNode("used an object."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d UseObject) Apply(ch Character) Character {return ch}
func (d UseObject) Undo(ch Character) Character {return ch}
func (d UseObject) GetRoundNo() int { return d.Round }
