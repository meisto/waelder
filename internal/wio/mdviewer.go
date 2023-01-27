// ======================================================================
// Author: Tobias Meisel (meisto)
// Creation Date: Thu 26 Jan 2023 02:30:08 PM CET
// Description: -
// ======================================================================
package wio

import (
   "regexp"
   "strings"
   "unicode/utf8"

   "waelder/internal/renderer"
)

type MDDocument struct {
   base     string
   Parsed   renderer.RenderField
   width    int
}

func parseLine(line string, linewidth int) []renderer.RenderLine {
   if line == "" {
      return []renderer.RenderLine{renderer.GenerateLineFromOne(
         linewidth, renderer.GenerateNoRenderNode(""),
      )}
   }


   headers := []*regexp.Regexp{
      regexp.MustCompile("^[ \t]*#[ ]*([[:alnum:]]+)$"),
      regexp.MustCompile("^[ \t]*##[ ]*([[:alnum:]]+)$"),
      regexp.MustCompile("^[ \t]*###[ ]*([[:alnum:]]+)$"),
      regexp.MustCompile("^[ \t]*####[ ]*([[:alnum:]]+)$"),
      regexp.MustCompile("^[ \t]*#####[ ]*([[:alnum:]]+)$"),
      regexp.MustCompile("^[ \t]*######[ ]*([[:alnum:]]+)$"),
   }

   headerStyles := []string{
      "MarkdownHeader1", 
      "MarkdownHeader2",
      "MarkdownHeader3",
      "MarkdownHeader4",
      "MarkdownHeader5",
      "MarkdownHeader6",
   }

   { // Headers
      for i := 0; i < len(headers); i++ {
         if headers[i].MatchString(line) {
            m := headers[i].FindStringSubmatch(line)

            return []renderer.RenderLine{
               renderer.GenerateLine(
                  linewidth,
                  []renderer.Renderable{
                     renderer.GenerateNoRenderNode(" "),
                     renderer.GenerateNode(" " + m[1] + " ", headerStyles[i]),
                  },
               ),
            }
         } 
      }
   }

   // No previous match => just normal text
   italic      := regexp.MustCompile(`\*[^\*]+\*`)
   bold        := regexp.MustCompile(`\*\*[^\*]+\*\*`)
   bolditalic  := regexp.MustCompile(`\*\*\*[^\*]+\*\*\*`)

   // Parse all elements in one line
   res := []renderer.Renderable{}
   for utf8.RuneCountInString(line) > 0 {
      i := italic.FindStringIndex(line)
      b := bold.FindStringIndex(line)
      bi := bolditalic.FindStringIndex(line)

      var x []int = b
      if x == nil || ( i != nil &&  i[0] < x[0]) { x = i }
      if x == nil || (bi != nil && bi[0] < x[0]) { x = bi }

      if x == nil {
         // No more relevant element in line
         upperBound := utf8.RuneCountInString(line)
         if linewidth < upperBound {
            upperBound = linewidth
         }
         res = append(res, renderer.GenerateNoRenderNode(line[:upperBound]))
         line = line[upperBound:]

      } else {
         // Cut off possible leading strings
         if x[0] > 0 {
            m := x[0]
            if m > linewidth { m = linewidth }

            res = append(res, renderer.GenerateNoRenderNode(line[:m]))
            line = line[m:]

            // Jump out to prevent infinite loop
            continue
         }

         // Generate appropriate render node
         seg := line[x[0]:x[1]]
         line = line[x[1]:]
         if bolditalic.MatchString(seg) {
            res = append(res, renderer.GenerateNode(seg[3:len(seg)-3], "bolditalic"))
         } else if bold.MatchString(seg) {
            res = append(res, renderer.GenerateNode(seg[2:len(seg)-2], "bold"))

         } else if italic.MatchString(seg) {
            res = append(res, renderer.GenerateNode(seg[1:len(seg)-1], "italic"))
         }
      }
   }

   // Split res in lines with fitting length
   a := []renderer.RenderLine{}
   for len(res) > 0 {
      remainingLength := linewidth
      line := []renderer.Renderable{}
      for len(res) > 0 && res[0].Length() <= remainingLength {
         line = append(line, res[0])
         remainingLength -= res[0].Length()

         res = res[1:]     
      }
      a = append(a, renderer.GenerateLine(linewidth, line))
   }


   return a
}

func GetMDDocument(base string, linewidth int) MDDocument {
   baselines := strings.Split(base, "\n")
   parsedlines := []renderer.RenderLine{}

   for _, i := range(baselines) {
      parsedlines = append(parsedlines, parseLine(i, linewidth)...)
   }

   return MDDocument{
      base: base,
      Parsed: renderer.GenerateField(parsedlines),
      width: linewidth,
   }
}



