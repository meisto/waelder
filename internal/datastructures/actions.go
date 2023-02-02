// ======================================================================
// Author: meisto
// Creation Date: Mon 16 Jan 2023 01:13:50 AM CET
// Description: -
// ======================================================================
package datastructures

import (
   "waelder/internal/renderer"
)

type Action interface {
   GetTargets() []string
	Display(Data) renderer.Renderable
	Apply(Character) Character
   Undo(Character) Character
	GetRoundNo() int
}

type Heal struct {
	Round      int
	Source     string
	Targets    []string
	HasHit     bool
	HpRegained int
   Medium     string
}


/** Reaction Action **/
type Reaction struct {
   Round    int
   Source   string
   Action   Action
}
func (re Reaction) GetTargets() []string {return re.Action.GetTargets()}
func (re Reaction) Display(data Data) renderer.Renderable {
   a := renderer.GenerateNoRenderNode(re.Source + " reacted:")
   b := re.Action.Display(data)

   return renderer.GenerateLine(
      a.Length() + b.Length(),
      []renderer.Renderable{a,b},
   )
}
func (re Reaction) Apply(ch Character) Character { return re.Action.Apply(ch)}
func (re Reaction) Undo(ch Character) Character { return re.Action.Undo(ch)}
func (re Reaction) GetRoundNo() int { return re.Round }


/** Dash Action **/
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


/** Disengage Action **/
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


/** Dodge Action **/
type Dodge struct {
   Round    int
   Source   string
}
func (d Dodge) GetTargets() []string { return []string{} }
func (d Dodge) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(d.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("prepared to dodge."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Dodge) Apply(ch Character) Character {return ch}
func (d Dodge) Undo(ch Character) Character {return ch}
func (d Dodge) GetRoundNo() int { return d.Round }


/** Help Action **/
type Help struct {
   Round    int
   Source   string
}
func (d Help) GetTargets() []string { return []string{} }
func (h Help) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(h.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("hid themself."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Help) Apply(ch Character) Character {return ch}
func (d Help) Undo(ch Character) Character {return ch}
func (d Help) GetRoundNo() int { return d.Round }


/** Hide Action **/
type Hide struct {
   Round    int
   Source   string
}
func (d Hide) GetTargets() []string { return []string{} }
func (h Hide) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(h.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("helped someone."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Hide) Apply(ch Character) Character {return ch}
func (d Hide) Undo(ch Character) Character {return ch}
func (d Hide) GetRoundNo() int { return d.Round }


/** Ready Action **/
type Ready struct {
   Round    int
   Source   string
}
func (r Ready) GetTargets() []string { return []string{} }
func (r Ready) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(r.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("prepared themself."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Ready) Apply(ch Character) Character {return ch}
func (d Ready) Undo(ch Character) Character {return ch}
func (d Ready) GetRoundNo() int { return d.Round }


/** Search action **/
type Search struct {
   Round    int
   Source   string
}
func (d Search) GetTargets() []string { return []string{} }
func (s Search) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(s.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("used an object."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d Search) Apply(ch Character) Character {return ch}
func (d Search) Undo(ch Character) Character {return ch}
func (d Search) GetRoundNo() int { return d.Round }


/** UseObject **/
type UseObject struct {
   Round    int
   Source   string
}
func (d UseObject) GetTargets() []string { return []string{} }
func (uo UseObject) Display( data Data) renderer.Renderable {
   ch := data.GetCharacter(uo.Source)
	t := []renderer.Renderable{
      renderer.GenerateNode(ch.Name, ch.GetStyle()),
      renderer.GenerateNoRenderNode("used an object."),
   }

   l := 0
   for _, i := range(t) {l += i.Length()}

   return renderer.GenerateLine(l, t)
}
func (d UseObject) Apply(ch Character) Character {return ch}
func (d UseObject) Undo(ch Character) Character {return ch}
func (d UseObject) GetRoundNo() int { return d.Round }
