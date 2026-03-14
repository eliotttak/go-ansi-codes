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

package cursor

import (
	"fmt"

	ansicodes "github.com/eliotttak/go-ansi-codes"
)

// MakeVisible and MakeInvisible make the cursor visible or invisible.
const (
	MakeVisible   ansicodes.EscapeSequence = "\033[?25h"
	MakeInvisible ansicodes.EscapeSequence = "\033[?25l"
)

// GoToHome puts the cursor at (0;0).
const GoToHome ansicodes.EscapeSequence = "\033[H"

// RequestPosition asks the terminal for the current cursor position. It responds in the standard
// input in the form 'ESC l ; c R' (without the spaces), where ESC is the escape character ('\033'
// in octal), l is the line number and c is the column number (R is written as it).
//
// The terminal considers, in the response, that (1;1) is the upper left corner of the screen.
const RequestPosition ansicodes.EscapeSequence = "\033[6n"

// OneLineUpWithScrolling moves up the cursor of one row, without changing the column, scrolling if
// necessary.
const OneLineUpWithScrolling ansicodes.EscapeSequence = "\033M"

// SavePositionDec and RestorePositionDec respectively saves and restores the cursor position, in
// DEC mode.
const (
	SavePositionDec    ansicodes.EscapeSequence = "\0337"
	RestorePositionDec ansicodes.EscapeSequence = "\0338"
)

// SavePositionSco and RestorePositionSco respectively saves and restores the cursor position, in
// SCO mode.
const (
	SavePositionSco    ansicodes.EscapeSequence = "\033[s"
	RestorePositionSco ansicodes.EscapeSequence = "\033[u"
)

// GoToLineColumn moves the cursor to the specified line and column.
func GoToLineColumn(line, column int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%d;%dH", line, column))
}

// MoveUpNLines moves up the cursor of the specified number of lines.
func MoveUpNLines(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dA", n))
}

// MoveDownNLines moves down the cursor of the specified number of lines.
func MoveDownNLines(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dB", n))
}

// MoveRightNColumns moves right the cursor of the specified number of columns.
func MoveRightNColumns(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dC", n))
}

// MoveLeftNColumns moves left the cursor of the specified number of columns.
func MoveLeftNColumns(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dD", n))
}

// MoveDownNLinesAndGoStart moves down the cursor of the specified number of lines, and moves it to
// the start of the line.
func MoveDownNLinesAndGoStart(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dE", n))
}

// MoveUpNLinesAndGoStart moves up the cursor of the specified number of lines, and moves it to the
// start of the line.
func MoveUpNLinesAndGoStart(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dF", n))
}

// MoveToColumnN moves the cursor to the specified column.
func MoveToColumnN(n int) ansicodes.EscapeSequence {
	return ansicodes.EscapeSequence(fmt.Sprintf("\033[%dG", n))
}
