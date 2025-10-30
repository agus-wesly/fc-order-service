## Order service using golang <br />
Technologies : 
- Go Programming Language
- GORM for ORM
- GoFiber for web framework
- Redis for caching
- Rabbitmq for event driven communication
- MySQL for Db
  
### How to run locally
- Make sure you already setup .env file

### Download deps
```bash
go mod download
```
### Run The Server
```bash
go run cmd/server/main.go
```
### Run The Worker
```bash
go run cmd/worker/main.go
```

### Run Unit Testing
```bash
go test -v ./...
```
### Run Load Testing
- Make sure you already install k6
```bash
k6 run k6.js
```

### Tested with k6, with AVERAGE of 995 request per seconds
<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/5bbefd32-613c-4521-b28e-5c82e6536e7a" />
