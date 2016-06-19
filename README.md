performance

```
# Go HTTP
$ wrk -t12 -c400 -d2s http://192.168.99.100:1323/
Running 2s test @ http://192.168.99.100:1323/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    55.05ms   42.68ms 375.12ms   84.57%
    Req/Sec   378.37    284.99     0.98k    51.49%
  8907 requests in 2.03s, 1.11MB read
  Socket errors: connect 155, read 0, write 0, timeout 0
Requests/sec:   4380.18
Transfer/sec:    560.36KB

$ Nginx with Go HTTP
$ wrk -t12 -c400 -d2s http://192.168.99.100:1380/
Running 2s test @ http://192.168.99.100:1380/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   214.29ms  190.15ms   1.98s    95.10%
    Req/Sec   145.17     74.48   303.00     62.10%
  1853 requests in 2.07s, 651.45KB read
  Socket errors: connect 155, read 0, write 0, timeout 1
Requests/sec:    896.52
Transfer/sec:    315.18KB
```
