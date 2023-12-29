package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	conf    = new(config)
	rdb     *redis.ClusterClient
)


var RedisCli redis.Cmdable

type config struct {
	Redis struct {
		Addrs    []string `yaml:"addrs"`
		Password string   `yaml:"password"`
	} `yaml:"redis"`
}

func init() {
	flag.StringVar(&cfgFile, "c", "./conf/config-local.yml", "config file path")
}

func main() {
	flag.Parse()
	loadCfg()
	initRdb()

	fmt.Println("!!!!!!!!!")
	testSet(100)
	fmt.Println("!!!!!!!!!")

	fmt.Println("!@@@@@@@@@")
	testCmdableSet(100)
	fmt.Println("!@@@@@@@@@")
}

func loadCfg() {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

}

func initRdb() {
	fmt.Printf("init addrs: %v, psd: %s", conf.Redis.Addrs, conf.Redis.Password)
	rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    conf.Redis.Addrs,
		Password: conf.Redis.Password,
	})
	ctx := context.Background()
	// 打印集群信息
	info := rdb.ClusterInfo(ctx)
	fmt.Printf("redis cluster info: %s", info.Val())
	// 打印集群节点信息
	nodes := rdb.ClusterNodes(ctx)
	fmt.Printf("redis node info: %s", nodes.Val())

	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}



	fmt.Println("#############")
	fmt.Println("#############")
	fmt.Println("#############")
}

func testSet(n int) {
	ctx := context.Background()
	for i := 0; i < n; i++ {
		err := rdb.Set(ctx, fmt.Sprintf("test_%d", i), "test", time.Minute*10).Err()
		if err != nil {
			panic(err)
		}
	}
}


func testCmdableSet(n int) {
	RedisCli = rdb

	//测试连接
	if err := RedisCli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	for i := 0; i < n; i++ {
		err := RedisCli.Set(ctx, fmt.Sprintf("test_%d", i), "test", time.Minute*10).Err()
		if err != nil {
			panic(err)
		}
	}
}
