package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanInput() (string, error) {
	fmt.Println("input format: $company_name $open/close")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.Replace(input, "\n", "", -1)
	return input, nil
}
