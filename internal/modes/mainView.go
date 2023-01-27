// ======================================================================
// Author: meisto
// Creation Date: Mon 21 Nov 2022 03:09:49 PM CET
// Description: -
// ======================================================================
package modes

/*
import(
   "fmt"
   "strings"

   tea "github.com/charmbracelet/bubbletea"
   "github.com/charmbracelet/bubbles/progress"
   "github.com/charmbracelet/lipgloss"

   "waelder/internal/config"
   "waelder/internal/views"
   ds "waelder/internal/datastructures"
)

// View method
func MainView(
   d ds.Data,
   windowHeight int,
   windowWidth int,
) string {
   var s string

   // Bar setup
   bar := progress.New(progress.WithSolidFill("#000000"))
   bar.ShowPercentage = false
   bar.Width = 10

   // Iterate over our choices
   for i1, x := range [4][]ds.Character{
      d.Players,
      d.Allies,
      d.Enemies,
      d.Neutrals,
   } {

      heading := "DEFAULT HEADING"
      switch i1 {
         case 0: heading = "Players"
         case 1: heading = "Allies"
         case 2: heading = "Enemies"
         case 3: heading = "Neutrals"
      }
      s += config.Style1.Render(fmt.Sprintf(" %s ", heading))
      s += "\n"

      for i2 := 0; i2 < len(x); i2++ {

         char := x[i2]


         var hpPercentage float64 = float64(char.Stats.Hp) / float64(char.Stats.Max_hp)

         isActive := false // i1 == cursor[1] && i2 == cursor[0]
         isDead := hpPercentage == 0.0

         f := func(s string) string {return views.FormatString(s, isActive, isDead)}
         comp1 := f(fmt.Sprintf(" %10s ",char.Name))
         health := views.FormatHealthString(
            10,
            hpPercentage,
            isActive,
         )
         healthNumeral := f(fmt.Sprintf("%03d/%03d", char.Stats.Hp, char.Stats.Max_hp))
         separator := f(" ")
         s+= strings.Join([]string{
            comp1,
            health,
            healthNumeral,
            f("🛡 1011"),
            f("👞20"),
         },
         separator) + "\n"

         /*
         // Grey out dead characters
         if isActive {
            line = formatStatusLine(components, config.StyleSelected)

         } else if hpPercentage == 0 {
            line = formatStatusLine(components, config.StyleDead)
         } else {
            line = strings.Join(components, " ")
         }

      }

      s += "\n"

   }

   return s
}



func formatStatusLine(content []string, style lipgloss.Style) string {

   for i := 0; i < len(content); i++ {
      content[i] = style.Render(content[i])
   }

   return strings.Join(content, style.Render(" "))
}

func MainUpdate(msg tea.KeyMsg) {
}
*/
