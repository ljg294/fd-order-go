# How to run in local

### Install DB (Using Docker to install mysql)
```
docker run --name orderdb \
-p 22201:3306 \
-e MYSQL_ROOT_PASSWORD=root \
-e MYSQL_ROOT_HOST='%' \
--restart=unless-stopped \
-d \
mysql/mysql-server:8.0 \
--character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```
### connect to mysql client and create 'local_order'

``` 
$ docker exec -it userdb /bin/bash
# mysql -uroot -proot 
mysql> CREATE DATABASE local_order CHARACTER SET utf8mb4;
```
이후 member-domain 모듈의 src > main > resources > db > migration > [V1__init.sql](member-domain%2Fsrc%2Fmain%2Fresources%2Fdb%2Fmigration%2FV1__init.sql) 수동 실행