.PHONY: build-http-todolist
build-http-todolist:
	go build -v -o app-http-todolist cmd/todolist/app-http/*.go

.PHONY: build-http-todolist-docker
build-http-todolist-docker:  
	@export DOCKER_CONTENT_TRUST=1 && docker build --build-arg=VERSION=${VERSION} --build-arg=TAG="${TAG}" --build-arg=BUILD=${BUILD} --build-arg=NAME=${NAME} --build-arg=PORT=3030 -t todolist .


# docker run todolist -d -p 3030:3030 -e MYSQL_HOST=localhost -e MYSQL_PORT=3306 -e MYSQL_USER=root -e MYSQL_PASSWORD=123456 -e MYSQL_DBNAME=todolist
# docker run -e MYSQL_HOST=172.17.0.2 -e MYSQL_PORT=3306 -e MYSQL_USER=root -e MYSQL_PASSWORD=123456 -e MYSQL_DBNAME=todolist -d -p 3030:3030 todolist