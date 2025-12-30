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
	"math"
	"strings"

	"github.com/fogleman/gg"
)

// GenerateNorthChart generates a North Indian style chart
// Fixed zodiac signs, houses move based on lagna (counter-clockwise)
func GenerateNorthChart(input ChartInput) ([]byte, error) {
	const size = 800
	const padding = 40
	const chartSize = float64(size - 2*padding)
	const centerX = float64(size) / 2
	const centerY = float64(size) / 2

	dc := gg.NewContext(size, size)
	dc.SetRGB(1, 1, 1) // White background
	dc.Clear()

	// Step 1: Define inner square (rotated 45 degrees)
	// Expand by 50% then another 15% then another 5%, then reduce by 2%: multiply by 1.5 * 1.15 * 1.05 * 0.98
	innerSquareSize := chartSize * 0.4 * 1.5 * 1.15 * 1.05 * 0.98
	innerHalfSize := innerSquareSize / 2

	// Step 2: Calculate outer square size
	// The center of each edge of the outer square should touch each corner vertex of the inner square
	// When inner square is rotated 45 degrees, distance from center to corner = innerHalfSize * sqrt(2)
	// The outer square's edge midpoints should be at this distance from center
	innerCornerDistance := innerHalfSize * math.Sqrt(2)
	outerHalfSize := innerCornerDistance

	// Step 3: Draw outer square (rotated 45 degrees)
	dc.SetRGB(0, 0, 0) // Black lines
	dc.SetLineWidth(3)

	dc.Push()
	dc.Translate(centerX, centerY)
	dc.Rotate(90 * math.Pi / 180) // Rotate 90 degrees clockwise (45 + 45)
	dc.DrawRectangle(-outerHalfSize, -outerHalfSize, outerHalfSize*2, outerHalfSize*2)
	dc.Stroke()
	dc.Pop()

	// Step 4: Draw inner square (rotated 45 degrees counter-clockwise)
	dc.SetLineWidth(2)
	dc.Push()
	dc.Translate(centerX, centerY)
	dc.Rotate(-45 * math.Pi / 180) // Rotate 45 degrees counter-clockwise
	dc.DrawRectangle(-innerHalfSize, -innerHalfSize, innerSquareSize, innerSquareSize)
	dc.Stroke()

	// Step 5: Draw two lines splitting each side of the inner square by 2
	// Extend these lines all the way to the outer square vertices
	// The outer square vertices are at distance outerHalfSize * sqrt(2) from center (diagonal distance)
	// Since we're in the inner square's rotated coordinate system, we need to extend far enough
	// to reach the outer square vertices. The outer square is rotated 90°, so its vertices
	// in global coordinates are at (outerHalfSize, 0), (0, outerHalfSize), etc.
	// In the inner square's coordinate system (-45°), we need to extend to reach these points.
	// A safe distance is the diagonal of the outer square: outerHalfSize * sqrt(2)
	extendDistance := outerHalfSize * math.Sqrt(2)

	// Line 1: horizontal line extending from inner square edge to outer square vertices
	dc.DrawLine(-extendDistance, 0, extendDistance, 0)
	dc.Stroke()
	// Line 2: vertical line extending from inner square edge to outer square vertices
	dc.DrawLine(0, -extendDistance, 0, extendDistance)
	dc.Stroke()

	dc.Pop()

	// Step 5a: Display Lagna rashi number (first number) at coordinates (400, 300)
	// Find Lagna rashi number
	var lagnaRashiNum int
	if input.Lagna != nil {
		lagnaRashiNum = RashiToNumber(input.Lagna.Rashi)
	}
	if lagnaRashiNum == 0 {
		lagnaRashiNum = 1 // Default to Aries
	}

	// Draw rashi number at global coordinates (400, 300)
	dc.SetRGB(0, 0, 0) // Black text
	// Load Matangi font from embedded data
	loadMatangiRegular(dc, 20)
	rashiStr := fmt.Sprintf("%d", lagnaRashiNum)
	// Position at coordinates (400, 300) in global coordinate system
	textX := 400.0
	textY := 300.0
	// Rotate the text by 5 degrees (was 15, now -10 counter-clockwise = 5)
	dc.Push()
	dc.Translate(textX, textY)
	dc.Rotate(5 * math.Pi / 180)                    // Rotate 5 degrees
	dc.DrawStringAnchored(rashiStr, 0, 0, 0.5, 0.5) // Center-aligned
	dc.Pop()

	// Step 6: Format the chart - add rashi numbers, planets, and lagna
	// In North Indian charts: FIXED ZODIAC (signs stay fixed), HOUSES MOVE (counter-clockwise)

	// Find Lagna rashi
	var lagnaRashi int
	if input.Lagna != nil {
		lagnaRashi = RashiToNumber(input.Lagna.Rashi)
	}
	if lagnaRashi == 0 {
		lagnaRashi = 1
	}

	// Define fixed positions for all 12 rashi numbers (counter-clockwise from lagna position)
	// Position 1 is lagna (already drawn above)
	// Positions 2-12 are the remaining positions
	rashiPositions := []struct {
		x, y             float64 // Rashi number position
		angle            float64 // Rotation angle in degrees
		planetX, planetY float64 // Planet display position
	}{
		// Position 2 (next after lagna, counter-clockwise)
		{220.0, 160.0, -5.0, 180.0, 70.0},
		// Position 3
		{70.0, 300.0, -1.0, 60.0, 150.0},
		// Position 4
		{220.0, 500.0, -1.0, 200.0, 310.0},
		// Position 5
		{70.0, 670.0, -1.0, 60.0, 500.0},
		// Position 6
		{130.0, 720.0, -1.0, 180.0, 640.0},
		// Position 7
		{400.0, 680.0, -1.0, 380.0, 480.0},
		// Position 8
		{650.0, 725.0, -1.0, 540.0, 660.0},
		// Position 9
		{730.0, 660.0, -1.0, 690.0, 500.0},
		// Position 10
		{580.0, 500.0, -1.0, 550.0, 330.0},
		// Position 11
		{720.0, 300.0, -1.0, 700.0, 130.0},
		// Position 12
		{580.0, 160.0, -1.0, 520.0, 70.0},
	}

	// Set up font for rashi numbers
	dc.SetRGB(0, 0, 0)
	// Load Matangi font from embedded data
	loadMatangiRegular(dc, 20)

	// Helper function to get rashi number for a position
	getRashiForPosition := func(position int) int {
		if position == 1 {
			return lagnaRashiNum
		}
		offset := position - 1
		rashiNum := (lagnaRashiNum + offset) % 12
		if rashiNum == 0 {
			rashiNum = 12
		}
		return rashiNum
	}

	// Draw rashi numbers in positions 2-12 and collect planets for each position
	// Position 1 is lagna, position 2 is lagna+1, position 3 is lagna+2, etc. (counter-clockwise)
	for i, pos := range rashiPositions {
		// Position number (2-12, where position 1 is lagna)
		// Position 2 should be lagna + 1, position 3 should be lagna + 2, etc.
		offset := i + 1 // Position 2 has offset 1, position 3 has offset 2, etc.
		rashiNum := (lagnaRashiNum + offset) % 12
		if rashiNum == 0 {
			rashiNum = 12
		}

		dc.Push()
		dc.Translate(pos.x, pos.y)
		dc.Rotate(pos.angle * math.Pi / 180)
		rashiStr := fmt.Sprintf("%d", rashiNum)
		dc.DrawStringAnchored(rashiStr, 0, 0, 0.5, 0.5) // Center-aligned
		dc.Pop()
	}

	// Now draw planets near each rashi number position
	// Load larger font for planets from embedded data
	loadMatangiBold(dc, 18)

	// Draw planets for position 1 (lagna position)
	position1Rashi := getRashiForPosition(1)
	regularPlanets1 := []string{}
	specialLagnas1 := []string{}

	// Add lagna if it's in this rashi
	if input.Lagna != nil && position1Rashi == lagnaRashiNum {
		abbrev := GetPlanetDisplayName("lagna", input.Lagna)
		// Lagna is never retrograde or combust (it's a point, not a planet)
		regularPlanets1 = append(regularPlanets1, abbrev)
	}

	// Add regular planets in this rashi, separate special lagnas
	for planetName, planet := range input.Planets {
		planetRashiNum := RashiToNumber(planet.Rashi)
		if planetRashiNum > 0 && planetRashiNum == position1Rashi {
			abbrev := GetPlanetDisplayName(planetName, planet)
			if planet.IsRetrograde {
				abbrev += "R"
			}
			if planet.IsCombust {
				abbrev += "C"
			}
			
			// Separate special lagnas from regular planets
			if IsSpecialLagnaAbbrev(abbrev, input) {
				specialLagnas1 = append(specialLagnas1, abbrev)
			} else {
				regularPlanets1 = append(regularPlanets1, abbrev)
			}
		}
	}

	// Draw planets near position 1 (lagna position at 400, 300)
	if len(regularPlanets1) > 0 || len(specialLagnas1) > 0 {
		leftX := 360.0  // Left side for regular planets
		rightX := 400.0 // Right side for special lagnas
		planetY := 140.0
		
		// Draw regular planets on the left
		for i, planetAbbrev := range regularPlanets1 {
			// Check if this is Ascendant and set color to saffron
			if strings.Contains(planetAbbrev, "Asc") {
				dc.SetRGB(1.0, 0.6, 0.2) // Saffron
			} else {
				dc.SetRGB(0, 0, 0) // Black
			}
			dc.DrawStringAnchored(planetAbbrev, leftX, planetY+float64(i*20), 1.0, 0.5)
		}
		
		// Draw special lagnas on the right, matching up with planets by index
		maxItems := len(regularPlanets1)
		if len(specialLagnas1) > maxItems {
			maxItems = len(specialLagnas1)
		}
		
		for i := 0; i < maxItems; i++ {
			// Draw special lagna if available at this index
			if i < len(specialLagnas1) {
				dc.SetRGB(1.0, 0.85, 0.0) // Yellow for special lagnas
				dc.DrawStringAnchored(specialLagnas1[i], rightX, planetY+float64(i*20), 0.0, 0.5)
			}
		}
		dc.SetRGB(0, 0, 0) // Reset to black
	}

	// Draw planets for positions 2-12
	for i, pos := range rashiPositions {
		positionNum := i + 2
		rashiNum := getRashiForPosition(positionNum)

		regularPlanets := []string{}
		specialLagnas := []string{}

		// Add lagna if it's in this rashi
		if input.Lagna != nil && rashiNum == lagnaRashiNum {
			abbrev := GetPlanetDisplayName("lagna", input.Lagna)
			// Lagna is never retrograde or combust (it's a point, not a planet)
			regularPlanets = append(regularPlanets, abbrev)
		}

		// Add regular planets in this rashi, separate special lagnas
		for planetName, planet := range input.Planets {
			planetRashiNum := RashiToNumber(planet.Rashi)
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

		// Draw planets near this rashi number
		if len(regularPlanets) > 0 || len(specialLagnas) > 0 {
			// Use specific planet position if set, otherwise calculate offset
			var baseX, baseY float64
			if pos.planetX != 0.0 || pos.planetY != 0.0 {
				baseX = pos.planetX
				baseY = pos.planetY
			} else {
				// Calculate offset position for planets (to the right of the number)
				// Use the rotation angle to determine offset direction
				angleRad := pos.angle * math.Pi / 180
				offsetX := 30.0 * math.Cos(angleRad)
				offsetY := 30.0 * math.Sin(angleRad)
				baseX = pos.x + offsetX
				baseY = pos.y + offsetY
			}

			// Calculate left and right positions
			leftX := baseX - 20  // Left side for regular planets
			rightX := baseX + 20 // Right side for special lagnas

			// Draw regular planets on the left
			for j, planetAbbrev := range regularPlanets {
				// Check if this is Ascendant and set color to saffron
				if strings.Contains(planetAbbrev, "Asc") {
					dc.SetRGB(1.0, 0.6, 0.2) // Saffron
				} else {
					dc.SetRGB(0, 0, 0) // Black
				}
				dc.DrawStringAnchored(planetAbbrev, leftX, baseY+float64(j*20), 1.0, 0.5)
			}

			// Draw special lagnas on the right, matching up with planets by index
			maxItems := len(regularPlanets)
			if len(specialLagnas) > maxItems {
				maxItems = len(specialLagnas)
			}

			for j := 0; j < maxItems; j++ {
				// Draw special lagna if available at this index
				if j < len(specialLagnas) {
					dc.SetRGB(1.0, 0.85, 0.0) // Yellow for special lagnas
					dc.DrawStringAnchored(specialLagnas[j], rightX, baseY+float64(j*20), 0.0, 0.5)
				}
			}
			dc.SetRGB(0, 0, 0) // Reset to black
		}
	}

	// Note: Center text is not supported for North Indian charts
	// as there is no empty space in the middle like South Indian charts
	// The center is occupied by the inner square and dividing lines

	return encodePNG(dc.Image())
}
