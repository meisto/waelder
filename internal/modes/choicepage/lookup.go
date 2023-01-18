// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Tue 17 Jan 2023 11:59:49 PM CET
// Description: -
// ======================================================================
package choicepage

import (
   "dntui/internal/views"
   "dntui/internal/config"
   ds "dntui/internal/datastructures"
)

// Struct to hold complete specification of a choice page
type choiceSubPage struct {
   content     [][]string
   style       [][]views.ColorHelper
   selection   selection
}

type ChoicePageKey int64
const (
   Empty ChoicePageKey  = iota
   Test1  ChoicePageKey = iota
   Test2  ChoicePageKey = iota
)

//
var currentKey ChoicePageKey = Empty
func SetChoicePage(key ChoicePageKey) {currentKey = key}


// Generate a local lookup for subpages
// var subPageLookup map[ChoicePageKey]choiceSubPage = 
//   make(map[ChoicePageKey]choiceSubPage)
var subPageLookup [3]choiceSubPage = [...]choiceSubPage{
   {
      content:    [][]string {
         {"Hello there.", "-", "Obi-Wan"},
         {"BUTTON1", "Not a button"},
         {"BUTTON2", "Not a button"},
         {"BUTTON3", "Not a button"},
      },
      style:      [][]views.ColorHelper{
         {},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
      },
      selection:  []selectionElement{
         {display: "BUTTON4", style: config.Style1, i1: 1, i2: 0, action: func(d *ds.Data){}},
         {display: "", style: config.Style1, i1: 2, i2: 0, action: func(d *ds.Data){SetChoicePage(Test1)}},
      },
   },
   {
      content:    [][]string {
         {"Which opponent should This Player attack?"},
         {"A", "The Ork"},
         {"B", "The Goblin"},
         {"C", "A friend"},
      },
      style:      [][]views.ColorHelper{
         {},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
      },
      selection:  []selectionElement{
         {display: "", style: config.Style1, i1: 1, i2: 0, action: func(d *ds.Data){
            SetChoicePage(Test2)

            secretFunc = func(d *ds.Data) {
               print("The ork sends his regards")
            }

         }},
         {display: "", style: config.Style1, i1: 2, i2: 0, action: func(d *ds.Data){
            SetChoicePage(Test2)

            secretFunc = func(d *ds.Data) {
               print("The goblin sends his regards")
            }

         }},
         {display: "", style: config.Style1, i1: 3, i2: 0, action: func(d *ds.Data){
            SetChoicePage(Test2)

            secretFunc = func(d *ds.Data) {
               print("A friend sends his regards")
            }

         }},
      },
   },
   {
      content:    [][]string {
         {"What do you do now?"},
         {"A", "Run away (silently)."},
         {"B", "Scream."},
         {"C", "Do nothing."},
      },
      style:      [][]views.ColorHelper{
         {},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
         {{Index: 0, Style: config.StyleDarkRedBg}},
      },
      selection:  []selectionElement{
         {display: "", style: config.Style1, i1: 1, i2: 0, action: func(d *ds.Data){returnFunction()}},
         {display: "", style: config.Style1, i1: 2, i2: 0, action: func(d *ds.Data){secretFunc(d)}},
         {display: "", style: config.Style1, i1: 3, i2: 0, action: func(d *ds.Data){}},
      },
   },
}

var secretFunc func(d *ds.Data)

func setFunction(f func(d*ds.Data)) {
   subPageLookup[Test2].selection[2].action = f
}


