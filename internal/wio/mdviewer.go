// ======================================================================
// Author: meisto
// Creation Date: Thu 26 Jan 2023 02:30:08 PM CET
// Description: -
// ======================================================================
package wio

import (
   "fmt"
   "log"
   "regexp"
   "strings"
   "strconv"
   "unicode/utf8"

   "waelder/internal/renderer"
)

type MDDocument struct {
   base     string
   renderNodes []parseNode
   isMultiselect bool
   index int

   triggerMap map[string]func()
}

func (d *MDDocument) SelectNext() {
   // No focusable elements present
   if d.index == -1 {return }

   // Unfocus current element
   a, ok := d.renderNodes[d.index].(focusableParseNode)
   if ok  {
      d.renderNodes[d.index] = a.Unfocus()
   }

   for i := 1; i < len(d.renderNodes) + 1; i++ {
      a, ok := d.renderNodes[(d.index + i) % len(d.renderNodes)].(focusableParseNode)

      // Check if this element can be focused, if yes focus it and return
      if ok {
         d.index = (d.index + i) % len(d.renderNodes)
         d.renderNodes[d.index] = a.Focus()

         return 
      }
   }
}

func (d *MDDocument) ActivateElement() {
   switch a := d.renderNodes[d.index].(type) {
      case toggleParseNode:
         a, ok := d.renderNodes[d.index].(toggleParseNode)
         if ok {
            if !d.isMultiselect {
               for i := 0; i < len(d.renderNodes); i++ {
                  b, ok := d.renderNodes[i].(toggleParseNode)
                  if ok {
                     d.renderNodes[i] = b.Deactivate()
                  }
               }
            }

            d.renderNodes[d.index] = a.Toggle()
         }
      case readInputParseNode:
         d.renderNodes[d.index] = a.readInput()
         d.SelectNext()

      case actionParseNode:
         id := a.GetId()
         action, ok := d.triggerMap[id]
         if !ok {
            log.Fatal("Could not retrieve action of id '", id, "'.")
         }
         action()

   }

}

func (d *MDDocument) ActivateMultiselect() { d.isMultiselect = true }
func (d *MDDocument) DeactivateMultiselect() { d.isMultiselect = false }
func (d MDDocument) GetActiveElements() []string {
   res := []string{}
   for i := 0; i < len(d.renderNodes); i++ {
      a, ok := d.renderNodes[i].(toggleParseNode)
      if ok && a.IsActive() {
         res = append(res, a.GetId())

         // Found all elements in singleselect mode 
         if !d.isMultiselect { break }
      }
   }
   return res
}



func (d MDDocument) Render(width int) renderer.RenderField {
   res := []renderer.RenderLine{}  
   
   // Keep line start
   l := []renderer.Renderable{}
   used := 0
   isBlockquote := false


   flush := func() {
      res = append(res, renderer.GenerateLine(width, l))
      l = []renderer.Renderable{}
      used = 0

      if isBlockquote {
         l = append(l, renderer.GenerateNode(" ", "MarkdownBlockquote"))
         l = append(l, renderer.GenerateNoRenderNode(" "))
         used += 2
      }
   }

   for i := 0; i < len(d.renderNodes); i++ {
      a := d.renderNodes[i]

      switch x := a.(type) {
         case LineBreakparseNode:
            isBlockquote = false
            flush()

         case fullLineparseNode:
            isBlockquote = false

            // Flush remaining documents
            if len(l) > 0 {
               flush()
            }
            res = append(res, renderer.GenerateLineFromOne(width, renderer.GenerateNode(x.content, x.style)))

         case AtomarparseNode:
            var a renderer.Renderable
            a = renderer.GenerateNode(x.content, x.style)

            if a.Length() > width {
               a = renderer.GenerateNoRenderNode("[PARSEERROR]")
            }

            if used + a.Length() >= width {
               flush()
            } 
            l = append(l, a)
            used += a.Length()

         case toggleParseNode:

            var button, label renderer.Renderable

            if x.IsFocused() {
               label = renderer.GenerateNode(x.label, "MarkdownFocusedElement")
            } else {
               label = renderer.GenerateNoRenderNode(x.label)
            }

            if x.isActivated {
               button = renderer.GenerateNode(x.button, "MarkdownToggleElementActive")
            } else {
               button = renderer.GenerateNode(x.button, "MarkdownToggleElement")
            }

            if button.Length() + label.Length() > width {
               log.Fatal("ERROR")
            }

            if used + button.Length() + label.Length() > width {
               flush()
            }

            l = append(l, button)
            l = append(l, label)
            used += button.Length() + label.Length()
         case readInputParseNode:
            var field renderer.Renderable

            s := fmt.Sprintf("%" + strconv.Itoa(x.maxLength) + "s", x.input)

            if x.IsFocused() {
               field = renderer.GenerateNode(s, "MarkdownInputFieldFocused")
            } else {
               field = renderer.GenerateNode(s, "MarkdownInputField")
            }

            if field.Length() > width {
               log.Fatal("ERROR")
            }

            if used + field.Length() > width {
               flush()
            }

            l = append(l, field)
            used += field.Length()

         case actionParseNode:
            var field renderer.Renderable
            if x.IsFocused() {
               field = renderer.GenerateNode(x.label, "MarkdownActionFieldFocused")
            } else {
               field = renderer.GenerateNode(x.label, "MarkdownActionField")

            }

            if used + field.Length() > width { 
               flush() }

            l = append(l, field)
            used += field.Length()


         case blockquoteparseNode:
            isBlockquote = true
            if len(l) > 0 { 
               flush() 
            } else {
               l = append(l, renderer.GenerateNode(" ", "MarkdownBlockquote"))
               l = append(l, renderer.GenerateNoRenderNode(" "))
               used += 2
            }

         case bareTextparseNode:
            s := x.GetContent() 

            if utf8.RuneCountInString(s) > width - used {
               l = append(l, renderer.GenerateNoRenderNode(s[:width-used]))
               s = s[width - used:]
               flush()
            }

            for utf8.RuneCountInString(s) > width {
               s1 := s[0:width - used - 1]
               s  = s[width - used - 1:]

               l = append(l, renderer.GenerateNoRenderNode(s1))
               flush()

            }
            a := renderer.GenerateNoRenderNode(s)
            l = append(l, a)
            used += a.Length()
      }
   }
   res = append(res, renderer.GenerateLine(width, l))

   return renderer.GenerateField(res)
}



