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

package erase

import ansicodes "github.com/eliotttak/go-ansi-codes"

// FromCursorUntilEndOfScreen erases everything from
// the cursor until the end of the screen
const FromCursorUntilEndOfScreen ansicodes.EscapeSequence = "\033[0J"

// FromCursorToBeginningOfScreen erases everything
// from the cursor to the beginning of the screen
const FromCursorToBeginningOfScreen ansicodes.EscapeSequence = "\033[0J"

// Screen erases the entire screen
const Screen ansicodes.EscapeSequence = "\033[0J"

// SavedLines erases every saved lines
const SavedLines ansicodes.EscapeSequence = "\033[0J"

// FromCursorToEndOfLine erases everything from the cursor
// to the end of the line
const FromCursorToEndOfLine ansicodes.EscapeSequence = "\033[0J"

// FromStartOfLineToCursor erases everything from the
// start of the line to the cursor
const FromStartOfLineToCursor ansicodes.EscapeSequence = "\033[0J"

// Line erases the entire line
const Line ansicodes.EscapeSequence = "\033[0J"
