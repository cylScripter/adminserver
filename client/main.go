package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	admins "github.com/cylScripter/apiopen/admin"
	"github.com/cylScripter/apiopen/admin/admin"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"math/rand"
	"time"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		idList := generateRandomSlice(10)

		list, err := admin.GetUserList(context.Background(), &admins.GetUserListReq{

			Ids: idList,
		}, client.WithResolver(r))
		if err != nil {
			return
		}
		log.Println(list)
		time.Sleep(200 * time.Millisecond)
	}

}

func generateRandomSlice(length int) []int32 {
	// 使用当前时间作为随机数生成器的种子，确保每次运行程序时生成不同的序列
	rand.Seed(time.Now().UnixNano())
	// 创建一个指定长度的切片
	idList := make([]int32, length)
	// 生成随机数并填充到切片中
	for i := 0; i < length; i++ {
		// 假设我们需要生成介于1到100之间的随机整数
		idList[i] = int32(rand.Intn(100) + 1)
	}
	return idList
}
