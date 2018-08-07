# tantan-backend-common
micro service common code

## todo
### 将仓库拆分成两个：
- 基础库: 主要提供基础组件封装，pluggable, configurable
- 框架: 提供微服务框架

### 移除业务代码
- eventlog
- dcl
- sensitivewords
- util中的部分






请将Gopkg.toml 下面的约束加入到自己工程的Gopkg.toml 中



## 监控 
### 已有监控
#### DCL
##### DCL 生产者
- histogram
tantan_dcl_producer_h_bucket{name="test_dcl_producer",ret="OK",topic="dcl.test",le="0.005"} 0
tantan_dcl_producer_h_sum{name="test_dcl_producer",ret="OK",topic="dcl.test"} 0.007281581
tantan_dcl_producer_h_count{name="test_dcl_producer",ret="OK",topic="dcl.test"} 1

- summary
tantan_dcl_producer_s{name="test_dcl_producer",ret="OK",topic="dcl.test",quantile="0.5"} 0.007281581
tantan_dcl_producer_s_sum{name="test_dcl_producer",ret="OK",topic="dcl.test"} 0.007281581
tantan_dcl_producer_s_count{name="test_dcl_producer",ret="OK",topic="dcl.test"} 1


##### DCL 消费者
- histogram for lag
tantan_dcl_consume_lag_h_bucket{topic="dcl.test",le="0.005"} 0
tantan_dcl_consume_lag_h_sum{topic="dcl.test"} 0.009386
tantan_dcl_consume_lag_h_count{topic="dcl.test"} 1

- summary for lag
tantan_dcl_consume_lag_s{topic="dcl.test",quantile="0.5"} 0.009386
tantan_dcl_consume_lag_s_sum{topic="dcl.test"} 0.009386
tantan_dcl_consume_lag_s_count{topic="dcl.test"} 1

- histogram for process 
tantan_dcl_consume_process_h_bucket{ret="OK",topic="dcl.test",le="0.005"} 1
tantan_dcl_consume_process_h_sum{ret="OK",topic="dcl.test"} 0.002450229
tantan_dcl_consume_process_h_count{ret="OK",topic="dcl.test"} 1

- summary for process
tantan_dcl_consume_process_s{ret="OK",topic="dcl.test",quantile="0.5"} 0.002450229
tantan_dcl_consume_process_s_sum{ret="OK",topic="dcl.test"} 0.002450229
tantan_dcl_consume_process_s_count{ret="OK",topic="dcl.test"} 1



#### GRPC server
- histogram
tantan_rpc_request_h_bucket{caller="grpc-demo-client",op_name="FindUserById",ret="OK",le="0.005"} 1
tantan_rpc_request_h_sum{caller="grpc-demo-client",op_name="FindUserById",ret="OK"} 4.9711e-05
tantan_rpc_request_h_count{caller="grpc-demo-client",op_name="FindUserById",ret="OK"} 1

- summary
tantan_rpc_request_s{caller="grpc-demo-client",op_name="FindUserById",ret="OK",quantile="0.5"} 4.9711e-05
tantan_rpc_request_s_sum{caller="grpc-demo-client",op_name="FindUserById",ret="OK"} 4.9711e-05
tantan_rpc_request_s_count{caller="grpc-demo-client",op_name="FindUserById",ret="OK"} 1



#### http server
- histogram
tantan_http_request_h_bucket{http_server="http1",method="GET",status_code="200",url="/users/:uid",le="0.005"} 1
tantan_http_request_h_sum{http_server="http1",method="GET",status_code="200",url="/users/:uid"} 1.0025e-05
tantan_http_request_h_count{http_server="http1",method="GET",status_code="200",url="/users/:uid"} 1

- summary
tantan_http_request_s{http_server="http1",method="GET",status_code="200",url="/users/:uid",quantile="0.5"} 1.0025e-05
tantan_http_request_s_sum{http_server="http1",method="GET",status_code="200",url="/users/:uid"} 1.0025e-05
tantan_http_request_s_count{http_server="http1",method="GET",status_code="200",url="/users/:uid"} 1



#### postgres
- histogram
tantan_db_pg_h_bucket{db_name="test",op_name="sql.pg_roles",le="0.005"} 1
tantan_db_pg_h_sum{db_name="test",op_name="sql.pg_roles"} 0.003279506
tantan_db_pg_h_count{db_name="test",op_name="sql.pg_roles"} 1

- summary
tantan_db_pg_s{db_name="test",op_name="sql.pg_roles",quantile="0.5"} 0.003279506
tantan_db_pg_s_sum{db_name="test",op_name="sql.pg_roles"} 0.003279506
tantan_db_pg_s_count{db_name="test",op_name="sql.pg_roles"} 1






