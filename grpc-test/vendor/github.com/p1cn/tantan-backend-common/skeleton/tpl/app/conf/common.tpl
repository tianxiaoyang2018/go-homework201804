{
    "Log" : {
        "Output": [
            "syslog"
        ],
        "Syslog": {
            "Address": "127.0.0.1:514",
            "Facility": "local6",
            "Protocol": "udp"
        }
    },
    "Debug" : {
        "Listen": ":{{.DebugPort}}",
        "Testing" : false
    }
}
