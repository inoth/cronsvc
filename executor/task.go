package executor

func ReceiverTask(msg []byte) {
	ecr.receiver <- msg
}
