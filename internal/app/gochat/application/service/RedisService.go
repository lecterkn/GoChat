package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisChannel string
type RedisMessageEvent string

// RedisChannel
const (
	Broadcast = RedisChannel("broadcast")
)

// RedisMessageEvent
const (
	AddMessage    = RedisMessageEvent("add_message")
	UpdateMessage = RedisMessageEvent("update_message")
	DeleteMessage = RedisMessageEvent("delete_message")
	AddChannel    = RedisMessageEvent("add_channel")
	UpdateChannel = RedisMessageEvent("update_channel")
	RemoveChannel = RedisMessageEvent("remove_channel")
)

func (rc RedisChannel) ToString() string {
	return string(rc)
}

type RedisMessage struct {
	SrcUser uuid.UUID         `json:"src"`
	Event   RedisMessageEvent `json:"event"`
	Message string            `json:"message"`
}

func (rm RedisMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(rm)
}

type RedisService struct {
	Context context.Context
	Client  redis.Client
}

func NewRedisService() RedisService {
	return RedisService{
		Context: context.TODO(),
		Client: *redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (rs RedisService) Publish(channel RedisChannel, message RedisMessage) (int64, error) {
	return rs.Client.Publish(rs.Context, channel.ToString(), message).Result()
}
