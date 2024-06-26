package command

import (
	"famcache/domain/command"
	"strconv"
	"strings"
)

type JobCommand struct {
	id         string
	command    command.CommandType
	isPeriodic bool
	delay      uint64
	jobId      string
}

func (c *JobCommand) ID() string {
	return c.id
}

func (c *JobCommand) Type() command.CommandType {
	return c.command
}

func (c *JobCommand) IsPeriodic() bool {
	return c.isPeriodic
}

func (c *JobCommand) Delay() uint64 {
	return c.delay
}

func (c *JobCommand) JobId() string {
	return c.jobId
}

// uuid JOB_REGISTER <delay> <is_periodic>
func newRegisterJobCommand(parts []string) command.JobCommand {
	if len(parts) < 3 {
		return nil
	}

	delayStr := parts[2]

	delay := uint64(0)

	if delayStr != "" {
		err := error(nil)
		delay, err = strconv.ParseUint(delayStr, 10, 64)

		if err != nil {
			return nil
		}
	}

	isPeriodic := false

	if len(parts) > 3 {
		isPeriodic = parts[3] == "true"
	}

	return &JobCommand{
		id:         parts[0],
		command:    command.CommandJobRegister,
		delay:      delay,
		isPeriodic: isPeriodic,
		jobId:      "",
	}
}

// uuid JOB_CANCEL <job_id>
func newCancelJobCommand(parts []string) command.JobCommand {
	if len(parts) < 2 {
		return nil
	}

	return &JobCommand{
		id:         parts[0],
		command:    command.CommandJobCancel,
		delay:      0,
		isPeriodic: false,
		jobId:      parts[1],
	}
}

func NewJobCommand(commandType command.CommandType, query string) command.JobCommand {
	parts := strings.Fields(strings.TrimSpace(query))

	switch commandType {
	case command.CommandJobRegister:
		return newRegisterJobCommand(parts)
	case command.CommandJobCancel:
		return newCancelJobCommand(parts)
	default:
		return nil
	}
}
