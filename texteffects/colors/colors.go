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

package colors

import (
	"github.com/eliotttak/go-ansi-codes/texteffects"
)

// DefaultF is the ANSI text effect number that reset the foreground color.
const DefaultF texteffects.TextEffectNumber = 39

// DefaultB is the ANSI text effect number that reset the background color.
const DefaultB texteffects.TextEffectNumber = 49

// Theses constants are 4-bits (16 colors) representations of basic, foreground colors.
const (
	BlackF4  texteffects.TextEffectNumber = 30
	RedF4    texteffects.TextEffectNumber = 31
	GreenF4  texteffects.TextEffectNumber = 32
	YellowF4 texteffects.TextEffectNumber = 33
	BlueF4   texteffects.TextEffectNumber = 34
	PurpleF4 texteffects.TextEffectNumber = 35
	CyanF4   texteffects.TextEffectNumber = 36
	WhiteF4  texteffects.TextEffectNumber = 37
)

// Theses constants are 4-bits (16 colors) representations of bright, foreground colors.
const (
	BrightBlackF4  texteffects.TextEffectNumber = 90
	BrightRedF4    texteffects.TextEffectNumber = 91
	BrightGreenF4  texteffects.TextEffectNumber = 92
	BrightYellowF4 texteffects.TextEffectNumber = 93
	BrightBlueF4   texteffects.TextEffectNumber = 94
	BrightPurpleF4 texteffects.TextEffectNumber = 95
	BrightCyanF4   texteffects.TextEffectNumber = 96
	BrightWhiteF4  texteffects.TextEffectNumber = 97
)

// Theses constants are 4-bits (16 colors) representations of basic, background colors.
const (
	BlackB4  texteffects.TextEffectNumber = 40
	RedB4    texteffects.TextEffectNumber = 41
	GreenB4  texteffects.TextEffectNumber = 42
	YellowB4 texteffects.TextEffectNumber = 43
	BlueB4   texteffects.TextEffectNumber = 44
	PurpleB4 texteffects.TextEffectNumber = 45
	CyanB4   texteffects.TextEffectNumber = 46
	WhiteB4  texteffects.TextEffectNumber = 47
)

// Theses constants are 4-bits (16 colors) representations of bright, background colors.
const (
	BrightBlackB4  texteffects.TextEffectNumber = 100
	BrightRedB4    texteffects.TextEffectNumber = 101
	BrightGreenB4  texteffects.TextEffectNumber = 102
	BrightYellowB4 texteffects.TextEffectNumber = 103
	BrightBlueB4   texteffects.TextEffectNumber = 104
	BrightPurpleB4 texteffects.TextEffectNumber = 105
	BrightCyanB4   texteffects.TextEffectNumber = 106
	BrightWhiteB4  texteffects.TextEffectNumber = 107
)

// ColorF8 returns a slice of the three values that defines a 8-bits (256 colors) foreground color.
// You can find the values to pass to ColorF8 in this image:
// https://user-images.githubusercontent.com/995050/47952855-ecb12480-df75-11e8-89d4-ac26c50e80b9.png
// (thanks to @ConnerWill https://www.github.com/ConnerWill). The colors are the same for background
// as for foreground.
func ColorF8(color texteffects.TextEffectNumber) []texteffects.TextEffectNumber {
	return []texteffects.TextEffectNumber{38, 5, color}
}

// ColorB8 returns a slice of the three values that defines a 8-bits (256 colors) background color.
// You can find the values to pass to ColorB8 in this image:
// https://user-images.githubusercontent.com/995050/47952855-ecb12480-df75-11e8-89d4-ac26c50e80b9.png
// (thanks to @ConnerWill https://www.github.com/ConnerWill). The colors are the same for background
// as for foreground.
func ColorB8(color texteffects.TextEffectNumber) []texteffects.TextEffectNumber {
	return []texteffects.TextEffectNumber{48, 5, color}
}

// TrueColorF returns a slice of the three values that defines a TrueColor (24-bits; 256³ = 16777216
// colors) foreground color. You can pass values from 0 to 255.
func TrueColorF(r, g, b texteffects.TextEffectNumber) []texteffects.TextEffectNumber {
	return []texteffects.TextEffectNumber{38, 2, r, g, b}
}

// TrueColorB returns a slice of the three values that defines a TrueColor (24-bits; 256³ = 16777216
// colors) background color. You can pass values from 0 to 255.
func TrueColorB(r, g, b texteffects.TextEffectNumber) []texteffects.TextEffectNumber {
	return []texteffects.TextEffectNumber{48, 2, r, g, b}
}
