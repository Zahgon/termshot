// Copyright © 2020 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package img

import (
	"image"
	"image/color"
	"io"
	"os"

	"github.com/gonvenience/bunt"
	imgfont "golang.org/x/image/font"
)

const (
	red    = "#ED655A"
	yellow = "#E1C04C"
	green  = "#71BD47"
)

const (
	defaultFontSize = 12
	defaultFontDPI  = 144
)

// commandIndicator is the string to be used to indicate the command in the screenshot
var commandIndicator = func() string {
	if val, ok := os.LookupEnv("TS_COMMAND_INDICATOR"); ok {
		return val
	}

	return "➜"
}()

type Scaffold struct {
	content bunt.String

	factor float64

	columns int

	defaultForegroundColor color.Color

	clipCanvas bool

	drawDecorations bool
	drawShadow      bool

	shadowBaseColor string
	shadowRadius    uint8
	shadowOffsetX   float64
	shadowOffsetY   float64

	padding float64
	margin  float64

	regular     imgfont.Face
	bold        imgfont.Face
	italic      imgfont.Face
	boldItalic  imgfont.Face
	lineSpacing float64
	tabSpaces   int
}

func NewImageCreator() Scaffold { _ = "STUB: not implemented"; return *new(Scaffold) }

func (s *Scaffold) SetFontFaceRegular(face imgfont.Face) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetFontFaceBold(face imgfont.Face) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetFontFaceItalic(face imgfont.Face) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetFontFaceBoldItalic(face imgfont.Face) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetColumns(columns int) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetMargin(margin float64) { _ = "STUB: not implemented"; return }

func (s *Scaffold) SetPadding(padding float64) { _ = "STUB: not implemented"; return }

func (s *Scaffold) DrawDecorations(value bool) { _ = "STUB: not implemented"; return }

func (s *Scaffold) DrawShadow(value bool) { _ = "STUB: not implemented"; return }

func (s *Scaffold) ClipCanvas(value bool) { _ = "STUB: not implemented"; return }

func (s *Scaffold) GetFixedColumns() int { _ = "STUB: not implemented"; return 0 }

func (s *Scaffold) AddCommand(args ...string) error { _ = "STUB: not implemented"; return nil }

func (s *Scaffold) AddContent(in io.Reader) error { _ = "STUB: not implemented"; return nil }

// Add an additional newline in case the column
// count is reached and line wrapping is needed

func (s *Scaffold) fontHeight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *Scaffold) measureContent() (width float64, height float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// temporary drawer for reference calucation

// width, either by using longest line, or by fixed column value

// unlimited: max width of all lines

// fixed: max width based on column count

// height, lines times font height and line spacing

func (s *Scaffold) image() (image.Image, error) {
	_ = "STUB: not implemented"
	return *new(image.Image), nil
}

// Make sure the output window is big enough in case no content or very few
// content will be rendered

// Optional: Apply blurred rounded rectangle to mimic the window shadow
//

// Draw rounded rectangle with outline to produce impression of a window
//

// Optional: Draw window decorations (i.e. three buttons) to produce the
// impression of an actional window
//

// Apply the actual text into the prepared content area of the window
//

// background color
//nolint:gocritic

// #nosec G115
// #nosec G115
// #nosec G115

// foreground color

// #nosec G115
// #nosec G115
// #nosec G115

// mitigate issue #1 by replacing it with a similar character

// There seems to be no font face based way to do an underlined
// string, therefore manually draw a line under each character

// Write writes the scaffold content as PNG into the provided writer
//
// Deprecated: Use [Scaffold.WritePNG] instead.
func (s *Scaffold) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// WritePNG writes the scaffold content as PNG into the provided writer
func (s *Scaffold) WritePNG(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Optional: Clip image to minimum size by removing all surrounding transparent pixels
//

// WriteRaw writes the scaffold content as-is into the provided writer
func (s *Scaffold) WriteRaw(w io.Writer) error { _ = "STUB: not implemented"; return nil }
