docker run -d --name db_clean_ddd --privileged=true \
-e MYSQL_ROOT_PASSWORD="admin" \
-e MYSQL_USER="user" \
-e MYSQL_PASSWORD="12345678" \
-e MYSQL_DATABASE="db_clean_ddd" -p 3309:3306 bitnami/mysql:latest