type parseNode interface {
   GetContent()         string
}

type fullLineparseNode struct {
   content  string
   style    string
}
func (n fullLineparseNode) GetContent() string {return n.content}

type AtomarparseNode struct {
content  string
style    string
}
func (n AtomarparseNode) GetContent() string {return n.content}


type bareTextparseNode struct { content  string }
func (n bareTextparseNode) GetContent() string {return n.content}


type LineBreakparseNode struct {}
func (n LineBreakparseNode) GetContent() string {return ""}

type blockquoteparseNode struct {}
func (n blockquoteparseNode) GetContent() string {return ""}

type focusableParseNode interface {
   GetContent() string
   IsFocused() bool
   Focus() focusableParseNode
   Unfocus() focusableParseNode
}

/** Node to read user input **/
type readInputParseNode struct {
   input string
   id string
   isFocused bool
   maxLength int
}
func (ri readInputParseNode) GetContent() string {return ""}
func (ri readInputParseNode) IsFocused() bool { return ri.isFocused }
func (ri readInputParseNode) Focus() focusableParseNode {
   ri.isFocused = true
   return ri
}
func (ri readInputParseNode) Unfocus() focusableParseNode {
   ri.isFocused = false
   return ri
}
func (ri readInputParseNode) readInput() focusableParseNode {
   ri.input = ReadLine(true, ri.maxLength)
   return ri
}

// Ensure that the type implements the interface during compiletime
var _ focusableParseNode = (*readInputParseNode)(nil)

/** Button that can be toggled on and off **/
type toggleParseNode struct {
   button   string
   label    string
   id       string
   isActivated bool
   isFocused bool
}
func (tp toggleParseNode) GetContent() string { return tp.button + " " + tp.label }
func (tp toggleParseNode) GetId() string { return tp.id }
func (tp toggleParseNode) IsActive() bool { return tp.isActivated }
func (tp toggleParseNode) Activate() toggleParseNode { 
   tp.isActivated = true
   return tp
}
func (tp toggleParseNode) Deactivate() toggleParseNode { 
   tp.isActivated = false
   return tp
}
func (tp toggleParseNode) Toggle() toggleParseNode { 
   tp.isActivated = !tp.isActivated
   return tp
}
func (tp toggleParseNode) IsFocused() bool { return tp.isFocused }
func (tp toggleParseNode) Focus() focusableParseNode { 
   tp.isFocused = true
   return tp
}
func (tp toggleParseNode) Unfocus() focusableParseNode {
   tp.isFocused = false
   return tp
}

