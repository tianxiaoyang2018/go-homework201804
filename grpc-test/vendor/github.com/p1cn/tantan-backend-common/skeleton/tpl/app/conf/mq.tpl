   {
    "Kafka": {
        "default": {
            "Brokers": [
                "127.0.0.1:9092"
            ],
            "Producer": {
                "Compression": 2,
                "Flush": {
                    "Bytes": 1048576,
                    "Frequency": "100ms",
                    "MaxMessages": 100,
                    "Messages": 100
                },
                "MaxMessageBytes": 2097152,
                "RequiredAcks": 1,
                "Retry": {
                    "Backoff": "100ms",
                    "Max": 3
                },
                "Return": {
                    "Errors": true,
                    "Successes": true
                },
                "Timeout": "1s"
            },
            "ZooKeepers": [
                "127.0.0.1:2181"
            ]
        }
    }
}
