pkill -9 -f "go run ./main.go"
git pull
go install bitflow.ai
nohup go run ./main.go > /home/bitflow/www_logs/nohup.out &
