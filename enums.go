// Package go-liblzma is a wrapper for liblzma and XZ file format.
package xzgo

const DefaultBufsize = 32768

type Action uint

const (
	// Continue coding.
	Run Action = iota
	// Make all the input available at output.
	SyncFlush
	// Finish encoding of the current Block.
	FullFlush
	// Finish the coding operation.
	Finish
)

type Errno uint

var _ error = Errno(0)

const (
	// Operation completed successfully.
	Ok Errno = iota
	// End of stream was reached.
	StreamEnd
	// Input stream has no integrity check.
	NoCheck
	// Cannot calculate the integrity check.
	UnsupportedCheck
	// Integrity check type is now available.
	GetCheck
	// Cannot allocate memory.
	MemError
	// Memory usage limit was reached.
	MemlimitError
	// File format not recognized.
	FormatError
	// Invalid or unsupported options.
	OptionsError
	// Data is corrupt.
	DataError
	// No progress is possible.
	BufError
	// Programming error.
	ProgError
)

var errorMsg = [...]string{
	"Operation completed successfully",
	"End of stream was reached",
	"Input stream has no integrity check",
	"Cannot calculate the integrity check",
	"Integrity check type is now available",
	"Cannot allocate memory",
	"Memory usage limit was reached",
	"File format not recognized",
	"Invalid or unsupported options",
	"Data is corrupt",
	"No progress is possible",
	"Programming error",
}

func (e Errno) Error() string {
	return errorMsg[e]
}

type Checksum uint

const (
	CheckNone   Checksum = 0
	CheckCRC32  Checksum = 1
	CheckCRC64  Checksum = 4
	CheckSHA256 Checksum = 10
)

type Preset uint32

const (
	Level0 Preset = iota
	Level1
	Level2
	Level3
	Level4
	Level5
	Level6
	Level7
	Level8
	Level9
)

const (
	// Default compression preset.
	LevelDefault Preset = Level6
	// Extreme compression preset. To be OR'ed with another preset.
	LevelExtreme Preset = 1 << 31
	// Mask for preset level. To AND with a Preset to extract the level.
	LevelMask Preset = 0x1f
)
