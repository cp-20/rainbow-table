package main

const (
	lowercaseCharset = "abcdefghijklmnopqrstuvwxyz"
	uppercaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitsCharset    = "0123456789"
	specialCharset   = "!@#$%^&*()_+-=[]{}|;:,<.>/? "
)

const (
	chainLength       = 5000
	numberOfChains    = 100000
	charset           = lowercaseCharset
	rainbowTableFile  = "rainbow_table_sha256.txt"
	goroutineNum      = 16
	passwordMinLength = 8
	passwordMaxLength = 14
)
