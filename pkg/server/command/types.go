package command

type CommandType = string

const (
	CommandGet    CommandType = "GET"
	CommandSet    CommandType = "SET"
	CommandDelete CommandType = "DELETE"

	CommandPublish   CommandType = "PUBLISH"
	CommandSubscribe CommandType = "SUBSCRIBE"
)
