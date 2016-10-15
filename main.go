package main

import (
	"strconv"
	"bytes"
	
	"github.com/andlabs/ui"
)

func main() {
	createMainWindow()
}

func closeMainWindow(*ui.Window) bool {
	ui.Quit()
	return true
}

func createMainWindow() {
	err := ui.Main(func() {
		window := ui.NewWindow("Calculator", 150, 300, false)
		box := ui.NewVerticalBox()
		display := ui.NewLabel("")
		box.Append(display, false)

		row1 := ui.NewHorizontalBox()
		box.Append(row1, false)
		row1.Append(createInputButton("1", display), true)
		row1.Append(createInputButton("2", display), false)
		row1.Append(createInputButton("3", display), false)
		row1.Append(createClearButton(display), true)

		row2 := ui.NewHorizontalBox()
		box.Append(row2, false)
		row2.Append(createInputButton("4", display), true)
		row2.Append(createInputButton("5", display), false)
		row2.Append(createInputButton("6", display), false)
		row2.Append(createInputButton("+", display), true)

		row3 := ui.NewHorizontalBox()
		box.Append(row3, false)
		row3.Append(createInputButton("7", display), true)
		row3.Append(createInputButton("8", display), false)
		row3.Append(createInputButton("9", display), false)
		row3.Append(createInputButton("-", display), true)

		row4 := ui.NewHorizontalBox()
		box.Append(row4, false)
		row4.Append(createInputButton("0", display), true)
		row4.Append(createCalcButton(display), false)

		
		window.SetChild(box)
		window.OnClosing(closeMainWindow)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func createInputButton(n string, display *ui.Label) ui.Control {
	num1 := ui.NewButton(n)
	num1.OnClicked(func(*ui.Button) {
		display.SetText(display.Text() + n)
	})
	return num1
}
	
func createClearButton(display *ui.Label) ui.Control {
	clear := ui.NewButton("C")
	clear.OnClicked(func(*ui.Button) {
		display.SetText("")
	})
	return clear
}

func createCalcButton(display *ui.Label) ui.Control {
	calc := ui.NewButton("=")
	calc.OnClicked(func(*ui.Button) {
		expression := display.Text()
		result := calculate(expression)
		display.SetText(result)
	})
	return calc
}

const (
	PLUS = iota
	MINUS = iota
)

func calculate(expr string) string {
	result := 0
	operation := PLUS
	buffer := bytes.NewBufferString("")
	for _, c := range expr {
		if c == '+' || c == '-' {
			num, _ := strconv.Atoi(buffer.String())
			result = operate(result, operation, num)
			buffer.Truncate(0)
			if c == '+' {
				operation = PLUS
			} else {
				operation = MINUS
			}
		} else {
			buffer.WriteRune(c)
		}	
	}
	num, _ := strconv.Atoi(buffer.String())
	result = operate(result, operation, num)
	return strconv.Itoa(result)
}

func operate(result int, operation int, value int) int {
	if operation == PLUS {
		result = result + value
	} else {
		result = result - value
	}
	return result
}
	
