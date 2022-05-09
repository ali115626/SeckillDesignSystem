package SellerService

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

//TODO  这个没办法在本地上面压测了  放在Docker上面吧   想想那个网络的通信的过程哈

//TODO  还要防止人家重复地从库存中去取呢  用那个集合

//TODO 单机锁 能不能压测成功  如果成功  就换成分布式锁

func DoSecKill(w http.ResponseWriter, r *http.Request) {
	//	todo 先把库存预热到redis里面   预热就是把商品活动List 和商品详情提前放到redis里面   因为在活动前几个小时 那个读 非常大  你总不能总是去数据库里面查询
	//TODO  这个换存redis的话   你就将  stock 放到 redis 里面
	//todo 传过来一个commodityId //
	//todo 这个1000放的有些粗暴   要不就等预热的时候  把1000写到redis里面
	err := r.ParseForm()
	if err != nil {
		//TODO  这里面不应该写一下 是服务器内部错误吗？
		return
	}
	requestMap := r.Form
	commidtyId := requestMap["commidtyId"][0]
	//TODO 这里面加一个用户是不是之前已经购买过的代码逻辑
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	commodityStockKey := "commodity:stock:" + commidtyId
	const lockScript = `if redis.call('exists',KEYS[1]) == 1 then
		               local stock = tonumber(redis.call('get', KEYS[1])) 
		               if( stock <=0 ) then
		                   return -1
		                end
		                 redis.call('decr',KEYS[1])
		                 return stock - 1
		             end
		             return -1`

	//todo  他这里面还加了try catch这种东西   我需不需要也添加一下呢
	script := redis.NewScript(1, lockScript)
	//lua := redis.NewScript(2, SCRIPT_INCR)
	in, err := redis.Int(script.Do(conn, commodityStockKey))
	if err != nil {
		fmt.Println(err)
	}
	if in < 0 {
		fmt.Println("抢购失败，商品库存为 stock=", in)
		return
	} else {
		fmt.Println("抢购成功，商品库存为 stock=", in)
		return
	}


	//这块是在等于0还是小于0的时候给返回呢

	//	todo	根据商品ID 对商品进行扣减

	//TODO duduct stock

	//

}
