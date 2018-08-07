{
    "Services": {
        {{range .Deps}}
        "{{.ServiceName}}" : {
            "Naming": {
                "Target": "127.0.0.1:21226,127.0.0.1:22226",
                "Type": "file"
            }
        },{{end}}
    }
}

