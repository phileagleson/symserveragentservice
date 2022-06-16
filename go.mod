module github.com/phileagleson/symserveragentservice

go 1.18

require (
	github.com/kardianos/service v1.2.1
	github.com/phileagleson/symserveragent v0.0.0-00010101000000-000000000000
)

require (
	github.com/phileagleson/ziputils v0.0.0-20220608180218-94096c355153 // indirect
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211 // indirect
)

replace github.com/phileagleson/symserveragent => ../symserveragent
