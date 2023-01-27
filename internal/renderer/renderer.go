// ======================================================================
// Author: meisto
// Creation Date: Wed 18 Jan 2023 07:27:51 PM CET
// Description: -
// ======================================================================
package renderer

import (
   "log"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"waelder/internal/config"
)

type Renderable interface {
	Length() int
	Render(*termenv.Output, int, int)
}

type Alignment int32

const (
	Top    Alignment = iota
	Right  Alignment = iota
	Bottom Alignment = iota
	Left   Alignment = iota
)

type RenderField struct {
	height    int
	width     int
	alignment Alignment
	content   []RenderLine
}

func (r RenderField) Length() int { return r.height }
func (r RenderField) GetContentLength() int { return len(r.content) }
func (r RenderField) Render(output *termenv.Output, x, y int) {
	for i, a := range r.content {
		a.Render(output, x, y+i)
	}
}
func (r RenderField) RenderBlock(
	output *termenv.Output,
	x,
	y,
	height int,
	padTop bool,
   offset int,
) {
   

   if offset < 0 {offset = 0}

	if r.height - offset <= height { offset = r.height - height }

   // Paint content
   if r.height > height {
      for i := 0; i < height; i++ {
         r.content[i+offset].Render(output, x, y + i)
      }
   } else {
      // Whole block can fit so we can use offset otherwise
      offset = 0
      if padTop { offset = height - r.height }

      r.Render(output, x, y + offset)
   }
}

func GenerateField(content []RenderLine) RenderField {
   if len(content) == 0 {return RenderField{height:0, content: []RenderLine{}}}

	l := content[0].Length()
	for _, i := range content {
		if i.Length() != l {
			log.Fatal("[ERROR] Code: 81318510 | " + fmt.Sprintf("%d != %d\n", l, i.Length()))
		}
	}

	return RenderField{height: len(content), width: l, content: content}
}
func (r RenderField) Join(r2 RenderField) RenderField {
	if r.width != r2.width {
		log.Fatal("[ERROR] Code 31204189")
	}

	return GenerateField(append(r.content, r2.content...))
}
func (r RenderField) GetHeight() int { return r.height }
func (r RenderField) GetWidth() int  { return r.width }

type RenderLine struct {
	width   int
	content []Renderable
}

func (r RenderLine) Length() int {
	return r.width
}
func (r RenderLine) Render(output *termenv.Output, x, y int) {
	l := 0
	output.MoveCursor(y + 1, x + 1)
   output.WriteString(strings.Repeat(" ", r.width))
	output.MoveCursor(y + 1, x + 1)

	for _, k := range r.content {
		l2 := k.Length()

		if l+l2 > r.Length() {
			// Cut off overlong components
			output.WriteString("-")
			return
		}

		k.Render(output, x+l, y)
		l += l2
	}

}
func (r RenderLine) Join(r2 RenderLine) RenderLine {
	l := 0
	for _, i := range r.content {
		l += i.Length()
	}

	a := r.content
	if l < r.width {
		a = append(a, GenerateNoRenderNode(strings.Repeat(" ", r.width-l)))
	}

	return GenerateLine(
		r.width+r2.width,
		append(a, r2.content...),
	)
}

func GenerateLineFromOne(width int, content Renderable) RenderLine {
   return GenerateLine(width, []Renderable{content})
}

func GenerateLine(width int, content []Renderable) RenderLine {
	l := 0
	for _, c := range content {
		l += c.Length()
	}

	if l > width {
		log.Fatal("[ERROR] Code: 12308961, content too long.")
	}

	return RenderLine{width: width, content: content}
}

type RenderNode struct {
	text  string
	style lipgloss.Style
}

func (r RenderNode) Length() int {
	//NOTE: This returns the number of bytes not the length of the string
	// return len(r.text)

	return utf8.RuneCountInString(r.text)

}
func (r RenderNode) Render(output *termenv.Output, x, y int) {
	output.WriteString(r.style.Render(r.text))
}
func GenerateNode(text string, style string) RenderNode {
	return RenderNode{text: strings.Clone(text), style: config.GetStyle(style)}
}

type NoRenderNode struct{ text string }

func (r NoRenderNode) Length() int {
	//NOTE: This returns the number of bytes not the length of the string
	// return len(r.text)

	return utf8.RuneCountInString(r.text)
}
func (r NoRenderNode) Render(output *termenv.Output, x, y int) {
	output.WriteString(r.text)
}
func GenerateNoRenderNode(text string) NoRenderNode {
	return NoRenderNode{text: text}
}
