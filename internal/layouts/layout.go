// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Wed 18 Jan 2023 03:22:02 AM CET
// Description: -
// ======================================================================
package layouts

import (
   "log"
   "math"
   "strings"

   // tea "github.com/charmbracelet/bubbletea"

   "waelder/internal/modes"
   ds "waelder/internal/datastructures"
)


type Layout struct {
   TotalHeight int
   TotalWidth  int
   rows        []Row
}

type Row struct {
   heightPer   float64
   height      int
   fields      []Field
}

type Field struct {
   widthPer    float64
   width       int
   mode        modes.Mode
}

var Fullscreen = Layout{
   TotalHeight:   100,
   TotalWidth:    100,
   rows: []Row {
      {
         heightPer: 1.0, 
         height: 100, 
         fields: []Field {
            {
               widthPer: 1.0,
               width: 100,
               mode: modes.ActiveMode,
            },
         },
      },
   },
}

var TwoThirdsHorizontalSplit = Layout {
   TotalHeight:   0,
   TotalWidth:    0,
   rows: []Row {
      {
         heightPer: 0.66, 
         height: 0, 
         fields: []Field {
            {
               widthPer: 1.0,
               width: 1.0,
               mode: modes.ActiveMode,
            },
         },
      },
      {
         heightPer: 0.33, 
         height: 0, 
         fields: []Field {
            {
               widthPer: 1.0,
               width: 1.0,
               mode: modes.ActiveMode,
            },
         },
      },
   },
}



func (lay Layout) Resize(height int, width int) Layout {
   newRows := make([]Row, len(lay.rows))

   remainingHeight := height - (len(lay.rows) - 1)
   for i := 0; i < len(lay.rows); i++ {
      row := lay.rows[i]

      var rowHeight int
      if i == len(lay.rows) - 1 {
         // Last row
         rowHeight = remainingHeight
      } else {
         rowHeight = int(math.Floor(float64(height - (len(lay.rows) - 1)) * row.heightPer))
      }

      row.height        = rowHeight
      remainingHeight   -= rowHeight

      newFields := make([]Field, len(row.fields))
      remainingWidth := width - (len(row.fields) - 1)
      for j := 0; j < len(row.fields); j++ {
         field := row.fields[j]

         var fieldWidth int
         if j == len(row.fields) - 1 {
            // Last field
            fieldWidth = remainingWidth
         } else {
            fieldWidth = int(math.Floor(float64(width - (len(row.fields) - 1)) * field.widthPer))
         }

         remainingWidth -= fieldWidth
         newFields[j] = Field {
            widthPer:   field.widthPer,
            width:      fieldWidth,
            mode:       field.mode,
         }
      }

      newRows[i] = Row{ heightPer: row.heightPer, height: rowHeight, fields: newFields}
   }

   {
      b := 0
      for _, r := range(newRows) {b += r.height}
      b += len(newRows) - 1
      if b != height {
         log.Fatalf("|| %d %d",b, height)
         log.Fatal("[ERROR] Code: 08052265")}
   }


   return Layout {
      TotalHeight:   height,
      TotalWidth:    width,
      rows:          newRows,
   }
}

func (lay *Layout) Display(data ds.Data) string {
   res := make([]string, lay.TotalHeight)
   offset := 0

   for i := 0; i < len(lay.rows); i++ {
      row := lay.rows[i]


      var rowContent [][]string

      for j := 0; j < len(row.fields); j++ {
         field := row.fields[j]

            var x []string = modes.ActiveView(
               data, 
               row.height,
               field.width,
            )

         rowContent = append(
            rowContent, 
            x,
            // modes.ModeLookup[field.mode].View(
         )
      }

      //TODO: Remove the sanity checks
      {
         a := len(rowContent[0])
         for _, r := range(rowContent) {
            if len(r) != a {log.Fatal("[ERROR] Code: 44160149")}
         }

         b := 0
         for _, r := range(lay.rows) { b += r.height }
         b += len(lay.rows) - 1
         if b != lay.TotalHeight {log.Fatal("[ERROR] Code: 13519900")}
      }


      for j := 0; j < row.height; j++ {

         // Connect rows into single strings
         s := ""
         for _, c := range(rowContent) { s += "|" + c[j]}
         s = s[1:]

         res[offset + j] = s
      }


      offset += row.height + 1
      if i < len(lay.rows) - 1 {
         res[offset - 1] = strings.Repeat("~", lay.TotalWidth)
      }

   }

   return strings.Join(res, "\n")
}

// func (lay *Layout) Delegate(data *ds.Data, msg tea.Msg) {


/*

func generateLayout(percentages [][]float64) Layout {
   var vertSum float64 = 0.0


   for _, row := range(percentages) {

      rowHeight := row[0]
      var rowSum float64 = 0.0

      for _, cell := range(row) {
         if cell != rowHeight {log.Fatal("[ERROR] Code: 34181701")}

      }
      vertSum += rowHeight

   }

   if vertSum != 1.0 {log.Fatal("[ERROR] Code: 2315880")}


}
*/
