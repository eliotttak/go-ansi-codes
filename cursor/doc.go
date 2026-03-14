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

// Package cursor defines ANSI escapes sequences that interacts with the terminal cursor.
//
// As in some others packages, the documentation says, for most of the constants and functions,
// "Foobar does something". However, the constants and the functions results have to be written to
// the standard output (for instance) to be executed, by inserting them in a string to be written,
// or directly using their methods [ansicodes.EscapeSequence.Run] and
// [ansicodes.EscapeSequence.Frun].
package cursor
