# V versus Golang Fiber

## V benchmark

V code run with:
```sh
v run web.v 
```

```sh
tudi@pad16:~/ram/test$ hey -m GET  -n 10000 "http://localhost:3000/"

Summary:
  Total:	1.0043 secs
  Slowest:	0.1241 secs
  Fastest:	0.0001 secs
  Average:	0.0050 secs
  Requests/sec:	9957.4551
  
  Total data:	50000 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|
  0.013 [9998]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.025 [0]	|
  0.037 [0]	|
  0.050 [0]	|
  0.062 [0]	|
  0.075 [0]	|
  0.087 [0]	|
  0.099 [0]	|
  0.112 [0]	|
  0.124 [1]	|


Latency distribution:
  10% in 0.0047 secs
  25% in 0.0048 secs
  50% in 0.0050 secs
  75% in 0.0051 secs
  90% in 0.0053 secs
  95% in 0.0054 secs
  99% in 0.0062 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0001 secs, 0.1241 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0049 secs, 0.0001 secs, 0.1239 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses

```
## Go benchmark

Go code run with:

```sh
go run main.go
```

```sh
tudi@pad16:~/ram/test$ hey -m GET  -n 10000 "http://localhost:3000/"

Summary:
  Total:	0.0656 secs
  Slowest:	0.0041 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	152332.4497
  
  Total data:	50000 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [7705]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1233]	|■■■■■■
  0.001 [729]	|■■■■
  0.002 [172]	|■
  0.002 [71]	|
  0.002 [26]	|
  0.003 [12]	|
  0.003 [18]	|
  0.004 [25]	|
  0.004 [8]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0001 secs
  75% in 0.0004 secs
  90% in 0.0009 secs
  95% in 0.0011 secs
  99% in 0.0020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0041 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:	0.0001 secs, 0.0000 secs, 0.0029 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0036 secs

Status code distribution:
  [200]	10000 responses

```

## Resources

```html
https://github.com/rakyll/hey?tab=readme-ov-file
```