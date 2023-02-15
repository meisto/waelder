// ======================================================================
// Author: meisto
// Creation Date: Thu 26 Jan 2023 03:07:49 PM CET
// Description: -
// ======================================================================
package modes

import (
	ds "waelder/internal/datastructures"
   "waelder/internal/wio"
   "waelder/internal/renderer"
)


var markdownRaw string = 
   "# Test\n" +
   "This is a Document\n"+
   "TThis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\nThis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\nThis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\nThis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\nThis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\nhis is a Documenttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt\n" +
   "## Test2\n" +
   "### Test3\n" +
   "#### Test4\n" +
   "##### Test5\n" +
   "###### Test6\n" +
   "This is a italic*Line*\n" +
   "This is a bold**Line**\n" +
   "This is a bolditalic***Line***\n" +
   "This is a Line\n" +
   ">This is a Blockquote\n" +
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
   "This is the Li51ne we watch\n" +
   "This is a Line  <!this is a action node|actionnode>  \n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a Line\n" +
   "This is a toggleable box  \n" +
   "< x | Label   |589>  \n" +
   "< x | Label2  |581> \n" +
   "< x | Label3  |5813>  \n" +
   "This is a toggleable box" +
   "This is a Line\n" +
   "1. This is a Line\n" +
   "2. This is a input field:\n" +
   "123 " +
   "<?213|3>  \n" +
   "3. This is a Line\n" +
   "- This is a Line\n" +
   "- This is a Line\n" +
   "- This is a Line\n" +
   "- This is a Line\n" +
   "This is the last Line\n"

var documentLookup map[string]wio.MDDocument = map[string]wio.MDDocument{
   "test": wio.GetMDDocument(markdownRaw),
   "test2": wio.GetMDDocument(wio.ReadLocalFileToString("data/test.md")),
}

var Markdown wio.MDDocument = documentLookup["test2"]


func mdView(
	d ds.Data,
	windowHeight int,
	windowWidth int,
) renderer.RenderField {


   return Markdown.Render(windowWidth)
}
