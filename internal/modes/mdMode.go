// ======================================================================
// Author: meisto
// Creation Date: Thu 26 Jan 2023 03:07:49 PM CET
// Description: -
// ======================================================================
package modes

import (
	"github.com/muesli/termenv"

	ds "waelder/internal/datastructures"
   "waelder/internal/wio"
   "waelder/internal/renderer"
)

var markdownRaw string = 
   "# Test\n" +
   "This is a Document\n"+
   "This is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\n" +
   "## Test2\n" +
   "### Test3\n" +
   "#### Test4\n" +
   "##### Test5\n" +
   "###### Test6\n" +
   "This is a italic*Line*\n" +
   "This is a bold**Line**\n" +
   "This is a bolditalic***Line***\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line1\n" +
   "This is a Li2ne\n" +
   "This is a Lin2e\n" +
   "This is a L41ine\n" +
   "This is a Li142ne\n" +
   "This is a Li14ne\n" +
   "This is a Li14ne\n" +
   "This is a Li142ne\n" +
   "This is a Li51ne\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is the last Line\n"

var markdown = wio.GetMDDocument(markdownRaw, 20)
func GetMarkdownLength() int { return markdown.Parsed.GetContentLength() }


func mdView(
	output *termenv.Output,
	d ds.Data,
	windowHeight int,
	windowWidth int,
) renderer.RenderField {

   markdown = wio.GetMDDocument(markdownRaw, windowWidth)

   return markdown.Parsed
}
