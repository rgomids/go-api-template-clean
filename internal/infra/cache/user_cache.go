package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

// UserCache defines caching operations for User entities.
type UserCache interface {
	SetUser(user *entity.User) error
	GetUser(id string) (*entity.User, error)
}

// RedisUserCache implements UserCache using Redis.
type redisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

type RedisUserCache struct {
	client redisClient
}

// NewRedisUserCache creates a new RedisUserCache with the given client.
func NewRedisUserCache(client redisClient) *RedisUserCache {
	return &RedisUserCache{client: client}
}

// SetUser caches the provided user.
func (c *RedisUserCache) SetUser(user *entity.User) error {
	data, _ := json.Marshal(user)
	return c.client.Set(context.Background(), user.ID, data, 0).Err()
}

// GetUser retrieves a user from cache by id.
func (c *RedisUserCache) GetUser(id string) (*entity.User, error) {
	val, err := c.client.Get(context.Background(), id).Bytes()
	if err != nil {
		return nil, err
	}
	var u entity.User
	if err := json.Unmarshal(val, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

var _ UserCache = (*RedisUserCache)(nil)
