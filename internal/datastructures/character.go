// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 21 Nov 2022 03:13:28 PM CET
// Description: -
// ======================================================================
package datastructures

import (
   // tea "github.com/charmbracelet/bubbletea"
)


type Character struct{
   Name        string
   Affiliation string
   Race        string
   Subrace     string
   Class       string

   Stats       CharacterStats
}

type CharacterStats struct {
   Hp       int
   Max_hp   int
   Initiative  int
}
