package models

import (
	"os"
	"os/exec"
	"runtime"
)

func CallClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//var clear map[string]func() //create a map for storing clear funcs
//
//func init() {
//	clear = make(map[string]func()) //Initialize it
//	clear["linux"] = func() {
//		cmd := exec.Command("clear") //Linux example, its tested
//		cmd.Stdout = os.Stdout
//		cmd.Run()
//	}
//	clear["windows"] = func() {
//		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
//		cmd.Stdout = os.Stdout
//		cmd.Run()
//	}
//}
//
//func CallClear() {
//	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
//	if ok {                          //if we defined a clear func for that platform:
//		value() //we execute it
//	} else { //unsupported platform
//		panic("Your platform is unsupported! I can't clear terminal screen :(")
//	}
//}

//func main() {
//	time.Sleep(2 * time.Second)
//	CallClear()
//}

//fmt.Print("\033[H\033[4J")

//func CallClear() {
//	h, l := 8, 8
//	board := make([]uint8, h*l)
//	for i := 0; i < 120; i++ {
//		clear()
//		cellFlip(rand.Intn(len(board)), board)
//		fmt.Println(plotted(l, board))
//		time.Sleep(400 * time.Millisecond)
//	}
//}
//
//func clear() {
//	cmd := exec.Command("clear")
//	cmd.Stdout = os.Stdout
//	cmd.Run()
//}
//
//func cellFlip(i int, b []uint8) {
//	if b[i] >= 3 {
//		return
//	}
//	b[i]++
//}
//
//func plotted(l int, b []uint8) string {
//	var s strings.Builder
//	s.Grow(len(b) + l)
//	for i, cell := range b {
//		s.WriteRune(toRune(cell))
//		if i%l != l-1 {
//			continue
//		}
//		s.WriteRune('\n')
//	}
//	return s.String()
//}
//
//func toRune(cell uint8) rune {
//	var c rune
//	switch cell {
//	case 0:
//		c = 0x53E3
//	case 1:
//		c = 0x4E00
//	case 2:
//		c = 0x4E8C
//	case 3:
//		c = 0x4E09
//	default:
//		c = 'x'
//	}
//	return c
//}
