FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=task_management
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

RUN echo "The ARG variable value is $MYSQL_ROOT_PASSWORD"

EXPOSE 3306