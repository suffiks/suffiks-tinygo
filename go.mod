module github.com/suffiks/suffiks-tinygo

go 1.20

require (
	github.com/dave/jennifer v1.6.1
	github.com/kubewarden/k8s-objects v1.27.0-kw1
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-openapi/strfmt v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
)

replace net => tinygo.org/x/drivers/net v0.8.0

replace github.com/go-openapi/strfmt => github.com/kubewarden/strfmt v0.1.0
