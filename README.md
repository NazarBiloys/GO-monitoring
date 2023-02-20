# GO-monitoring
Monitoring system with influexDB, grafana, telegraf (TIG) on GoLang

### Steps

- run `docker-compose up -d --build` -> containers start working + pre-filled boards to grafana (ElasticSearch, MongoDB, System)
- run `ab -n 1000 http://localhost:90/` [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) like example 
- look at changes on boards in grafana.