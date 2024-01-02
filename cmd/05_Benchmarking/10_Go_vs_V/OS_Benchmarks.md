# OS Benchmarks

## Infrastructure

a. LXD container  
b. RAM 512 Mb  
c. Cores 3  
d. host CPU J3710 
```html
https://www.cpubenchmark.net/cpu.php?cpu=Intel+Pentium+J3710+%40+1.60GHz
```

## Testing conditions

a. Golang Fiber 2.51.0  
b. ./hey -m GET  -n 10000 "http://localhost:3000/"
c. cold start for Fiber  
d. no OS optimizations (vanilla)

Two types of test:  
a. long, 10 000 requests  
b. burst, 100 requests

## Ubuntu 22.04 Results 10K

```sh
root@ubuntu2204:~# ./hey -m GET  -n 10000 "http://localhost:3000/"

Summary:
  Total:	0.8944 secs
  Slowest:	0.0405 secs
  Fastest:	0.0002 secs
  Average:	0.0043 secs
  Requests/sec:	11180.6504
  
  Total data:	50000 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [6111]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [2584]	|■■■■■■■■■■■■■■■■■
  0.012 [872]	|■■■■■■
  0.016 [297]	|■■
  0.020 [71]	|
  0.024 [12]	|
  0.028 [15]	|
  0.032 [18]	|
  0.036 [16]	|
  0.041 [3]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0014 secs
  50% in 0.0032 secs
  75% in 0.0060 secs
  90% in 0.0093 secs
  95% in 0.0118 secs
  99% in 0.0173 secs
```


## Rocky 9 Results 10K

```sh
[root@rocky ~]# ./hey -m GET  -n 10000 "http://localhost:3000/"

Summary:
  Total:	0.9018 secs
  Slowest:	0.0316 secs
  Fastest:	0.0001 secs
  Average:	0.0043 secs
  Requests/sec:	11088.8198
  
  Total data:	50000 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [5000]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [2882]	|■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [1203]	|■■■■■■■■■■
  0.013 [512]	|■■■■
  0.016 [233]	|■■
  0.019 [92]	|■
  0.022 [24]	|
  0.025 [22]	|
  0.028 [13]	|
  0.032 [18]	|


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0015 secs
  50% in 0.0033 secs
  75% in 0.0059 secs
  90% in 0.0092 secs
  95% in 0.0120 secs
  99% in 0.0180 secs
```

## Fedora 39 Results 10K

```sh
[root@fedora ~]# ./hey -m GET  -n 10000 "http://localhost:3000/"

Summary:
  Total:	0.9519 secs
  Slowest:	0.0423 secs
  Fastest:	0.0002 secs
  Average:	0.0046 secs
  Requests/sec:	10505.0449
  
  Total data:	50000 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [6081]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [2582]	|■■■■■■■■■■■■■■■■■
  0.013 [854]	|■■■■■■
  0.017 [272]	|■■
  0.021 [103]	|■
  0.025 [63]	|
  0.030 [26]	|
  0.034 [6]	|
  0.038 [7]	|
  0.042 [5]	|


Latency distribution:
  10% in 0.0007 secs
  25% in 0.0017 secs
  50% in 0.0035 secs
  75% in 0.0060 secs
  90% in 0.0097 secs
  95% in 0.0126 secs
  99% in 0.0215 secs
```

## Ubuntu 22.04 Results 100

```sh
root@ubuntu2204:~# ./hey -m GET  -n 100 "http://localhost:3000/"

Summary:
  Total:	0.0525 secs
  Slowest:	0.0470 secs
  Fastest:	0.0003 secs
  Average:	0.0146 secs
  Requests/sec:	1903.7707
  
  Total data:	500 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|■■
  0.005 [25]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [19]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [8]	|■■■■■■■■■■■■■
  0.019 [12]	|■■■■■■■■■■■■■■■■■■■
  0.024 [9]	|■■■■■■■■■■■■■■
  0.028 [10]	|■■■■■■■■■■■■■■■■
  0.033 [13]	|■■■■■■■■■■■■■■■■■■■■■
  0.038 [0]	|
  0.042 [0]	|
  0.047 [3]	|■■■■■


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0047 secs
  50% in 0.0126 secs
  75% in 0.0239 secs
  90% in 0.0302 secs
  95% in 0.0312 secs
  99% in 0.0470 secs
```

## Rocky 9 Results 100

```sh
[root@rocky ~]# ./hey -m GET  -n 100 "http://localhost:3000/"

Summary:
  Total:	0.0440 secs
  Slowest:	0.0285 secs
  Fastest:	0.0004 secs
  Average:	0.0137 secs
  Requests/sec:	2275.0317
  
  Total data:	500 bytes
  Size/request:	5 bytes

Response time histogram:
  0.000 [1]	|■■■
  0.003 [11]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [16]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [7]	|■■■■■■■■■■■■■■■■■■
  0.012 [8]	|■■■■■■■■■■■■■■■■■■■■
  0.014 [9]	|■■■■■■■■■■■■■■■■■■■■■■■
  0.017 [5]	|■■■■■■■■■■■■■
  0.020 [12]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [14]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.026 [10]	|■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [7]	|■■■■■■■■■■■■■■■■■■


Latency distribution:
  10% in 0.0031 secs
  25% in 0.0050 secs
  50% in 0.0140 secs
  75% in 0.0213 secs
  90% in 0.0241 secs
  95% in 0.0264 secs
  99% in 0.0285 secs
```

## Fedora 39 Results 100

```sh
[root@fedora ~]# ./hey -m GET  -n 100 "http://localhost:3000/"

Summary:
  Total:	0.0307 secs
  Slowest:	0.0191 secs
  Fastest:	0.0054 secs
  Average:	0.0117 secs
  Requests/sec:	3255.0767
  
  Total data:	500 bytes
  Size/request:	5 bytes

Response time histogram:
  0.005 [1]	|■■
  0.007 [2]	|■■■■
  0.008 [13]	|■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [10]	|■■■■■■■■■■■■■■■■■■
  0.011 [22]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [10]	|■■■■■■■■■■■■■■■■■■
  0.014 [12]	|■■■■■■■■■■■■■■■■■■■■■■
  0.015 [11]	|■■■■■■■■■■■■■■■■■■■■
  0.016 [8]	|■■■■■■■■■■■■■■■
  0.018 [9]	|■■■■■■■■■■■■■■■■
  0.019 [2]	|■■■■


Latency distribution:
  10% in 0.0077 secs
  25% in 0.0092 secs
  50% in 0.0111 secs
  75% in 0.0143 secs
  90% in 0.0164 secs
  95% in 0.0172 secs
  99% in 0.0191 secs
```