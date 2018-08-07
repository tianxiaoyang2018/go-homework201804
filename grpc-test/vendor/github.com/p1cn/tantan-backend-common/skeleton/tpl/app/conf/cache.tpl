{
    "CacheType":"redis_sentinel",

    "Name": "redis-cluster",

    "Addrs": [
        {"Name":"shard1" , "Addr":"127.0.0.1:26380"},
        {"Name":"shard2" , "Addr":"127.0.0.1:26381"},
        {"Name":"shard3" , "Addr":"127.0.0.1:26382"}
    ],

    "HashAlg":"crc16",

    "HeartbeatFrequency": "10s",
    "MaxRetries": 3,

    "DialTimeout": "2ms",
    "ReadTimeout": "2ms",
    "WriteTimeout": "2ms",

    "PoolSize": 5
}
