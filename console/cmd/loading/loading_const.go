package loading

var PROGRESS = []string{
	"╍╍╍╍╍╍╍╍╍╍  ",
	"▰╍╍╍╍╍╍╍╍╍  ",
	"▰▰╍╍╍╍╍╍╍╍  ",
	"▰▰▰▰╍╍╍╍╍╍  ",
	"▰▰▰▰▰╍╍╍╍╍  ",
	"▰▰▰▰▰▰╍╍╍╍  ",
	"▰▰▰▰▰▰▰╍╍╍  ",
	"▰▰▰▰▰▰▰▰╍╍  ",
	"▰▰▰▰▰▰▰▰▰╍  ",
	"▰▰▰▰▰▰▰▰▰▰  "}

var CUBES = []string{
	"□□□□□□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■□□□□□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■□□□□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■□□□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■□□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■□□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■□□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■□□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■□□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■□□□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■□□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■□□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■□□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■□□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■□□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■□□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■□□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■□□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■□□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■□□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■□□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■□□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■□□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■□□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■□□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■■□□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■■■□□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■■■■□□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■■■■■□ ",
	"■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ ",
}

var BAR = []string{
	"▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"███████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"████████████████████████████▒▒▒▒▒▒▒▒▒▒▒▒ ",
	"█████████████████████████████▒▒▒▒▒▒▒▒▒▒▒ ",
	"██████████████████████████████▒▒▒▒▒▒▒▒▒▒ ",
	"███████████████████████████████▒▒▒▒▒▒▒▒▒ ",
	"████████████████████████████████▒▒▒▒▒▒▒▒ ",
	"█████████████████████████████████▒▒▒▒▒▒▒ ",
	"██████████████████████████████████▒▒▒▒▒▒ ",
	"███████████████████████████████████▒▒▒▒▒ ",
	"█████████████████████████████████████▒▒▒ ",
	"██████████████████████████████████████▒▒ ",
	"███████████████████████████████████████▒ ",
	"████████████████████████████████████████ ",
}