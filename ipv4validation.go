package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusiv
*/
func Is_valid_ip(ip string) bool {
	split := strings.Split(ip, ".")
	if len(split) != 4 {
		return false
	}

	for _, x := range split {
		if x[0] == '0' && len(x) != 1 {
			return false
		}
		if value, err := strconv.ParseInt(x, 10, 64); err == nil {
			if value < 0 || value > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func main() {
	if Is_valid_ip("12.34.56 .1") {
		fmt.Printf("fuap")
	}
}
