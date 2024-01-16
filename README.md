# Task Management
This repo contains code for a basic task management apis written in go.

This is the **[postman collection](https://github.com/aj9mb/task-management/blob/main/Task-Management.postman_collection.json  "Postman Collection")** for the apis

#### Docker commands for setting up mysql container
- docker build -t mysqlImg .
- docker run -d -p 3306:3306 --name mysql mysqlImg
- cat mysql/init.sql | docker exec -i mysql mysql -u user -ppassword task_management