package command

type CommandType = string

const (
	// Store commands
	CommandGet    CommandType = "GET"
	CommandSet    CommandType = "SET"
	CommandDelete CommandType = "DELETE"

	// Pubsub commands
	CommandPublish     CommandType = "PUBLISH"
	CommandSubscribe   CommandType = "SUBSCRIBE"
	CommandUnsubscribe CommandType = "UNSUBSCRIBE"

	// Job commands
	CommandJobRegister CommandType = "JOB_REGISTER"
	CommandJobStatus   CommandType = "JOB_STATUS"
	CommandJobCancel   CommandType = "JOB_CANCEL"
	CommandJobExecute  CommandType = "JOB_EXECUTE"
)
