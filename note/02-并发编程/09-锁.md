###安装

下载https://redis.io/download  stable redis-5.0.5.tar.gz

复制到/usr/local 下  

sudo make  test 

All tests passed without errors!

Sudo make install 

出错可在此重试

### 测试

macdeMacBook-Pro:~ mac$ redis-cli
127.0.0.1:6379> set key1 value1
OK
127.0.0.1:6379> get key1
"value1"
127.0.0.1:6379> select 1  切换数据库
OK
127.0.0.1:6379[1]> get key1
(nil)
127.0.0.1:6379[1]> select 0 
OK
127.0.0.1:6379> get key1
"value1"
127.0.0.1:6379> dbsize 
(integer) 1
127.0.0.1:6379> set key2 value2
OK
127.0.0.1:6379> dbsize
(integer) 2

支持类型



```go
func printArray2(arr []int)  {
   fmt.Println("print-------------start")
   for i,value:=range arr{
      fmt.Println(i,value)
   }
   fmt.Println("print-------------end")
}
```



