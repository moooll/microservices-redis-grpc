package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	CompanyName string
	Open bool
	Err error
}

func ScanInput() (in Input) {
	fmt.Println("input format: $company_name $open/close")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return Input{"", false, err}
	}

	input = strings.Replace(input, "\n", "", -1)
	i := strings.Split(input, "")
	if len(i) == 2 {
		companyName := i[0]
		if i[1] == "open" {
			return Input{companyName, true, nil}
		} else if i[1] == "close" {
			return Input{companyName, false, nil}
		} else {
			e := errors.New("wrong input")
			return Input{companyName, false, e}
		}
	}
	
	return Input{"", false, nil}
}
