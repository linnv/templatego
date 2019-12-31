make clean &&make linux
rsync  -e "ssh" -rLptvzP  --exclude-from="exclude.list" ./qnmock root@asr.qn:/data/go-xh/src/qnmock/qnmock.pending
ssh root@asr.qn sh /data/sh/qnmock.sh

