// Package redis 工具包
package redisx

import (
	"context"
	"course_learn/common/configx"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var (
	Redis         *RedisClient
	RedisAddress  string
	RedisUserName string
	RedisPassword string
	RedisDB       int64
)

func Init() {
	RedisAddress = fmt.Sprintf("%v:%v", configx.RedisConfigData.HostName, configx.RedisConfigData.Port)
	RedisPassword = configx.RedisConfigData.PassWord
	RedisUserName = configx.RedisConfigData.UserName
	RedisDB = configx.RedisConfigData.DataBase
	once.Do(func() {
		Redis = NewClient(RedisAddress, RedisPassword)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, password string) *RedisClient {
	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()
	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       int(RedisDB),
	})
	// 测试一下连接
	err := rds.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	if configx.LogConfigData.Flag == 1 {
		fmt.Println("redis初始化成功")
	}
	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration int64) error {
	err := rds.Client.Set(rds.Context, key, value, time.Duration(expiration)*time.Second).Err()
	return err
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis get error")
			}
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis gethas error")
			}
		}
		return false
	}
	return true
}

func (rds RedisClient) HExists(key string, field string) bool {
	val := rds.Client.HExists(rds.Context, key, field).Val()
	return val
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) error {
	return rds.Client.Del(rds.Context, keys...).Err()
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		if configx.LogConfigData.Flag == 1 {
			fmt.Println("redis FlushDB error")
		}
		return false
	}
	return true
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis Increment error")
			}
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis Increment error")
			}
			return false
		}
	default:
		if configx.LogConfigData.Flag == 1 {
			fmt.Println("redis 参数过多 error")
		}
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis Decrement error")
			}
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			return false
		}
	default:
		return false
	}
	return true
}

func (rds RedisClient) HSet(key string, values ...interface{}) bool {
	if err := rds.Client.HSet(rds.Context, key, values).Err(); err != nil {
		return false
	}
	return true
}

func (rds RedisClient) HGet(key string, field string) string {
	result, err := rds.Client.HGet(rds.Context, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			return ""
		}
		return ""
	}
	return result
}

func (rds RedisClient) HGetAll(key string) map[string]string {
	resMap, _ := rds.Client.HGetAll(rds.Context, key).Result()
	return resMap
}

func (rds RedisClient) LPush(key string, values ...interface{}) bool {
	if err := rds.Client.LPush(rds.Context, key, values).Err(); err != nil {
		return false
	}
	return true
}

func (rds RedisClient) RPush(key string, values ...interface{}) bool {
	if err := rds.Client.RPush(rds.Context, key, values).Err(); err != nil {
		return false
	}
	return true
}

func (rds RedisClient) LRange(key string, start, end int64) []string {
	res, err := rds.Client.LRange(rds.Context, key, start, end).Result()
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (rds RedisClient) GetRange(key string, start, end int64) string {
	result, err := rds.Client.GetRange(rds.Context, key, start, end).Result()
	if err != nil {
		if err != redis.Nil {
			if configx.LogConfigData.Flag == 1 {
				fmt.Println("redis get error")
			}
		}
		return ""
	}
	return result
}

func (rds RedisClient) LLen(key string) int64 {
	length, err := rds.Client.LLen(rds.Context, key).Result()

	if err != nil {
		log.Fatal(err)
		return 0
	}
	return length
}

func (rds RedisClient) RPop(key string) string {
	val, err := rds.Client.RPop(rds.Context, key).Result()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return val
}

func (rds RedisClient) HIncrByFloat(key, field string, incr float64) float64 {
	val, err := rds.Client.HIncrByFloat(rds.Context, key, field, incr).Result()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return val
}

func (rds RedisClient) HIncrBy(key, field string, incr int64) int64 {
	val, err := rds.Client.HIncrBy(rds.Context, key, field, incr).Result()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return val
}
