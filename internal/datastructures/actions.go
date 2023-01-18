// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Mon 16 Jan 2023 01:13:50 AM CET
// Description: -
// ======================================================================
package datastructures


type MeeleAttack struct {
   Source   string
   Targets  []string
   HasHit   bool
   Damage   int
}

type RangedAttack struct {
   Source   string
   Targets  []string
   HasHit   bool
   Damage   int
}

type MagicAttack struct {
   Source   string
   Targets  []string
   HasHit   bool
   Damage   int
}

type Healing struct {
   Source      string
   Targets     []string
   HasHit      bool
   HpRegained  int
}
