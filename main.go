package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	ticTacToe()
}

func ticTacToe() {

	rand.Seed(time.Now().UnixNano())

	t := [3][3]string{}
	for {
		t = userTurn(t)

		nicePrint(t)

		if checkWin("x", t) {
			fmt.Println("user wins")
			break
		}
		if checkGameOver(t) {
			fmt.Println("game over: tie")
			break
		}
		t = computerTurn(t)

		nicePrint(t)

		if checkWin("o", t) {
			fmt.Println("computer wins")
			break
		}
		if checkGameOver(t) {
			fmt.Println("game over: tie")
			break
		}

	}
}

func checkWin(symbol string, t [3][3]string) bool {

	var a, b, c string

	// check columns
	for col := 0; col <= 2; col++ {
		a, b, c = t[0][col], t[1][col], t[2][col]
		if a == symbol && b == symbol && c == symbol {
			return true
		}
	}

	// check rows
	for row := 0; row <= 2; row++ {
		a, b, c = t[row][0], t[row][1], t[row][2]
		if a == symbol && b == symbol && c == symbol {
			return true
		}
	}

	// check diagonals
	a, b, c = t[0][0], t[1][1], t[2][2]
	if a == symbol && b == symbol && c == symbol {
		return true
	}

	a, b, c = t[0][2], t[1][1], t[2][0]
	if a == symbol && b == symbol && c == symbol {
		return true
	}
	return false
}

func checkAlmostWin(symbol string, t [3][3]string) (bool, int, int) {

	// check columns
	for col := 0; col <= 2; col++ {

		sum := 0
		emptyRow := -1
		emptyCol := -1

		for row := 0; row <= 2; row++ {
			cell := t[row][col]

			if cell == symbol {
				sum = sum + 1
			} else if cell == "" {
				emptyRow = row
				emptyCol = col
			}
		}
		if sum == 2 {
			found := emptyRow != -1
			return found, emptyRow, emptyCol
		}
	}

	// check rows

	for row := 0; row <= 2; row++ {

		sum := 0
		emptyCol := -1
		emptyRow := -1
		for col := 0; col <= 2; col++ {
			cell := t[row][col]

			if cell == symbol {
				sum = sum + 1
			} else if cell == "" {
				emptyCol = col
				emptyRow = row
			}
		}
		if sum == 2 {
			found := emptyCol != -1
			return found, emptyRow, emptyCol
		}
	}

	// check 1st diagonal
	sum := 0
	emptyIndex := -1
	for row := 0; row <= 2; row++ {
		cell := t[row][row]
		if cell == symbol {
			sum = sum + 1
		} else if cell == "" {
			emptyIndex = row

		}

	}
	if sum == 2 {
		found := emptyIndex != -1
		return found, emptyIndex, emptyIndex
	}

	// check 2nd diagonal
	sum = 0
	emptyIndex = -1

	for row := 0; row <= 2; row++ {
		col := 2 - row
		cell := t[row][col]
		if cell == symbol {
			sum = sum + 1
		} else if cell == "" {
			emptyIndex = row
		}
	}
	if sum == 2 {
		found := emptyIndex != -1
		return found, emptyIndex, emptyIndex
	}

	return false, -1, -1
}

func tryCorner(t [3][3]string) (bool, [3][3]string) {
	if checkEmptyCell(0, 0, t) {
		t = writeO(0, 0, t)
		return true, t
	}
	if checkEmptyCell(0, 2, t) {
		t = writeO(0, 2, t)
		return true, t
	}
	if checkEmptyCell(2, 2, t) {
		t = writeO(2, 2, t)
		return true, t
	}
	if checkEmptyCell(2, 0, t) {
		t = writeO(2, 0, t)
		return true, t
	}
	return false, t
}

func computerTurn(t [3][3]string) [3][3]string {
	found, x, y := checkAlmostWin("o", t)
	if found {
		t = writeO(x, y, t)
		return t
	}

	found, x, y = checkAlmostWin("x", t)
	if found {
		t = writeO(x, y, t)
		return t
	}

	ok, t := tryCorner(t)
	if ok {
		return t
	}

	for {
		x := rand.Intn(3)
		y := rand.Intn(3)
		fmt.Println("random x and y:", x+1, y+1)
		cellIsEmpty := checkEmptyCell(x, y, t)
		if cellIsEmpty {

			t = writeO(x, y, t)
			break
		} else {
			fmt.Println("cell is taken", t[x][y])
		}

	}
	return t
}

func userTurn(t [3][3]string) [3][3]string {

	var x, y int

	for {
		input := readString()
		k := strings.SplitN(input, ",", 2)

		x = stringToInt(k[0]) - 1
		y = stringToInt(k[1]) - 1
		if x >= 0 && x <= 2 && y >= 0 && y <= 2 {
			break
		}
		fmt.Println("number not fitting in the game")
	}

	if checkEmptyCell(x, y, t) {
		fmt.Println("OK")
		t[x][y] = "x"
	} else {
		fmt.Println("cell is not empty", t[x][y])
	}
	return t
}

func printMatrix(t [3][3]string) {
	for i := 0; i <= 2; i++ {
		lineArray := []string{}
		for j := 0; j <= 2; j++ {
			lineArray = append(lineArray, t[i][j])
		}

		line := strings.Join(lineArray, " | ")
		fmt.Println(line)
		fmt.Println("_________")

	}
}

func nicePrint(matrix [3][3]string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 1, 3, ' ', tabwriter.Debug|tabwriter.AlignRight)

	for i := 0; i <= 2; i++ {
		lineArray := []string{}
		for j := 0; j <= 2; j++ {
			lineArray = append(lineArray, matrix[i][j])
		}

		line := strings.Join(lineArray, "\t")
		fmt.Fprintln(w, line)
		//fmt.Fprintln(w, "---\t---\t---\t")
	}
	fmt.Fprintln(w, "")
	w.Flush()
}

func checkEmptyCell(x int, y int, matrix [3][3]string) bool {
	if matrix[x][y] == "" {
		return true

	} else {
		return false
	}
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	return text
}
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0
	}

	return i

}

func writeO(x int, y int, t [3][3]string) [3][3]string {

	if checkEmptyCell(x, y, t) {
		t[x][y] = "o"

	} else {
		fmt.Println("position already taken", x, y)
		debug.PrintStack()
	}
	return t
}

func checkGameOver(t [3][3]string) bool {
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			if checkEmptyCell(x, y, t) {
				return false
			}
		}
	}
	return true
}
