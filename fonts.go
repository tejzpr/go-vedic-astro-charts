// Copyright (c) 2024 Tejus Pratap <tejzpr@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package parashari

import (
	_ "embed"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/basicfont"
)

// Embed font files into the binary using go:embed
// These fonts will be included in the compiled binary when you build the application

//go:embed fonts/matangi/fonts/ttf/Matangi-Regular.ttf
var matangiRegularFont []byte

//go:embed fonts/matangi/fonts/ttf/Matangi-Bold.ttf
var matangiBoldFont []byte

// loadEmbeddedFont loads a font from embedded bytes and sets it on the context
// If loading fails, falls back to basic font
func loadEmbeddedFont(dc *gg.Context, fontData []byte, size float64) error {
	tt, err := opentype.Parse(fontData)
	if err != nil {
		dc.SetFontFace(basicfont.Face7x13)
		return err
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		dc.SetFontFace(basicfont.Face7x13)
		return err
	}

	dc.SetFontFace(face)
	return nil
}

// loadMatangiRegular loads Matangi Regular font from embedded data
func loadMatangiRegular(dc *gg.Context, size float64) {
	if err := loadEmbeddedFont(dc, matangiRegularFont, size); err != nil {
		// Fallback already set in loadEmbeddedFont
	}
}

// loadMatangiBold loads Matangi Bold font from embedded data
func loadMatangiBold(dc *gg.Context, size float64) {
	if err := loadEmbeddedFont(dc, matangiBoldFont, size); err != nil {
		// Fallback already set in loadEmbeddedFont
	}
}

