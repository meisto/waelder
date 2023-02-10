// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Thu 09 Feb 2023 11:50:15 PM CET
// Description: -
// ======================================================================
package root 

import (
   ds "waelder/internal/datastructures"
   "waelder/internal/layouts"
)


func turn(
   action ds.Action,
   data *ds.Data,
   layout layouts.Layout,
) {
   data.Step(action); 
   layout.Reset(*data)
}

