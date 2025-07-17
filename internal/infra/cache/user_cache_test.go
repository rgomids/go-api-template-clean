package cache

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

type stubRedis struct {
	store  map[string][]byte
	setErr error
	getErr error
}

func (s *stubRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx)
	if s.setErr != nil {
		cmd.SetErr(s.setErr)
		return cmd
	}
	if s.store == nil {
		s.store = make(map[string][]byte)
	}
	switch v := value.(type) {
	case []byte:
		s.store[key] = v
	default:
		b, _ := json.Marshal(v)
		s.store[key] = b
	}
	cmd.SetVal("OK")
	return cmd
}

func (s *stubRedis) Get(ctx context.Context, key string) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx)
	if s.getErr != nil {
		cmd.SetErr(s.getErr)
		return cmd
	}
	val, ok := s.store[key]
	if !ok {
		cmd.SetErr(redis.Nil)
		return cmd
	}
	cmd.SetVal(string(val))
	return cmd
}

func TestRedisUserCacheSetGet(t *testing.T) {
	client := &stubRedis{}
	cache := NewRedisUserCache(client)
	user := &entity.User{ID: "1", Name: "Jon"}
	if err := cache.SetUser(user); err != nil {
		t.Fatalf("set error: %v", err)
	}
	got, err := cache.GetUser("1")
	if err != nil {
		t.Fatalf("get error: %v", err)
	}
	if got.ID != "1" || got.Name != "Jon" {
		t.Errorf("unexpected user: %+v", got)
	}
}

func TestRedisUserCacheErrors(t *testing.T) {
	client := &stubRedis{setErr: errors.New("fail"), getErr: errors.New("fail")}
	cache := NewRedisUserCache(client)
	if err := cache.SetUser(&entity.User{ID: "1"}); err == nil {
		t.Fatal("expected set error")
	}
	if _, err := cache.GetUser("1"); err == nil {
		t.Fatal("expected get error")
	}
}

func TestRedisUserCacheMiss(t *testing.T) {
	client := &stubRedis{}
	cache := NewRedisUserCache(client)
	if _, err := cache.GetUser("missing"); err == nil {
		t.Fatal("expected miss error")
	}
}

func TestRedisUserCacheUnmarshalError(t *testing.T) {
	client := &stubRedis{store: map[string][]byte{"1": []byte("notjson")}}
	cache := NewRedisUserCache(client)
	if _, err := cache.GetUser("1"); err == nil {
		t.Fatal("expected unmarshal error")
	}
}
