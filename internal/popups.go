// ======================================================================
// Author: meisto
// Creation Date: Wed 25 Jan 2023 05:40:23 PM CET
// Description: -
// ======================================================================
package root

import (
   "fmt"
   "log"
   "strconv"

   "waelder/internal/asciiart"
	ds "waelder/internal/datastructures"
	"waelder/internal/layouts"
   "waelder/internal/modes"
   "waelder/internal/renderer"
   "waelder/internal/wio"
)

func getDodge(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Dodge{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}
func getHelp(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Help{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}
func getHide(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Hide{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}
func getReady(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Ready{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}
func getSearch(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Search{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}
func getUseObject(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.UseObject{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}

func getDisengage(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Disengage{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}

func getDash(
   data *ds.Data,
   layout layouts.Layout,
) {
   ac := ds.Dash{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
   }
   turn(ac, data, layout)
}


func getAttack(
   data     *ds.Data,
   layout   layouts.Layout,
   eventChannel chan string,
) {

   width := 40
   x := layout.TotalWidth / 2 - width / 2
   y := 20


   // Choice module to pick target
   var c1 modes.Choice
   labels  := append(data.Enemies, data.Neutrals...)
   labels   = append(labels, data.Allies...)
   labels   = append(labels, data.Players...)
   { // Generate first popup menu
      l := len(data.Players) + len(data.Allies) + len(data.Enemies) +
         len(data.Neutrals)
      buttons := []string{
         "1", "2", "3", "4", "5", "6","7","8","9","0","q","w","e","r","t",
         "z","u","i","l","y",
      }

      for i := 0; i < len(buttons); i++ {
         buttons[i] = fmt.Sprintf(" %s ", buttons[i])
      }
      styledLabels := make([]renderer.Renderable, len(labels))

      f := func(x renderer.Renderable) renderer.Renderable {
         return renderer.GenerateLine(
            x.Length() + 1,
            []renderer.Renderable{renderer.GenerateNoRenderNode(" "), x},
         )
      }

      i := 0
      for j := 0; j < len(data.Enemies); j++ {
         ch := data.GetCharacter(data.Enemies[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Neutrals); j++ {
         ch := data.GetCharacter(data.Neutrals[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Allies); j ++ {
         ch := data.GetCharacter(data.Allies[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }
      for j := 0; j < len(data.Players); j++ {
         ch := data.GetCharacter(data.Players[j])
         styledLabels[i] = f(renderer.GenerateNode(ch.Name, ch.GetStyle()))
         i += 1
      }

      c1 = modes.GetChoice(
         fmt.Sprintf(
            " Whom should %s attack?",
            data.CombatLog.Current.ActiveCharacter,
         ),
         buttons[:l],
         styledLabels,
      )
   }


   // Choice module to pick target
   var c2 modes.Choice
   {

      f := func(a, b string) renderer.RenderLine {
         return renderer.GenerateLine(
            width -5,
            []renderer.Renderable{
               renderer.GenerateNode(a, "italic"),
               renderer.GenerateNode(b, "bold"),
            },
         )
      }

      labels := []renderer.Renderable{
         f(" Meele    ", asciiart.OneLineSword),
         f(" Ranged   ", asciiart.OneLineArrow),
         f(" Magical  ", asciiart.OneLineFire),
         f(" Healing  ", asciiart.OneLineFire),
      }

      c2 = modes.GetChoice(
         "Which mode of attack?",
         []string{" 1 "," 2 "," 3 "," 4 "},
         labels,
      )
   }

   var c3 modes.Choice
   {
      b := []string{" Yes ", " No "}
      labels := []renderer.Renderable {
         renderer.GenerateNoRenderNode(""),
         renderer.GenerateNoRenderNode(""),
      }

      c3 = modes.GetChoice("Did they hit?", b, labels)
   }
 
   redrawChoicePopup := func(c modes.Choice) {
      con := c.ToRenderField(x,y,width)
      pop := layouts.NewPopupField(x,y, con)
      log.Print(con.GetWidth(), " ", con.GetHeight())
      layout.DisplayPopup(pop)
   }


   { // Poll for target
      for {
         redrawChoicePopup(c1)
         s, ok := <- eventChannel
         if !ok { return }
         if s == "<ENTER>" { break }

         switch s {
            case "j": c1.Forward()
            case "k": c1.Backward()
            case "0": c1.GoTo(9)
            case "1": c1.GoTo(0)
            case "2": c1.GoTo(1)
            case "3": c1.GoTo(2)
            case "4": c1.GoTo(3)
            case "5": c1.GoTo(4)
            case "6": c1.GoTo(5)
            case "7": c1.GoTo(6)
            case "8": c1.GoTo(7)
            case "9": c1.GoTo(8)
         }
      }
   }
   layout.Reset(*data)


   { // Poll for attack mode
      for {
         redrawChoicePopup(c2)
         s, ok := <- eventChannel
         if !ok { return }

         if s == "<ENTER>" { break }

         switch s {
            case "j": c2.Forward()
            case "k": c2.Backward()
            case "1": c2.GoTo(0);
            case "2": c2.GoTo(1);
            case "3": c2.GoTo(2);
            case "4": c2.GoTo(3);
         }
      }
   }
   layout.Reset(*data)


   { // Poll for hit
      for {
         redrawChoicePopup(c3)
         s, ok := <- eventChannel
         if !ok { return }
         if s == "<ENTER>" { break }

         switch s {
            case "j": c3.Forward();
            case "k": c3.Backward();
            case "1": c3.GoTo(0);
            case "2": c3.GoTo(1);
         }
      }
   }
   layout.Reset(*data)


   // Prompt for damage if not hit
   var damage int = 0
   if c3.GetIndex() == 0 {
      content := renderer.GenerateField(
         []renderer.RenderLine{
            renderer.GenerateLineFromOne(
               width - 2,
               renderer.GenerateNoRenderNode("How much damage did they deal?"),
            ),
         },
      )
      layout.DisplayPopup(layouts.NewPopupField(x,y, content))

      // Try to cast string to int
      n, err := strconv.Atoi(<- wio.ReadLine(true))
      for err != nil || n < 0 || n > 1000 {
         content = renderer.GenerateField(
            []renderer.RenderLine{
               renderer.GenerateLineFromOne(
                  width - 2,
                  renderer.GenerateNoRenderNode("Illegal input, ty again."),
               ),
            },
         )
         layout.DisplayPopup(layouts.NewPopupField(x, y, content))
         n, err = strconv.Atoi(<- wio.ReadLine(true))
      }
      damage = n
   } 

   at :=ds.Attack{
      Round: data.CombatLog.Current.RoundNumber,
      Source: data.CombatLog.Current.ActiveCharacter,
      Targets: []string{labels[c1.GetIndex()]},
      HasHit: c3.GetIndex() == 0,
      Damage: damage,
      Range: ds.ToAttackType(c3.GetIndex()),
   }
   turn(at, data, layout)   
}
