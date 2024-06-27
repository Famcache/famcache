package command

import "fmt"

func SetCommand(id, key, value string, ttl *uint64) string {
	if ttl != nil {
		return fmt.Sprintf("%s SET %s %s %d", id, key, value, *ttl)
	}
	return fmt.Sprintf("%s SET %s %s", id, key, value)
}

func GetCommand(id, key string) string {
	return fmt.Sprintf("%s GET %s", id, key)
}

func DeleteCommand(id, key string) string {
	return fmt.Sprintf("%s DELETE %s", id, key)
}

func PublishCommand(id, topic, data string) string {
	return fmt.Sprintf("%s PUBLISH %s %s", id, topic, data)
}

func SubscribeCommand(id, topic string) string {
	return fmt.Sprintf("%s SUBSCRIBE %s", id, topic)
}

func UnsubscribeCommand(id, topic string) string {
	return fmt.Sprintf("%s UNSUBSCRIBE %s", id, topic)
}

func JobExecuteCommand(jobId string) string {
	return fmt.Sprintf("JOB_EXECUTE %s\n", jobId)
}
