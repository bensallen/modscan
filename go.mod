module github.com/bensallen/modscan

go 1.14

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/fanyang01/radix v0.0.0-20160415095728-e1747dd9eeac
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/u-root/u-root v7.0.0+incompatible
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68
)

replace github.com/u-root/u-root v7.0.0+incompatible => github.com/u-root/u-root v1.0.1-0.20201119150355-04f343dd1922
