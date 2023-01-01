module github.com/dundunlabs/bunapp/example

go 1.19

replace github.com/dundunlabs/bunapp => ../

require (
	github.com/dundunlabs/bunapp v0.0.0-00010101000000-000000000000
	github.com/uptrace/bun v1.1.9
	github.com/uptrace/bun/dialect/pgdialect v1.1.9
	github.com/uptrace/bun/driver/pgdriver v1.1.9
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/uptrace/bunrouter/extra/reqlog v1.0.19 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/uptrace/bunrouter v1.0.19
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	mellium.im/sasl v0.3.0 // indirect
)
