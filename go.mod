module github.com/bensallen/modscan

go 1.14

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/fanyang01/radix v0.0.0-20160415095728-e1747dd9eeac
	github.com/klauspost/compress v1.10.11 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/u-root/u-root v7.0.0+incompatible
	github.com/ulikunitz/xz v0.5.8 // indirect
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f
)

replace github.com/u-root/u-root => github.com/bensallen/u-root v0.0.0-20200908023917-b3c793943d5b
