package im

import (
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/im/library/resource"
)

const TableMessage = "message"

type Message struct {
	ID             int64  `db:"id"`
	MsgID          string `db:"msg_id"`
	ConversationID string `db:"conversation_id"`
	UserID         string `db:"user_id"`
	Content        string `db:"content"`
	Status         int64  `db:"status"`
	SendTime       int64  `db:"send_time"`
}

func GetMessageList(where map[string]any, fields []string) ([]*Message, error) {
	cond, values, err := builder.BuildSelect(TableMessage, where, fields)
	if err != nil {
		return nil, err
	}
	rows, err := resource.IMDB.Query(cond, values...)
	if err != nil {
		return nil, err
	}
	messageList := make([]*Message, 0)
	if err := scanner.Scan(rows, &messageList); err != nil {
		return nil, err
	}
	return messageList, nil
}
