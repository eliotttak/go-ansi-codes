// Copyright © 2026 Eliott TAKVORIAN
//
// This file is a part of go-ansi-colors.
//
// go-ansi-colors is free software: you can redistribute it and/or modify it under the terms of the
// GNU Lesser General Public License as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without
// even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License along with this program.
// If not, see <https://www.gnu.org/licenses/>

package screen

import (
	"strconv"

	ansicodes "github.com/eliotttak/go-ansi-codes"
)

// OptionCode defines a ANSI screen mode number that can be parsed into
// an [ansicodes.EscapeSequence] with its methods [OptionCode.Set]() and
// [OptionCode.Unset]().
type OptionCode int8

// MonochromeText40x25 sets or unsets the screen to display monochrome
// text with a size of 40 column and 25 lines.
const MonochromeText40x25 OptionCode = 0

// ColoredText40x25 sets or unsets the screen to display colored
// text with a size of 40 column and 25 lines.
const ColoredText40x25 OptionCode = 1

// MonochromeText80x25 sets or unsets the screen to display monochrome
// text with a size of 80 column and 25 lines.
const MonochromeText80x25 OptionCode = 2

// ColoredText80x25 sets or unsets the screen to display colored
// text with a size of 80 column and 25 lines.
const ColoredText80x25 OptionCode = 3

// FourColorsGraphics320x200 sets or unsets the screen to display
// 4-colors graphics with a size of 320 columns and 200 lines.
const FourColorsGraphics320x200 OptionCode = 4

// MonochromeGraphics320x200 sets or unsets the screen to display
// monochrome graphics with a size of 320 columns and
// 200 lines.
const MonochromeGraphics320x200 OptionCode = 5

// MonochromeGraphics640x200 sets or unsets the screen to display
// monochrome graphics with a size of 640 columns and
// 200 lines.
const MonochromeGraphics640x200 OptionCode = 6

// LineWrapping enables or disables the line wrapping.
const LineWrapping OptionCode = 7

// FourColorsGraphics320x200 sets or unsets the screen to display
// colored graphics with a size of 320 columns and 200 lines.
const ColoredGraphics320x200 OptionCode = 13

// SixteenColorsGraphics640x200 sets or unsets the screen to display
// 16-colors graphics with a size of 640 columns and 200 lines.
const SixteenColorsGraphics640x200 OptionCode = 14

// MonochromeGraphics640x350 sets or unsets the screen to display
// monochrome graphics with a size of 640 columns and
// 350 lines.
const MonochromeGraphics640x350 OptionCode = 15

// SixteenColorsGraphics640x350 sets or unsets the screen to display
// 16-colors graphics with a size of 640 columns and 350 lines.
const SixteenColorsGraphics640x350 OptionCode = 16

// MonochromeGraphics640x480 sets or unsets the screen to display
// monochrome graphics with a size of 640 columns and
// 480 lines.
const MonochromeGraphics640x480 OptionCode = 17

// SixteenColorsGraphics640x480 sets or unsets the screen to display
// 16-colors graphics with a size of 640 columns and 480 lines.
const SixteenColorsGraphics640x480 OptionCode = 18

// TwoHundredFiftySixColorsGraphics320x200 sets or unsets the screen
// to display 256-colors graphics with a size of 320 columns and 200 lines.
const TwoHundredFiftySixColorsGraphics320x200 OptionCode = 19

// Set creates the [ansicodes.EscapeSequence] that enables the [OptionCode]
// on which Set is run.
func (c OptionCode) Set() ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence("\033[=" + strconv.Itoa(int(c)) + "h")
}

// Unset creates the [ansicodes.EscapeSequence] that disables the [OptionCode]
// on which Unset is run.
func (c OptionCode) Unset() ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence("\033[=" + strconv.Itoa(int(c)) + "l")
}

// SaveScreen saves the current screen to restore it later with [RestoreScreen].
const SaveScreen ansicodes.EscapeSequence = "\033[?47h"

// RestoreScreen restores the screen previously saved by [SaveScreen].
const RestoreScreen ansicodes.EscapeSequence = "\033[?47l"
