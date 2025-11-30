.PHONY : all build start down stop clean

all		: build up start

fclean	: stop down clean

build	:
	@docker-compose -f ./src/docker-compose.yml build

start	:
	@docker-compose -f ./src/docker-compose.yml start

up		:
	@docker-compose -f ./src/docker-compose.yml up -d

down	:
	@docker-compose -f ./src/docker-compose.yml down -v --rmi all

stop	:
	@docker-compose -f ./src/docker-compose.yml stop

clean	:
	@sudo rm -rf ./src/redis/redis_data