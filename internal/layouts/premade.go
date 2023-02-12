// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 10 Feb 2023 11:12:05 PM CET
// Description: -
// ======================================================================
package layouts

import (
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
	"waelder/internal/modes"

   "waelder/internal/renderer"
)

func TwoOneHorizontalSplit(height, width int, output *termenv.Output, data ds.Data) Layout {

	l1 := int(float64(width) * 0.5)
	l2 := int(float64(height) * 0.8)

	// ─
   var content renderer.RenderField

   l := Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			{
            x: 0,
            y: 0,
            width: l1,
            height: l2,
            mode: modes.ActiveMode,
            content: content,
            padding: [4]int{0,0,0,0},
            scrollIndex: 100,
            borders: [4]bool{true,false,true,false},
            borderStyle: DoubleBorderStyle,
            startTop: false,
            output: output,
         },
			{
			   l1,
				0,
				width - l1,
				l2,
				modes.MdViewMode,
            content,
            [4]int{1,2,1,2},
            0,
				[4]bool{true, true, true, true},
				DoubleBorderStyle,
            false,
            output,
			},
			 {
				0,
            l2,
				width,
				height - l2 - 1,
				modes.ActionMode,
            content,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, true, false},
				DoubleBorderStyle,
            false,
            output,
			},
			{
				0,
            height - 1,
				width,
				1,
				modes.HelpMode,
            content,
            [4]int{0,0,0,0},
            -1,
				[4]bool{false, false, false, false},
				DoubleBorderStyle,
            false,
            output,
         },
		},
      output: output,
	}
   return l 
}


func Fullscreen(height, width int, output *termenv.Output, data ds.Data) Layout {

   var content renderer.RenderField
   return Layout{
		TotalHeight: height,
		TotalWidth:  width,
		Fields: []Field{
			{
            x: 0,
            y: 0,
            width: width,
            height: height,
            mode: modes.MdViewMode,
            content: content,
            padding: [4]int{0,1,0,1},
            scrollIndex: 1000,
            borders: [4]bool{true,true,true,true},
            borderStyle: DoubleBorderStyle,
            startTop: true,
            output: output,
         },
		},
      output: output,
	}
}
