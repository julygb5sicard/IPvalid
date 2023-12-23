package main

import (
"fmt"
"net"
"os"
)

func main() {
if len(os.Args) != 2 {
fmt.Println("Usage: go run script.go <IP address>")
os.Exit(1)
}

ipAddress := os.Args[1]
if isValidIP(ipAddress) {
fmt.Println("Entered IP address", ipAddress, "valid.")
} else {
fmt.Println("The entered IP address", ipAddress, "invalid.")
}
}

func isValidIP(ipAddress string) bool {
parsedIP := net.ParseIP(ipAddress)
return parsedIP != nil
}
