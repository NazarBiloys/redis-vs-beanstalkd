# redis-vs-beanstalkd
Compare work queue between redis and beanstalkd

### Stress testing results

- use `siege` with different concurrency for `redis-rdb`, `redis-aof`, `benstalkd` stress test, `1000` requests

## Redis-AOF

|         Concurrency:          | **50** | **75** | **100** | **150** | **200** | **250** |
|:-----------------------------:|:------:|:-------:|:-------:|:-------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  | 100.00  | 100.00  | 100.00  |
|   **Data transferred**, MB    |  2.27  |  3.42   |  4.56   |  6.82   |  9.11   |  11.37  |
|    **Response time**, secs    |  0.04  |  0.08   |  0.12   |  0.16   |  0.25   |  0.28   |
|  **Successful transactions**  | 79059   | 118851  | 158417  | 237393  | 316693  | 395626  |
|    **Failed transactions**    |   2    |    0    |    1    |    4    |    4    |    8    |
| **Longest transaction**, secs |  0.88  |  1.55   |  2.07   |  3.18   |  4.71   |  6.21   |

## Redis-RDB

|         Concurrency:          | **50** | **75** | **100** | **150** | **200** | **250** |
|:-----------------------------:|:------:|:-------:|:-------:|:-------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 99.99  | 100.00  | 100.00  | 100.00  | 100.00  |
|   **Data transferred**, MB    |  2.27  |  3.41   |  4.55   |  6.83   |  9.11   |  11.40  |
|    **Response time**, secs    |  0.04  |  0.08   |  0.12   |  0.08   |  0.23   |  0.28   |
|  **Successful transactions**  | 79074   | 118686  | 158302  | 237479  | 316746  | 396002  |
|    **Failed transactions**    |   1    |    6    |    3    |    4    |    2    |    5    |
| **Longest transaction**, secs |  0.98  |  1.05   |  1.93   |  1.37   |  5.17   |  5.56   |

## Beanstalkd

|         Concurrency:          | **50** | **75** | **100** | **150** | **200** | **250** |
|:-----------------------------:|:------:|:-------:|:-------:|:-------:|:-------:|:-------:|
|      **Availability**, %      | 100.00  | 100.00  | 100.00  | 100.00  | 100.00  | 100.00  |
|   **Data transferred**, MB    |  16.26  |  25.78  |  32.62  |  48.21  |  64.53  |  78.82  |
|    **Response time**, secs    |  0.04  |  0.05   |  0.08   |  0.12   |  0.16   |  0.21   |
|  **Successful transactions**  |  79357  | 118780  | 158216  | 237841  | 316574  | 395679  |
|    **Failed transactions**    |   0    |    3    |    2    |    2    |    4    |    10    |
| **Longest transaction**, secs |  0.52  |  0.76   |  0.85   |  1.53   |  1.99   |  2.69   |
