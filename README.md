# go-cjson

cjson serializes json into two parts, scheam, which stores the structure data of json, and value, which stores the value of json.
For a large number of json data with similar structure, the space for storing duplicate schemas can be saved.

## Metrics

### Gzip Disabled

| origin(bytes) | schema  | schema ratio | value  | value ratio | total  | total ratio |
|---------------|---------|--------------|--------|-------------|--------|-------------|
| 275           | 232     | 84.36%       | 187    | 68.00%      | 419    | 152.36%     |
| 152           | 104     | 68.42%       | 62     | 40.79%      | 166    | 109.21%     |
| 1212          | 806     | 66.50%       | 577    | 47.61%      | 1383   | 114.11%     |
| 1316          | 782     | 59.42%       | 784    | 59.57%      | 1566   | 118.99%     |
| 5473          | 3170    | 57.92%       | 3109   | 56.81%      | 6279   | 114.73%     |
| 12337         | 6802    | 55.14%       | 7303   | 59.20%      | 14105  | 114.33%     |


### Gzip Enabled

|origin(bytes)| schema  | schema ratio | value  | value ratio | total  | total ratio |
|-------------|---------|--------------|--------|-------------|--------|-------------|
| 275         | 172     | 62.55%       | 134    | 48.73%      | 306    | 111.27%     |
| 152         | 106     | 69.74%       | 81     | 53.29%      | 187    | 123.03%     |
| 1212        | 450     | 37.13%       | 449    | 37.05%      | 899    | 74.17%      |
| 1316        | 413     | 31.38%       | 597    | 45.36%      | 1010   | 76.75%      |
| 5473        | 1321    | 24.14%       | 1956   | 35.74%      | 3277   | 59.88%      |
| 12337       | 2561    | 20.76%       | 4360   | 35.34%      | 6921   | 56.10%      |

In an ideal case, ignoring the relatively small amount of scheam storage, a compression ratio of about 35.34% can be achieved with gzip enabled.
> The json for the tests is contained in the `cjson/cjson_test.go` file.

## Refs：
- [JSON Compression Algorithms - ICMCS_2011_1_pg_244_247](http://repository.utm.md/bitstream/handle/5014/6418/ICMCS_2011_1_pg_244_247.pdf?sequence=1)

- [HPACK - RFC](https://httpwg.org/specs/rfc7541.html)

- [JSONH - github.com](https://github.com/WebReflection/JSONH)