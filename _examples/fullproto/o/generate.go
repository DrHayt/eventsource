package orders

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative ./events.proto
//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative ./commands.proto
//go:generate protoc --eventsource_out=. events.proto
//go:generate protoc --commands_out=. commands.proto
