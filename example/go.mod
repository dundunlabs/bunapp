module github.com/dundunlabs/bunapp/example

go 1.19

replace (
	github.com/dundunlabs/bunapp => ../
	github.com/dundunlabs/bunapp/x/errorhdr => ../x/errorhdr
	github.com/dundunlabs/bunapp/x/panichdr => ../x/panichdr
)

require (
	github.com/dundunlabs/bunapp v0.0.0-00010101000000-000000000000
	github.com/dundunlabs/bunapp/x/errorhdr v0.0.0-00010101000000-000000000000
	github.com/dundunlabs/bunapp/x/panichdr v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
	github.com/uptrace/bun v1.1.9
	github.com/uptrace/bun/dialect/pgdialect v1.1.9
	github.com/uptrace/bun/driver/pgdriver v1.1.9
	github.com/uptrace/bunrouter v1.0.19
	github.com/uptrace/bunrouter/extra/reqlog v1.0.19
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.opentelemetry.io/otel v1.9.0 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	mellium.im/sasl v0.3.0 // indirect
)
