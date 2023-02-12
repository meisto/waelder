// ======================================================================
// Author: meisto
// Creation Date: Thu 26 Jan 2023 02:30:08 PM CET
// Description: -
// ======================================================================
package wio

import (
   "log"
   "regexp"
   "strings"
   "unicode/utf8"

   "waelder/internal/renderer"
)

type MDDocument struct {
   base     string
   renderNodes []ParseNode
   width    int
}
func (d MDDocument) Render() renderer.RenderField {
   res := []renderer.RenderLine{}  
   l := []renderer.Renderable{}
   used := 0

   flush := func() {
      i := 0
      for _, x := range(l) { i += x.Length() }
      log.Print(i, " ", d.width)

      res = append(res, renderer.GenerateLine(d.width, l))
      l = []renderer.Renderable{}
      used = 0
   }

   for i := 0; i < len(d.renderNodes); i++ {
      a := d.renderNodes[i]

      log.Print("Start")

      switch x := a.(type) {
         case LineBreakParseNode:
            flush()
         case FullLineParseNode:

            // Flush remaining documents
            if len(l) > 0 {
               flush()
            }
            res = append(res, renderer.GenerateLineFromOne(d.width, renderer.GenerateNode(x.content, x.style)))

         case AtomarParseNode:
            var a renderer.Renderable
            a = renderer.GenerateNode(x.content, x.style)

            if a.Length() > d.width {
               a = renderer.GenerateNoRenderNode("[PARSEERROR]")
            }

            if used + a.Length() >= d.width {
               flush()
            } 
            l = append(l, a)
            used += a.Length()

         case BareTextParseNode:
            s := x.GetContent() 

            if utf8.RuneCountInString(s) > d.width - used {
               l = append(l, renderer.GenerateNoRenderNode(s[:d.width-used]))
               s = s[d.width - used:]
               flush()
            }

            for utf8.RuneCountInString(s) > d.width - 3 {
               s1 := s[0:d.width - used - 1]
               s  = s[d.width - used - 1:]

               l = append(l, renderer.GenerateNoRenderNode(s1))
               flush()

            }
            a := renderer.GenerateNoRenderNode(s)
            l = append(l, a)
            used += a.Length()
      }
      log.Print("End")
   }
   res = append(res, renderer.GenerateLine(d.width, l))

   return renderer.GenerateField(res)
}



type ParseNode interface {
   GetContent()         string
}

type TodoParseNode struct {
   content  string
}
func (n TodoParseNode) GetContent() string {return n.content}

type FullLineParseNode struct {
   content  string
   style    string
}
func (n FullLineParseNode) GetContent() string {return n.content}

type AtomarParseNode struct {
content  string
style    string
}
func (n AtomarParseNode) GetContent() string {return n.content}


type BareTextParseNode struct { content  string }
func (n BareTextParseNode) GetContent() string {return n.content}

type LineBreakParseNode struct {}
func (n LineBreakParseNode) GetContent() string {return ""}





func parse(line string) []ParseNode {
   { // Paragraph separators
      l := strings.Split(line, "\n\n")

      if len(l) > 1 {
         res := []ParseNode{}
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(l[i])...)
         }
      }
   }

   { // Parse lines individually 
      if strings.Contains(line, "\n") {
         l := strings.Split(line, "\n") 
            res := []ParseNode{} 
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
            return []ParseNode{
               FullLineParseNode {
                  strings.Replace(line, "#", "", -1) + " ",
                  styles[j],
               },
            }
         }
      }
   }

   { // capture linebreaks
      if regexp.MustCompile(`(.*  $)`).MatchString(line) {
         return append(parse(line[:len(line)-2]), LineBreakParseNode{})
      }

   }


   // Linebreak

   // Blockquote
   regexp.MustCompile(`(^|\n)>`)


   { // Capture all bolditalic
      l := regexp.MustCompile(`\*\*\*[^\*]+\*\*\*`).FindAllStringIndex(line, -1)

      if len(l) > 0 {
         res := []ParseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarParseNode {
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
         res := []ParseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarParseNode {
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
         res := []ParseNode{}

         lowerBound := 0
         for i := 0; i < len(l); i ++ {
            res = append(res, parse(line[lowerBound:l[i][0]])...)
            res = append(res, AtomarParseNode {
               content: line[l[i][0]+1:l[i][1]-1],
               style:   "italic",
            })
            lowerBound = l[i][1]
         }
         res = append(res, parse(line[lowerBound:])...)
         return res
      }
   }

   return []ParseNode{BareTextParseNode{content: line}}
}


func GetMDDocument(base string, linewidth int) MDDocument {

   parsed := parse(base)

   return MDDocument{
      base:          base,
      renderNodes:   parsed,
      width: linewidth,
   }
}



