package main

import "fmt"

type DivError struct {
	dive int
	divr int
}

func (de *DivError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dive : %d
	divr : 0
	`
	return fmt.Sprintf(strFormat, de.dive)
}

func Divide(varDivdee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dDate := DivError{
			dive: varDivdee,
			divr: varDivider,
		}
		errorMsg = dDate.Error()
		return
	} else {
		return varDivdee / varDivider, ""
	}
}

func main() {
	_, msg := Divide(100, 0)
	fmt.Println(msg)
}
