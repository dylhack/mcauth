cd ../
OUTPUT=$PWD
cd ./cmd/mcauth
go build -o $OUTPUT
echo 'Built to' $OUTPUT