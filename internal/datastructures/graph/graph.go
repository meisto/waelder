// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Fri 20 Jan 2023 12:55:59 AM CET
// Description: -
// ======================================================================
package graph

import ds "waelder/internal/datastructures"

type Node int32
type Edge struct {
   Root     Node
   Label    string
   Target   Node
   action   func()
}


var nodeIndex int32 = 0
func GetNode() Node {
   nodeIndex += 1

   return Node(nodeIndex - 1)
}

func GetEdge(root Node, label string, target Node, action func()) Edge {
   return Edge{root, label, target, action}
}


type Graph struct {
   activeNode  Node
   nodes       []Node
   edges       []Edge

   lookup map[Node][]Edge
}
func GetGraph(rootNode Node) Graph { 

   g := Graph{
      activeNode: rootNode,
      lookup: make(map[Node][]Edge),
   }

   g.AddNode(rootNode)

   return g
}
func (g *Graph) AddNode(n Node) bool {

   for _, i := range(g.nodes) { if i == n { return false } }

   g.nodes = append(g.nodes, n)

   return true
   
}
func (g *Graph) AddEdge(e Edge) bool {
   hasRoot     := false
   hasTarget   := false

   for _, i := range(g.nodes) {
      if i== e.Root { hasRoot = true }
      if i== e.Target { hasTarget = true }
   }

   if hasRoot {
      g.edges = append(g.edges, e)

      g.lookup[e.Root] = append(g.lookup[e.Root], e)

      if !hasTarget { g.AddNode(e.Target) }


      return true
   }
   
   return false
}

func (g *Graph) Step(key string, d *ds.Data) bool {


   for _, i := range(g.lookup[g.activeNode]) {
      if i.Label == key {
         g.activeNode = i.Target

         i.action()

         return true
      }
   }
   return false

}




