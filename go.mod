module github.com/JZiegener/notes-app

go 1.17

require github.com/JZiegener/notes-app/commands v0.0.0

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/JZiegener/notes-app/commands => ./commands
