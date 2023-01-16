// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 05:57:30 PM CET
// Description: -
// ======================================================================
package data

import (
   cm "dntui/internal/model/charactermodel"
)


type Data struct {
   SQLitePath        string

   CharacterStore    map[string]cm.Character
   Players           []string
   Allies            []string
   Enemies           []string
   Neutrals          []string
}



type Data2 struct {
   Allies     []cm.Character
   Pcs      []cm.Character
   Enemies     []cm.Character
   Neutral   []cm.Character
}
