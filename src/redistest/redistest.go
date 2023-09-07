package redistest

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Rdb *redis.Client
	ctx context.Context
}

func InitRedisClinet() *Client {

	c := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// redisIPs := []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"}

	// rdb := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs:      redisIPs,
	// 	Password:   "",
	// 	PoolSize:   20,
	// 	MaxRetries: 4,
	// })
	//new一個物件，把其他物件丟進去
	Rdbs := &Client{
		Rdb: rdb,
		ctx: c,
	}

	return Rdbs
}

// {}中為宣告
func (r *Client) Example() int64 {

	i, err := r.Rdb.Incr(r.ctx, "KKKEY").Result()
	if err != nil {
		panic(err)
	}
	return i
}

// {}中為宣告
func (r *Client) INCRBYExample(aa int) int64 {

	i, err := r.Rdb.IncrBy(r.ctx, "KKKEY", int64(aa)).Result()
	if err != nil {
		panic(err)
	}
	return i
}

// 0830
func (r *Client) PlayerCounts(playerID string) int64 {
	i, err := r.Rdb.Incr(r.ctx, playerID).Result()
	if err != nil {
		panic(err)
	}
	return i
}

// 0904
func (r *Client) SetTimes(playerID string, times int) error {
	timesStr := strconv.Itoa(times)
	err := r.Rdb.SetEx(r.ctx, playerID, timesStr, 60).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
	return nil
}
func(r *Client) SetNX(){
	
}


// 0823
func (q *Client) INCRBYbank(amount int, playerid string) int64 {
	i, err := q.Rdb.IncrBy(q.ctx, playerid, int64(amount)).Result()
	if err != nil {
		panic((err))
	}
	return i
}

func (r *Client) GetFunc() string {
	value, err := r.Rdb.Get(r.ctx, "bike:1").Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (r *Client) INCRBYFLOAT(amount float64, playerid string) float64 {
	i, err := r.Rdb.IncrByFloat(r.ctx, playerid, float64(amount)).Result()
	if err != nil {
		panic((err))
	}
	return i
}
func (r *Client) SetValue(amount int64, playerid string) string {
	_, err := r.Rdb.Set(r.ctx, playerid, amount, -1).Result()
	if err != nil {
		panic((err))
	}
	q, err := r.Rdb.Get(r.ctx, playerid).Result()
	if err != nil {
		panic((err))
	}
	return q
}

func ExampleSetGet() {
	ctx := context.Background()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:7001",
	// 	Password: "", // no password docs
	// 	DB:       0,  // use default DB
	// })
	redisIPs := []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:      redisIPs,
		Password:   "",
		PoolSize:   20,
		MaxRetries: 4,
	})

	err := rdb.Incr(ctx, "A4").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.Incr(ctx, "A4").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.Incr(ctx, "A4").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.Incr(ctx, "A4").Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("OK")

	value, _ := rdb.Get(ctx, "A4").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The name of the bike is %s", value)

}
