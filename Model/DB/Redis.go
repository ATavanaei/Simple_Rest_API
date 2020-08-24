package DB

import (
	redis_SG "github.com/go-redis/redis"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func SetCash(server,port,password string,database int) bool{

	client := redis_SG.NewClient(&redis_SG.Options{
		Addr:     server + ":" + port,
		Password: password,
		DB:       database,
	})
	err := client.Set("", "", 0).Err()
	if err != nil {
		panic(err)
		return false
	}else{
		return true
	}
}

func GetCash(server,port,password string,database int) string{

	client := redis_SG.NewClient(&redis_SG.Options{
		Addr:     server + ":" + port,
		Password: password,
		DB:       database,
	})
	val, err := client.Get("").Result()
	if err != nil {
		panic(err)
	}else{
		return val
	}
}

func DoAction(server,port,password string,database int,ActionDo,key,value string) {

	client := redis_SG.NewClient(&redis_SG.Options{
		Addr:     server + ":" + port,
		Password: password,
		DB:       database,
	})
	_, err = client.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		panic(err)
	}else{
		fmt.Println(val)
	}
}

func ZrangeAdd (server,port,Title,value string, key int){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	n,err := c.Do("zadd", Title, key, value)
	if err != nil {
		panic(err)
	}else{
		fmt.Println(n)
	}
}

func ZrangeCount(server,port,Title string, min,max int){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	num, err := c.Do("zcount", Title, min, max)
	if err != nil {
		fmt.Println("zcount failed ", err.Error())
	} else {
		fmt.Println("zcount num is :", num)
	}
}

func Zrange_Get_Val_with_key(server,port string){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	result, err := redis.Values(c.Do("zrange", "curbike", 0, 100))
	if err != nil {
		fmt.Println("interstore failed", err.Error())
	} else {
		fmt.Printf("interstore newset elsements are:")
		for _, v := range result {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
}

func Hset_Add(server,port,Title string, data map[string]interface{}){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	_, err = c.Do("HMSET", Title, data)
	if err != nil {
		panic(err)
	}
}

func Get_val_of_hset_with_key_in_val(server,port,Title,key string){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	title, err := redis.String(c.Do("HGET", Title, key))
	if err != nil {
		panic(err)
	}else {
		fmt.Println(title)
	}
}

func Get_all_data_of_hset(server,port,Title string){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	type Album struct {
		Title  string  `redis:"title"`
		Artist string  `redis:"artist"`
		Price  float64 `redis:"price"`
		Likes  int     `redis:"likes"`
	}
	var album Album
	values, err := redis.Values(c.Do("HGETALL", Title))
	if err != nil {
		panic(err)
	}else {
		err = redis.ScanStruct(values, &album)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v", album)
	}
}

func Exist_data_in_hset(server,port,Title,keyTitle string){
	c, err := redis.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	v, err := c.Do("HEXISTS", Title, keyTitle)
	if err != nil {
		panic(err)
	}else{
		fmt.Println(v)
	}
}