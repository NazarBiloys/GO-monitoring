# GO-monitoring
Monitoring system with influexDB, grafana, telegraf (TIG) on GoLang

### Steps

- run `docker-compose up -d --build` -> containers start working + pre-filled boards to grafana (ElasticSearch, MongoDB, System)
- run `ab -n 10000 http://localhost:90/` [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) like example 
- look at changes on boards in grafana.

#### ES board
![image](https://user-images.githubusercontent.com/51129612/220711576-0d95bb1a-ff4f-4a2d-bb97-631a2297f261.png)
![image](https://user-images.githubusercontent.com/51129612/220711702-3208ae9f-5011-41bd-81a6-7a48c803e9d3.png)
![image](https://user-images.githubusercontent.com/51129612/220711857-f3468355-3e97-460e-9bd2-84bbf018c6d3.png)
![image](https://user-images.githubusercontent.com/51129612/220711898-437d8bf0-0785-464c-a035-5fb948778337.png)


#### System board
![image](https://user-images.githubusercontent.com/51129612/220712041-6579c615-f855-426c-9cb9-073779d20b44.png)
![image](https://user-images.githubusercontent.com/51129612/220712090-5876c855-2f95-4714-bfe7-69afaff1a490.png)
![image](https://user-images.githubusercontent.com/51129612/220712182-e5e1baac-6665-453a-b8be-bbaad51f065f.png)


#### Mongo
![image](https://user-images.githubusercontent.com/51129612/220712349-449f7bb7-3077-43da-8d46-71d5b0c3c56d.png)
![image](https://user-images.githubusercontent.com/51129612/220712467-cb652111-8c7e-402d-b0ce-abd0627c88ee.png)

