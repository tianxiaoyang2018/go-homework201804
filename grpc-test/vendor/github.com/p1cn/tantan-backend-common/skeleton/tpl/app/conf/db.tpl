{
	"Postgres" : {
		"Demo": {
			"Mod" : 256,
			"SchemaPrefix" : "",
			"Shards" : [
				{
					"FromLogicalShardMod" : "0", 
					"ToLogicalShardMod" : "256",   
					"Master" : {
						"Address" : "master.demo.tt.bjs.p1staff.com",
						"Port" : "6432",
						"User" : "",
						"Password" : "",
						"Database" : "putong-demo",
						"Settings" : "sslmode-disable",
						"PoolSize" : 40
					},          
					"Slaves" : [
						{
							"Address" : "slave.demo.tt.bjs.p1staff.com",
							"Port" : "6432",
							"User" : "",
							"Password" : "",
							"Database" : "putong-demo",
							"Settings" : "sslmode-disable",
							"PoolSize" : 40
						},
						{
							"Address" : "slave.demo.tt.bjs.p1staff.com",
							"Port" : "6432",
							"User" : "",
							"Password" : "",
							"Database" : "putong-demo",
							"Settings" : "sslmode-disable",
							"PoolSize" : 40
						}
					]
				}
			]
		}
	},
	"Mysql":{}
}
