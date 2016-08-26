package termcol

type Color string

const (
	_Black   Color = "\x1b[0;30"
	_Red     Color = "\x1b[0;31"
	_Green   Color = "\x1b[0;32"
	_Yellow  Color = "\x1b[0;33"
	_Blue    Color = "\x1b[0;34"
	_Magenta Color = "\x1b[0;35"
	//	light blue
	_Cyan  Color = "\x1b[0;36"
	_White Color = "\x1b[0;37"

	//	Default
	_Default Color = "\x1b[0;39"

	_Reset Color = "\x1b[0m"
)

const (
	Black   Color = "\x1b[0;30m"
	Red     Color = "\x1b[0;31m"
	Green   Color = "\x1b[0;32m"
	Yellow  Color = "\x1b[0;33m"
	Blue    Color = "\x1b[0;34m"
	Magenta Color = "\x1b[0;35m"
	//	light blue
	Cyan  Color = "\x1b[0;36m"
	White Color = "\x1b[0;37m"

	//	Default
	Default Color = "\x1b[0;39m"

	Reset Color = "\x1b[0m"
)

var Normal = struct {
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default Color
}{
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default,
}

var Bright = struct {
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default Color
}{
	_Black + ";1m",
	_Red + ";1m",
	_Green + ";1m",
	_Yellow + ";1m",
	_Blue + ";1m",
	_Magenta + ";1m",
	_Cyan + ";1m",
	_White + ";1m",
	_Default + ";1m",
}
