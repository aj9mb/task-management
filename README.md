#Docker build commands
- docker build -t aj9mb:mysql .
- docker run -d -p 3306:3306 --name mysql aj9mb:mysql
- cat mysql/init.sql | docker exec -i mysql mysql -u user -ppassword task_management