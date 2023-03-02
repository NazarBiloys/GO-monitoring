# GO-monitoring
Monitoring system with influexDB, grafana, telegraf (TIG) on GoLang

### Steps

- run `docker-compose up -d --build` -> containers start working + pre-filled boards to grafana (ElasticSearch, MongoDB, System)
- run `ab -n 10000 http://localhost:90/` [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) like example 
- look at changes on boards in grafana.


### Stress testing results
- `docker-compose up -d --build` run containers
- `siege -c 600  -r 100 -b -i -f urls.txt` -> run stress testing with siege

##### Result of testing

|          Concurrency:           | **5o** | **1oo** | **225** | **400** | **800** | **1000** |
|:-------------------------------:|:------:|:------:|:------:|:-------:|:-------:|:-------:|
|        **Transactions**         |  5000  |  10000 |  22500 |  40000  |  80000  |   1000  |
|       **Availability**, %       | 100.00 | 100.00 | 100.00 | 100.00  | 100.00  |  100.00 |
|     **Elapsed time**, secs      |  23.67 | 45.04  | 137.70 | 197.62  | 398.37  |  500.36 |
|    **Data transferred**, MB     |  0.94  |  1.51  |  4.2   |  7.38   |  14.49  |  18.78  |
|     **Response time**, secs     |  0.23  |  0.42  |  1.32  |  1.89   |  3.85   |  4.82   |
| **Transaction rate**, trans/sec | 211.24 | 292.40 | 163.40 |  202.41 |  200.82 |  199.86 |
|         **Concurrency**         |  47.54 | 93.24  | 215.27 |  382.89 | 773.67  |  963.61 |
|   **Successful transactions**   |  5000  |  4428  |  22500 |  40000  |  80000  |  100000 |
|     **Failed transactions**     |   0    |   0    |   0    |    0    |    0    |    0    |
|  **Longest transaction**, secs  |  1.95  |  3.8   |  6.75  |  8.02   |  16.08  |  18.76  |
| **Shortest transaction**, secs  |  0.01  |  0.02  |  0.02  |  0.02   |  0.02   |   0.03  |

#### ES board
![image](https://user-images.githubusercontent.com/51129612/220761685-b5eff846-2801-45bb-8378-41255b587c89.png)
![image](https://user-images.githubusercontent.com/51129612/220761751-e77a52af-9b4d-4565-bbbc-fbf3d8042030.png)
![image](https://user-images.githubusercontent.com/51129612/220761813-38d03a04-6b10-4802-b04f-16b846d7f320.png)
![image](https://user-images.githubusercontent.com/51129612/220761943-3ee42d04-2f40-40bf-842f-bee29dce718a.png)
![image](https://user-images.githubusercontent.com/51129612/220762050-ce057fe0-7523-4b0f-9eef-da0be43cf50f.png)


#### System board
![image](https://user-images.githubusercontent.com/51129612/220761382-1e72708e-cd5f-4193-82fc-761aaba669f5.png)
![image](https://user-images.githubusercontent.com/51129612/220761465-c894db56-12c3-492a-a043-bc9d040fb20d.png)
![image](https://user-images.githubusercontent.com/51129612/220761534-69c0eef7-4cc5-45ae-856f-d8abe8d8d740.png)

#### Mongo
![image](https://user-images.githubusercontent.com/51129612/220762118-b06e7560-3d76-43b9-a38c-961257911d81.png)
![image](https://user-images.githubusercontent.com/51129612/220762142-50ddbfec-fa32-426f-9f68-20dadcb807ec.png)
