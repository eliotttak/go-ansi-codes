# `go-ansi-codes`

`go-ansi-codes` is a Go package that implements most of ANSI escape sequences.

## Supported escape sequences
- Text effects:
	+ Bold
	+ Dim
	+ Italic
	+ Simple and double underline
	+ Blinking
	+ Inverse (invert foreground and background)
	+ Invisible
	+ Strikethrough
- Colors:
	+ A set of 16 colors (4 bits) for foreground and background, with regular and bright modes:
		* Black
		* Red
		* Green
		* Yellow
		* Blue
		* Purple
		* Cyan
		* White
	+ Support for 8-bits colors, or 256 colors. [The image below](#8-bits-color-codes) defines the
	  8-bits colors codes.
	+ Support for TrueType colors (24-bits, or 256³ = 16777216 colors).
- Cursor management:
	+ Visibility
	+ Position
- Erasing text on screen
- Screen modes:
	+ Monochrome, text, 40x25
	+ Colored, text, 40x25
	+ Monochrome, text, 80x25
	+ Colored, text, 80x25
	+ 4-colors, graphics, 320x200
	+ Monochrome, graphics, 320x200
	+ Monochrome, graphics, 640x200
	+ Colored, graphics, 320x200
	+ 16-colors, graphics, 640x200
	+ Monochrome, graphics, 640x350
	+ 16-colors, graphics, 640x350
	+ Monochrome, graphics, 640x480
	+ 16-colors, graphics, 640x480
	+ 256-colors, graphics, 320x200
	+ Enabling and disabling line wrapping
	> [!NOTE]
	> 
	> The commands for the screen mode are obsolete, and should no longer be used on recent terminal 
	emulators.
	
### 8-bits color codes

[![Click here to see the image][8-bits-codes]][8-bits-codes]

Thanks to [@ConnerWill](https://gist.github.com/ConnerWill)

## Installing
```bash
go get github.com/eliotttak/go-ansi-codes
```

## Documentation

Find the `go-ansi-codes` documentation here: [Go Packages Documentation]

## Licensing

This library is released under the GNU Lesser General Public License version 3 (LGPLv3).

You can find the entire license content in [LICENSE.md](./LICENSE.md).

---

 [8-bits-codes]: https://user-images.githubusercontent.com/995050/47952855-ecb12480-df75-11e8-89d4-ac26c50e80b9.png "Click to open the image"
 [Go Packages Documentation]: https://pkg.go.dev/github.com/eliotttak/go-ansi-codes
