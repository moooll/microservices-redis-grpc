package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ScanInput() (companyName string, open bool, err error) {
	fmt.Println("input format: $company_name $open/close")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", false, err
	}

	input = strings.Replace(input, "\n", "", -1)
	i := strings.Split(input, "")
	if len(i) == 2 {
		companyName = i[0]
		if i[1] == "open" {
			return companyName, true, nil
		} else if i[1] == "close" {
			return companyName, false, nil
		} else {
			return companyName, false, errors.New("wrong input")
		}
	}
	
	return "", false, nil
}
