package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Id          string
	CompanyName string
	Open        bool
	Err         error
}

func ScanInput() (in Input) {
	fmt.Println("input format: $company_name $open/close")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return Input{"", "", false, err}
	}

	input = strings.Replace(input, "\n", "", -1)
	i := strings.Split(input, " ")
	if len(i) == 2 {
		companyName := i[0]
		if i[1] == "open" {
			return Input{"", companyName, true, nil}
		} else {
			e := errors.New("wrong input")
			return Input{"", companyName, false, e}
		}
	} else if len(i) == 3 && i[1] == "close" {
		companyName := i[0]
		id := i[2]
		return Input{id, companyName, false, nil}
	}

	return Input{"", "", false, nil}
}

// func ScanInput() Input {
// 	return Input{
// 		CompanyName: "apple",
// 		Err:         nil,
// 		Open:        false,
// 	}
// }
