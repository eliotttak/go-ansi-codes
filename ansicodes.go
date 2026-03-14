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

package ansicodes

import (
	"io"
	"os"
)

// EscapeSequence is based on (but is not an alias for) string.
// It contains an ANSI escape sequence.
type EscapeSequence string

// Run writes the escape sequence to the standard output.
func (es EscapeSequence) Run() (int, error) {
	return os.Stdout.WriteString(string(es))
}

// Frun writes the escape sequence to the specified writer, which
// can be a file, the standard error, a [strings.Builder], etc.
func (es EscapeSequence) Frun(w io.Writer) (int, error) {
	return w.Write([]byte(es))
}
