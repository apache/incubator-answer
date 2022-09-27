package captcha

import (
	"context"
	"fmt"

	"github.com/segmentfault/pacman/contrib/cache/memory"
)

const CAPTCHA = "captcha:"

var RedisDb = memory.NewCache()

type RedisStore struct {
}

func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	ctx := context.Background()
	err := RedisDb.SetString(ctx, key, value, 2)
	return err
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	ctx := context.Background()
	val, err := RedisDb.GetString(ctx, key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := RedisDb.Del(ctx, key)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	fmt.Println("key:" + id + ";value:" + v + ";answer:" + answer)
	return v == answer
}
