module my-app

go 1.25.5

require (
	github.com/google/uuid v1.6.0
	my-local-lib v0.0.0
)

replace my-local-lib => ../local-lib
