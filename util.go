package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// err: Actual error
// customError: User defined error message for easy debugging
// exitProgram: Specify if the program has to terminate or not (if true, the program will gracefully exit with os.Exit(1))
func HandleError(err error, customError string, exitProgram bool) bool {
	if err != nil {
		color.Red("Error: %s | Custom message: %s", err, customError)
		if exitProgram {
			os.Exit(1)
		}
		return true
	}

	return false
}

// converts a string of characters into a integer array, you can specify the seperator and you can just pass empty string if you aren't aware
func StringToIntArr(inp string, separator string) (resp []int, err error) {
	for _, v := range strings.Split(inp, separator) {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}

		resp = append(resp, i)
	}

	return
}

// return comma seperated string values for the given int array
func IntArrToStr(inp []int) (res string) {
	for i, v := range inp {
		if i == 0 {
			res = fmt.Sprintf("%d", v)
			continue
		}

		res = fmt.Sprintf("%s,%d", res, v)
	}

	return
}

func Contains(data []any, toFind any) bool {
	for _, v := range data {
		if v == toFind {
			return true
		}
	}

	return false
}
