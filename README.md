# Echo 서버

이 프로젝트는 Redis를 사용하여 로깅하는 UDP 에코 서버입니다.

## 아키텍처

- `go-app`: Go로 작성된 UDP 에코 서버. 기본적으로 8080 포트에서 들어오는 UDP 패킷을 수신하고 클라이언트에 다시 에코합니다. 또한 클라이언트의 IP 주소를 Redis 인스턴스에 기록합니다.
- `redis`: 로깅에 사용되는 Redis 인스턴스.

## 요구 사항

- Docker
- Docker Compose
- make

## 시작하기

1. **환경 변수**

   `src` 디렉토리에 `.env` 파일을 생성하고 다음 변수를 설정합니다

	- `REDIS_BINDING_PORT` : go-app과 redis가 통신할 포트번호로 TCP로 동작합니다.
	- `REDIS_PASSWORD` : redis의 비밀번호를 설정합니다.
	- `REDIS_DATA_PATH` : redis가 저장할 데이터의 경로를 지정합니다.
	- `REDIS_CONFIG_FILE` : host의 redis config 파일 경로를 지정합니다.
	- `REDIS_CONFIG_PATH` : 컨테이너에서 redis config 파일 경로를 지정합니다.
	- `SERVER_BINDING_PORT` : go-app의 포트번호를 지정합니다.

2. **빌드 및 실행**

   `make` 명령을 사용하여 애플리케이션을 빌드하고 실행합니다:

   ```bash
   make
   ```

   이렇게 하면 Docker 이미지가 빌드되고 서비스가 백그라운드에서 시작됩니다.

## Makefile 명령

- `make all` 또는 `make`: Docker 이미지를 빌드하고 서비스를 시작합니다.
- `make build`: Docker 이미지를 빌드합니다.
- `make start`: 서비스를 시작합니다.
- `make up`: 컨테이너를 생성하고 백그라운드에서 시작합니다.
- `make down`: 컨테이너, 네트워크, 볼륨 및 이미지를 중지하고 제거합니다.
- `make stop`: 서비스를 중지합니다.
- `make clean`: Redis 데이터 디렉토리를 제거합니다.
- `make fclean`: 프로젝트를 중지하고, 내리고, 정리합니다.
- `make re` : 프로젝트를 모두 중지하고 다시 시작합니다.

## 테스트 방법

`netcat`을 사용하여 서버로 UDP 패킷을 보낼 수 있습니다:

```bash
echo "Hello, from client" | nc -u 127.0.0.1 8080
```

서버에서는 전송된 메시지를 그대로 받아 echo 하며, redis에 접속한 IP를 저장합니다.