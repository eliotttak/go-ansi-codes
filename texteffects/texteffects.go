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

package texteffects

import (
	"strconv"
	"strings"

	ansicodes "github.com/eliotttak/go-ansi-codes"
)

// TextEffectNumber is a text effect number that has to be parsed in to an
// [ansicodes.EscapeSequence] using [EscapeSequence]().
type TextEffectNumber int

// ToAnsi returns the [ansicodes.EscapeSequence] corresponding to ten.
func (ten TextEffectNumber) ToAnsi() ansicodes.EscapeSequence {
	return EscapeSequence(ten)
}

// EscapeSequence parses one or more [TextEffectNumber] into an [ansicodes.EscapeSequence]. It
// returns an empty string if there is no argument.
func EscapeSequence(n ...TextEffectNumber) ansicodes.EscapeSequence {
	if len(n) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteByte(033)
	b.WriteRune('[')
	first := true
	for _, e := range n {
		if !first {
			b.WriteRune(';')
		} else {
			first = false
		}
		b.WriteString(strconv.Itoa(int(e)))

	}
	b.WriteRune('m')
	return ansicodes.EscapeSequence(b.String())
}

// Reset is the ANSI text effect number that reset every text style or color.
const Reset TextEffectNumber = 0

// Theses constants are ANSI text effect numbers that enable a text style.
const (
	Bold            TextEffectNumber = 1
	Dim             TextEffectNumber = 2
	Italic          TextEffectNumber = 3
	Underline       TextEffectNumber = 4
	Blinking        TextEffectNumber = 5
	Inverse         TextEffectNumber = 7
	Invisible       TextEffectNumber = 8
	Strikethrough   TextEffectNumber = 9
	DoubleUnderline TextEffectNumber = 21
)

// Theses constants are ANSI text effect numbers that disable one or more text style(s).
const (
	ResetBoldAndDim    TextEffectNumber = 22
	ResetItalic        TextEffectNumber = 23
	ResetUnderline     TextEffectNumber = 24
	ResetBliking       TextEffectNumber = 25
	ResetInverse       TextEffectNumber = 27
	ResetInvisible     TextEffectNumber = 28
	ResetStrikethrough TextEffectNumber = 29
)
