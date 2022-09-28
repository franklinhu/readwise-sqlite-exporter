module github.com/franklinhu/readwise-sqlite-exporter

go 1.19

require (
	github.com/ethanholz/readwise-go v0.1.0
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/stretchr/testify v1.8.0
)

// Remove after closing https://github.com/ethanholz/readwise-go/pull/1
replace github.com/ethanholz/readwise-go v0.1.0 => github.com/franklinhu/readwise-go v0.0.0-20220928042856-a4ce1fd995bf

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
