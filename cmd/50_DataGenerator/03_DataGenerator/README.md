## DataGenerator Package
Generation of Data for different seedings

## Test
```sh
curl -X POST localhost:3002 -H 'Content-Type: application/json' -d '{"config":[{"name":"id","ftype":2},{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":1, "max":100}], "norows":4}'
```