/** Action Parse node **/
type actionParseNode struct {
   label       string
   id          string
   isFocused   bool
}
func (a actionParseNode) GetContent() string { return a.label }
func (a actionParseNode) IsFocused() bool { return a.isFocused }
func (a actionParseNode) Focus() focusableParseNode {
   a.isFocused = true
   return a
}
func (a actionParseNode) Unfocus() focusableParseNode {
   a.isFocused = false
   return a
}
func (a actionParseNode) GetId() string { return a.id }



func parse(line string) []parseNode {
   { // Paragraph separators
      l := strings.Split(line, "\n\n")

      if len(l) > 1 {
         res := []parseNode{}
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(l[i])...)
            res = append(res, LineBreakparseNode{})
         }
      }
   }

   { // Capture linesbreaks
      l := regexp.MustCompile(`  \n`).FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, LineBreakparseNode{})
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   { // Parse lines individually 
      if strings.Contains(line, "\n") {
         l := strings.Split(line, "\n") 
            res := []parseNode{} 
            for _, x := range(l) {
               res = append(res, parse(x)...)
            }
            return res
      }
   }


   { // Capture headings or list elements
      captures := [8]*regexp.Regexp {
         regexp.MustCompile(`^\s*#[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*##[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*###[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*####[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*#####[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*######[ ]*([[:alnum:]]+)$`),
         regexp.MustCompile(`^\s*[0-9]+\..+$`),
         regexp.MustCompile(`^\s*-.+$`),
      }

      styles := [8]string {
         "MarkdownHeader1",
         "MarkdownHeader2",
         "MarkdownHeader3",
         "MarkdownHeader4",
         "MarkdownHeader5",
         "MarkdownHeader6",
         "MarkdownOrderedList",
         "MarkdownUnorderedList",
      }

      for j := 0; j < len(captures); j++ {
         if captures[j].MatchString(line) {
            return []parseNode{
               fullLineparseNode {
                  strings.Replace(line, "#", "", -1) + " ",
                  styles[j],
               },
            }
         }
      }
   }


   { // Blockquote
      if regexp.MustCompile(`^>.*$`).MatchString(line) {
         return append(
            []parseNode{blockquoteparseNode{}}, 
            parse(line[1:])...
         )
      }
   }


   { // Capture all bolditalic
      l := regexp.MustCompile(`\*\*\*[^\*]+\*\*\*`).FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarparseNode {
               content: line[l[i][0]+3:l[i][1]-3],
               style:   "bolditalic",
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }
   { // Capture all bold
      l := regexp.MustCompile(`\*\*[^\*]+\*\*`).FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarparseNode {
               content: line[l[i][0]+2:l[i][1]-2],
               style:   "bold",
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }
   { // Capture all italic
      l := regexp.MustCompile(`\*[^\*]+\*`).FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarparseNode {
               content: line[l[i][0]+1:l[i][1]-1],
               style:   "italic",
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   { // Capture all toggle elements
      reg := regexp.MustCompile(`<([ 0-9a-zA-Z]+)\|([ 0-9a-zA-Z]+)\|([0-9a-zA-Z]+)>`)
      l := reg.FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)

            a := reg.FindStringSubmatch(line[l[i][0]:l[i][1]])
         
            res = append(res, toggleParseNode {
               button: a[1],
               label: a[2],
               id: a[3],
               isActivated: false,
               isFocused: false,
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   { // Capture read line elements
      reg := regexp.MustCompile(`<\?([0-9a-zA-Z]+)\|([0-9]+)>`)
      l := reg.FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)

            a := reg.FindStringSubmatch(line[l[i][0]:l[i][1]])
         
            maxLength, err := strconv.Atoi(a[2])
            if err != nil { log.Fatal("ERROR")}
            res = append(res, readInputParseNode {
               id: a[1],
               maxLength: maxLength,
               isFocused: false,
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   { // Capture action elements
      reg := regexp.MustCompile(`<\!([ 0-9a-zA-Z]+)\|([ a-zA-Z0-9]+)>`)
      l := reg.FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []parseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)

            a := reg.FindStringSubmatch(line[l[i][0]:l[i][1]])
         
            res = append(res, actionParseNode {
               label: a[1],
               id: a[2],
               isFocused: false,
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   return []parseNode{bareTextparseNode{content: line}}
}


func GetMDDocument(base string) MDDocument {

   parsed := parse(base)

   index := -1
   for i, x := range(parsed) {
      a, ok := x.(focusableParseNode)

      if ok {
         index = i
         parsed[index] = a.Focus()
         break
      }
   }

   return MDDocument{
      base:          base,
      renderNodes:   parsed,
      isMultiselect: false,
      index:         index,
   }
}



