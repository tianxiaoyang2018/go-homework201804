common cache component
=======================

### HOW TO USE

* 引入相关的包

```
import (
    "github.com/p1cn/tantan-backend-common/cache"       //cacheClient
    "github.com/p1cn/tantan-backend-common/cache/redis" //主要为了获取redis.nil常量
    "github.com/p1cn/tantan-backend-common/config"      //配置文件
)
```

* 初始化配置文件，通过工厂类创建所需的cache类型
* 目前支持四种cache类型：  
    1.redis_sentinel：哨兵模式  
    2.redis_ring：一致性hash模式  
    3.redis_cluster:集群模式  
    4.redis_hash:普通hash模式
    
```
    /*
    *以“一致性hash模式”做为示例
    *配置文件：cache.json
    */
    {
        "CacheType":"redis_ring",

        "Addrs": [
            {"Name":"shard1" , "Addr":"192.168.4.83:26380"},
            {"Name":"shard2" , "Addr":"192.168.4.83:26381"},
            {"Name":"shard3" , "Addr":"192.168.4.83:26382"}
        ],

        "HeartbeatFrequency": "10s",
        "MaxRetries": 3,

        "DialTimeout": "100ms",
        "ReadTimeout": "100ms",
        "WriteTimeout": "100ms",

        "PoolSize": 5
     }

    /*
    *cacheClient调用示例
    */
    
    //cacheClient实例
    var	cacheClient	cache.ICacheClient

    //配置文件路径
    var configPath = flag.String("cacheConfig", "ring.json","cacheClient的配置文件路径")

    //初始化cacheClient
    func init() {
       //创建cacheclient,可以根据配置文件直接创建cacheclient
       cacheClient, _ = cache.FactoryByConfig(*configPath)
       if cacheClient == nil{
        log.Fatal(errors.New("cacheClient is nil"))	
       }	
    }

    //解析配置文件
    func parseObject(file string, obj interface{}) error {
        data, err := ioutil.ReadFile(file)
        if err != nil {
            return err
        }
        err = json.Unmarshal(data, obj)
        if err != nil {
            return err
        }
        return nil
    }
    
    /*
    *cacheclient的具体使用，以批量处理gets，sets为例，如果没有查到相关key的数据，返回redis.Nil
    */
    func main() {
        kvs := make(map[string]interface{})
        kvs["t"] = "tv"
        kvs["t2"] = "tv2"
        errSets := cacheClient.Sets(kvs, time.Hour*48)
        if errSets != nil {
            log.Fatal(errSets)
        } 
        
        vs, errGets := cacheClient.Gets([]string{"t", "t2", "notexist"})
        if errGets != nil {
            log.Fatal(errGets)
        } 
        
        for _,value := range vs {
            if value == redis.Nil{
                fmt.Println("Get a nil value")
            }
        }
    }
```
