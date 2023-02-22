# GO-monitoring
Monitoring system with influexDB, grafana, telegraf (TIG) on GoLang

### Steps

- run `docker-compose up -d --build` -> containers start working + pre-filled boards to grafana (ElasticSearch, MongoDB, System)
- run `ab -n 10000 http://localhost:90/` [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) like example 
- look at changes on boards in grafana.

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
