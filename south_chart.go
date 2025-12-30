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
	"fmt"
	"image"
	"strings"

	"github.com/fogleman/gg"
)

// GenerateSouthChart generates a South Indian style chart
// Houses are fixed, rashis rotate based on Lagna
func GenerateSouthChart(input ChartInput) ([]byte, error) {
	const size = 800
	const padding = 40
	const gridSize = size - 2*padding

	dc := gg.NewContext(size, size)
	dc.SetRGB(1, 1, 1) // White background
	dc.Clear()

	// Draw outer square
	dc.SetRGB(0, 0, 0) // Black lines
	dc.SetLineWidth(2)
	dc.DrawRectangle(float64(padding), float64(padding), float64(gridSize), float64(gridSize))
	dc.Stroke()

	// Calculate cell size (4x4 grid = 16 cells, but we use 12 houses around perimeter)
	cellSize := float64(gridSize) / 4

	// STEP 1, 2, 3 & 4: Draw Houses 1-4
	// Top row: House 1 (Aries), House 2 (Taurus), House 3 (Gemini)
	// Right side: House 3 (corner), House 4 (Cancer) below House 3

	dc.SetLineWidth(1)

	// Draw the boundaries for top row houses
	// Left edge: vertical line at x = padding + cellSize (from top to first horizontal line)
	x1 := float64(padding) + cellSize
	dc.DrawLine(x1, float64(padding), x1, float64(padding)+cellSize)
	dc.Stroke()

	// Right edge of House 1 (also left edge of House 2): vertical line at x = padding + 2*cellSize
	x2 := float64(padding) + 2*cellSize
	dc.DrawLine(x2, float64(padding), x2, float64(padding)+cellSize)
	dc.Stroke()

	// Right edge of House 2 (also left edge of House 3): vertical line at x = padding + 3*cellSize
	x3 := float64(padding) + 3*cellSize
	// Top part: from top to first horizontal line (House 3)
	dc.DrawLine(x3, float64(padding), x3, float64(padding)+cellSize)
	dc.Stroke()
	// Bottom part: from first horizontal line to second horizontal line (left edge of House 4)
	dc.DrawLine(x3, float64(padding)+cellSize, x3, float64(padding)+2*cellSize)
	dc.Stroke()

	// Right edge of House 3 (also right edge of Houses 4, 5, and 6): vertical line at x = padding + 4*cellSize (outer edge)
	// This is the right side of the chart, so draw from top to bottom
	x4 := float64(padding) + 4*cellSize
	// Top part: from top to first horizontal line (House 3)
	dc.DrawLine(x4, float64(padding), x4, float64(padding)+cellSize)
	dc.Stroke()
	// Middle part: from first horizontal line to bottom (Houses 4, 5, and 6)
	dc.DrawLine(x4, float64(padding)+cellSize, x4, float64(padding)+4*cellSize)
	dc.Stroke()

	// Left edge of House 5: extend x3 line down to third horizontal line
	dc.DrawLine(x3, float64(padding)+2*cellSize, x3, float64(padding)+3*cellSize)
	dc.Stroke()

	// Left edge of House 6: extend x3 line down to bottom
	dc.DrawLine(x3, float64(padding)+3*cellSize, x3, float64(padding)+4*cellSize)
	dc.Stroke()

	// Left edge of House 7 (also right edge of House 8): vertical line at x = padding + 2*cellSize (from third horizontal line to bottom)
	x2Bottom := float64(padding) + 2*cellSize
	dc.DrawLine(x2Bottom, float64(padding)+3*cellSize, x2Bottom, float64(padding)+4*cellSize)
	dc.Stroke()

	// Left edge of House 8 (also right edge of House 9): vertical line at x = padding + cellSize (from third horizontal line to bottom)
	x1Bottom := float64(padding) + cellSize
	dc.DrawLine(x1Bottom, float64(padding)+3*cellSize, x1Bottom, float64(padding)+4*cellSize)
	dc.Stroke()

	// Left edge of House 9: vertical line at x = padding (from third horizontal line to bottom)
	// This is the left edge of the chart, already part of outer square, but we need the bottom part
	x0Bottom := float64(padding)
	dc.DrawLine(x0Bottom, float64(padding)+3*cellSize, x0Bottom, float64(padding)+4*cellSize)
	dc.Stroke()

	// Left edge of House 10: vertical line at x = padding (from second horizontal line to third horizontal line)
	// This is the left edge of the chart, extend upward
	dc.DrawLine(x0Bottom, float64(padding)+2*cellSize, x0Bottom, float64(padding)+3*cellSize)
	dc.Stroke()

	// Left edge of House 11: vertical line at x = padding (from first horizontal line to second horizontal line)
	// This is the left edge of the chart, extend further upward
	dc.DrawLine(x0Bottom, float64(padding)+cellSize, x0Bottom, float64(padding)+2*cellSize)
	dc.Stroke()

	// Left edge of House 12: vertical line at x = padding (from top to first horizontal line)
	// This is the left edge of the chart, top-left corner
	dc.DrawLine(x0Bottom, float64(padding), x0Bottom, float64(padding)+cellSize)
	dc.Stroke()

	// Right edge of House 10 (also right edge of House 11): vertical line at x = padding + cellSize (from first horizontal line to third horizontal line)
	// This is also the left edge of House 12 and right edge of House 1
	dc.DrawLine(x1Bottom, float64(padding)+cellSize, x1Bottom, float64(padding)+3*cellSize)
	dc.Stroke()

	// Right edge of House 12: vertical line at x = padding + cellSize (from top to first horizontal line)
	// This is also the left edge of House 1
	dc.DrawLine(x1Bottom, float64(padding), x1Bottom, float64(padding)+cellSize)
	dc.Stroke()

	// Top edge: already part of outer square
	// Bottom edge of top row: horizontal line at y = padding + cellSize (from left edge of House 1 to right edge)
	y1 := float64(padding) + cellSize
	// Right part: from x1 to x4 (bottom edge of top row houses 1, 2, 3)
	dc.DrawLine(x1, y1, x4, y1)
	dc.Stroke()
	// Left part: from x0Bottom to x1Bottom (top edge of House 11, bottom edge of House 12)
	dc.DrawLine(x0Bottom, y1, x1Bottom, y1)
	dc.Stroke()

	// Bottom edge of House 4: horizontal line at y = padding + 2*cellSize (from left edge to right edge of House 4)
	y2 := float64(padding) + 2*cellSize
	// Right part: from x3 to x4 (bottom edge of House 4)
	dc.DrawLine(x3, y2, x4, y2)
	dc.Stroke()
	// Left part: from x0Bottom to x1Bottom (top edge of House 10, bottom edge of House 11)
	dc.DrawLine(x0Bottom, y2, x1Bottom, y2)
	dc.Stroke()

	// Top edge of Houses 7, 8, and 9 (also bottom edge of House 5 and House 10): horizontal line at y = padding + 3*cellSize
	// This line goes from left edge of House 9 to right edge (separating House 5 from Houses 7, 8, and 9, and House 10 from House 9)
	y3 := float64(padding) + 3*cellSize
	// Left part: from x0Bottom to x1Bottom (top edge of House 9, bottom edge of House 10)
	dc.DrawLine(x0Bottom, y3, x1Bottom, y3)
	dc.Stroke()
	// Middle-left part: from x1Bottom to x2Bottom (top edge of House 8)
	dc.DrawLine(x1Bottom, y3, x2Bottom, y3)
	dc.Stroke()
	// Middle-right part: from x2Bottom to x3 (top edge of House 7)
	dc.DrawLine(x2Bottom, y3, x3, y3)
	dc.Stroke()
	// Right part: from x3 to x4 (bottom edge of House 5)
	dc.DrawLine(x3, y3, x4, y3)
	dc.Stroke()

	// Bottom edge of Houses 6, 7, 8, and 9: horizontal line at y = padding + 4*cellSize (from left edge of House 9 to right edge of House 6)
	// This is the bottom of the chart, already part of outer square
	y4 := float64(padding) + 4*cellSize
	dc.DrawLine(x0Bottom, y4, x4, y4)
	dc.Stroke()

	// Find Lagna rashi
	// For South Indian charts, rashi numbers are FIXED positions:
	// 1=Aries, 2=Taurus, 3=Gemini, ..., 8=Scorpio, ..., 12=Pisces
	// These numbers don't change - they're always in the same positions
	var lagnaRashi int

	// Get lagna rashi from input parameter
	if input.Lagna != nil {
		lagnaRashi = RashiToNumber(input.Lagna.Rashi)
	}

	// If lagna not provided or invalid, default to Aries
	if lagnaRashi == 0 {
		lagnaRashi = 1
	}

	// House positions as rectangles (arranged around perimeter)
	// Top row: 12 (left), 1 (left-center), 2 (right-center), 3 (right corner)
	// Right side: 3 (corner), 4 (top), 5 (middle), 6 (bottom corner)
	// Bottom row: 6 (corner), 7 (right-center), 8 (left-center), 9 (left corner)
	// Left side: 9 (corner), 10 (bottom), 11 (middle), 12 (top corner)
	houseRects := map[int]image.Rectangle{
		// Top row (left to right)
		12: image.Rect(int(padding), int(padding), int(padding+cellSize), int(padding+cellSize)),                   // Top-left corner
		1:  image.Rect(int(padding)+int(cellSize), int(padding), int(padding+2*cellSize), int(padding+cellSize)),   // Top left-center
		2:  image.Rect(int(padding)+int(2*cellSize), int(padding), int(padding+3*cellSize), int(padding+cellSize)), // Top right-center
		3:  image.Rect(int(padding)+int(3*cellSize), int(padding), int(padding+4*cellSize), int(padding+cellSize)), // Top-right corner

		// Right side (top to bottom, excluding corners)
		4: image.Rect(int(padding)+int(3*cellSize), int(padding)+int(cellSize), int(padding+4*cellSize), int(padding+2*cellSize)),   // Right top
		5: image.Rect(int(padding)+int(3*cellSize), int(padding)+int(2*cellSize), int(padding+4*cellSize), int(padding+3*cellSize)), // Right middle
		// House 6 is bottom-right corner (shared with bottom row)

		// Bottom row (right to left)
		6: image.Rect(int(padding)+int(3*cellSize), int(padding)+int(3*cellSize), int(padding+4*cellSize), int(padding+4*cellSize)), // Bottom-right corner
		7: image.Rect(int(padding)+int(2*cellSize), int(padding)+int(3*cellSize), int(padding+3*cellSize), int(padding+4*cellSize)), // Bottom right-center
		8: image.Rect(int(padding)+int(cellSize), int(padding)+int(3*cellSize), int(padding+2*cellSize), int(padding+4*cellSize)),   // Bottom left-center
		9: image.Rect(int(padding), int(padding)+int(3*cellSize), int(padding+cellSize), int(padding+4*cellSize)),                   // Bottom-left corner

		// Left side (bottom to top, excluding corners)
		10: image.Rect(int(padding), int(padding)+int(2*cellSize), int(padding+cellSize), int(padding+3*cellSize)), // Left bottom
		11: image.Rect(int(padding), int(padding)+int(cellSize), int(padding+cellSize), int(padding+2*cellSize)),   // Left middle
		// House 12 is top-left corner (already defined above)
	}

	// Draw rashi numbers and planets in each house
	dc.SetRGB(0, 0, 0)
	// Load Matangi font for rashi numbers from embedded data
	loadMatangiRegular(dc, 16)

	// STEP 1-12: Draw all 12 Houses
	// In South Indian charts, rashi numbers are FIXED positions:
	// Position 1 = Aries (1), Position 2 = Taurus (2), ..., Position 8 = Scorpio (8), etc.
	// These numbers never change - they're always in the same positions
	for houseNum := 1; houseNum <= 12; houseNum++ {
		rect := houseRects[houseNum]
		// Rashi number is always the same as the position (1-12)
		// Position 1 = Aries (1), Position 2 = Taurus (2), etc.
		rashiNum := houseNum

		// Draw rashi number (no L marker) - always display the rashi number
		rashiStr := fmt.Sprintf("%d", rashiNum)

		// Position text in bottom-right of the rectangle
		// Use bottom-right anchor with some padding from edges
		// Move text up more to avoid being crossed by bottom border
		textX := float64(rect.Max.X) - 10
		textY := float64(rect.Max.Y) - 29 // Moved up by another 2px (was 27, now 29)

		// Ensure rashi number is drawn in black
		dc.SetRGB(0, 0, 0)
		// Draw rashi number (anchored to bottom-right)
		dc.DrawStringAnchored(rashiStr, textX, textY, 1.0, 1.0)

		// Draw two parallel diagonal lines at bottom-left corner if this is the lagna rashi position
		// These form parallel diagonal lines (like //) at the corner
		if input.Lagna != nil && lagnaRashi > 0 && rashiNum == lagnaRashi {
			cornerX := float64(rect.Min.X) + 15 // Left border + 15px offset
			cornerY := float64(rect.Max.Y)      // Bottom border
			lineLength := 15.0                  // Length of each diagonal line
			offset := 3.0                       // Offset between the two parallel lines

			// Rotate by 270 degrees total (90 + 90 + 90): use rotation matrix
			// Original direction is -45 degrees, rotate by 270 degrees total
			cos90 := 0.0 // cos(90°)
			sin90 := 1.0 // sin(90°)

			// Original direction vector (lineLength, -lineLength)
			// First rotate by 90 degrees
			dx1 := lineLength*cos90 - (-lineLength)*sin90
			dy1 := lineLength*sin90 + (-lineLength)*cos90
			// Then rotate by another 90 degrees
			dx2 := dx1*cos90 - dy1*sin90
			dy2 := dx1*sin90 + dy1*cos90
			// Then rotate by another 90 degrees
			dx := dx2*cos90 - dy2*sin90
			dy := dx2*sin90 + dy2*cos90

			dc.SetLineWidth(2)
			// First diagonal: rotated line from bottom-left corner
			dc.DrawLine(cornerX, cornerY, cornerX+dx, cornerY+dy)
			dc.Stroke()
			// Second diagonal: parallel line, slightly offset
			dc.DrawLine(cornerX+offset, cornerY-offset, cornerX+dx+offset, cornerY+dy-offset)
			dc.Stroke()
			dc.SetLineWidth(1) // Reset line width
		}

		// Collect planets, grahas, and upagrahas in this house based on their Rashi
		// Planets should be placed in the house that contains their rashi
		var regularPlanets []string
		var specialLagnas []string

		// Add planets and lagna - treat lagna just like any other planet
		// First add lagna if this is the lagna rashi position
		if input.Lagna != nil && lagnaRashi > 0 && rashiNum == lagnaRashi {
			abbrev := GetPlanetDisplayName("lagna", input.Lagna)
			// Lagna is never retrograde or combust (it's a point, not a planet)
			regularPlanets = append(regularPlanets, abbrev)
		}

		// Add regular planets and separate special lagnas
		for planetName, planet := range input.Planets {
			planetRashiNum := RashiToNumber(planet.Rashi)
			// Check if this planet's rashi matches the rashi number of this position
			if planetRashiNum > 0 && planetRashiNum == rashiNum {
				abbrev := GetPlanetDisplayName(planetName, planet)

				if planet.IsRetrograde {
					abbrev += "R"
				}
				if planet.IsCombust {
					abbrev += "C"
				}

				// Separate special lagnas from regular planets
				if IsSpecialLagnaAbbrev(abbrev, input) {
					specialLagnas = append(specialLagnas, abbrev)
				} else {
					regularPlanets = append(regularPlanets, abbrev)
				}
			}
		}

		// Draw planets in top center of the box with larger font
		// Load larger Matangi font for planets from embedded data
		loadMatangiBold(dc, 22)
		centerX := float64(rect.Min.X+rect.Max.X) / 2 // Center horizontally
		planetY := float64(rect.Min.Y) + 25           // Top with padding

		// Calculate spacing: planets on left, special lagnas on right
		leftX := centerX - 25  // Left side for regular planets
		rightX := centerX + 25 // Right side for special lagnas

		// Draw regular planets on the left
		for i, planetAbbrev := range regularPlanets {
			// Check if this is Ascendant and set color to saffron
			if strings.Contains(planetAbbrev, "Asc") {
				dc.SetRGB(1.0, 0.6, 0.2) // Saffron
			} else {
				dc.SetRGB(0, 0, 0) // Black
			}
			dc.DrawStringAnchored(planetAbbrev, leftX, planetY+float64(i*25), 1.0, 0.5)
		}

		// Draw special lagnas on the right, matching up with planets by index
		maxItems := len(regularPlanets)
		if len(specialLagnas) > maxItems {
			maxItems = len(specialLagnas)
		}

		for i := 0; i < maxItems; i++ {
			// Draw special lagna if available at this index
			if i < len(specialLagnas) {
				dc.SetRGB(1.0, 0.85, 0.0) // Yellow for special lagnas
				dc.DrawStringAnchored(specialLagnas[i], rightX, planetY+float64(i*25), 0.0, 0.5)
			}
		}
		// Reset color back to black after drawing planets
		dc.SetRGB(0, 0, 0)
		// Reset font back to smaller size for rashi numbers
		loadMatangiRegular(dc, 16)
	}

	// Draw center text if provided
	if input.CenterText != "" {
		// Center of the chart (the 4 empty squares in the middle)
		centerX := float64(padding) + 2*cellSize
		centerY := float64(padding) + 2*cellSize

		// Load font for center text from embedded data
		loadMatangiRegular(dc, 18)

		dc.SetRGB(0, 0, 0) // Black text

		// Split text by newlines and draw each line
		lines := strings.Split(input.CenterText, "\n")
		lineHeight := 25.0                                     // Height between lines
		startY := centerY - float64(len(lines)-1)*lineHeight/2 // Center vertically

		for i, line := range lines {
			if line != "" { // Skip empty lines
				dc.DrawStringAnchored(line, centerX, startY+float64(i)*lineHeight, 0.5, 0.5)
			}
		}
	}

	return encodePNG(dc.Image())
}
