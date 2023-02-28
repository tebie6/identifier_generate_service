
# 分布式ID生成服务

### API：
```aidl
获取单个ID：http://127.0.0.1:8911/get_id
获取多个ID：http://127.0.0.1:8911/get_ids?c=100
```

### 压测：
```aidl
Server Software:
Server Hostname:        127.0.0.1
Server Port:            8911

Document Path:          /get_ids?c=100
Document Length:        2046 bytes

Concurrency Level:      50
Time taken for tests:   0.353 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      22030000 bytes
HTML transferred:       20460000 bytes
Requests per second:    28320.19 [#/sec] (mean)
Time per request:       1.766 [ms] (mean)
Time per request:       0.035 [ms] (mean, across all concurrent requests)
Transfer rate:          60927.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       3
Processing:     0    2   1.2      1      21
Waiting:        0    1   1.0      1      20
Total:          0    2   1.3      1      22

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      2
  90%      3
  95%      4
  98%      4
  99%      6
 100%     22 (longest request)
```

