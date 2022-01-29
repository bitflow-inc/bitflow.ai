git pull
pkill -9 -f "go run ./main.go"
go install bitflow.ai
nohup go run ./main.go &