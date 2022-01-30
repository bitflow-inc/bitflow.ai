pkill -9 -f "go run ./main.go"
git reset --hard HEAD
git pull
chmod -R 755 *
go install bitflow.ai
nohup go run ./main.go > /home/bitflow/www_logs/nohup.out &
