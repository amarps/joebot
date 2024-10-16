package server

import (
	"net"

	"github.com/amarps/joebot/src/models"
	"github.com/amarps/joebot/src/task"
	"github.com/amarps/joebot/src/utils"
	"github.com/pkg/errors"
)

type ClientInfoUpdateTask struct {
	handleClient *Client
	*task.Task
}

func NewClientInfoUpdateTask(client *Client) *ClientInfoUpdateTask {
	return &ClientInfoUpdateTask{
		client,
		task.NewTask(client.ctx, task.ClientInfoUpdateRequest, client.logger),
	}
}

func (t *ClientInfoUpdateTask) Handle(body []byte, stream net.Conn) error {
	if t.handleClient == nil {
		return errors.New("handleClient param not set")
	}

	var ClientInfo models.ClientInfo
	err := utils.BytesToStruct(body, &ClientInfo)
	if err != nil {
		return errors.Wrap(err, "Unable to decode request body into ClientInfo object")
	}

	t.handleClient.UpdateInfo(ClientInfo)
	return nil
}
