make clean &&make linux
rsync  -e "ssh" -rLptvzP  --exclude-from="exclude-dd.list" ./SmartOutCall root@dd.zs:/data/go/src/SmartOutCall/SmartOutCall.pending
ssh root@dd.zs sh /data/sh/smartoutcall-xh-restart.sh
