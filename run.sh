docker build -t smoc:1.0 .
docker run --rm --name smoc -v /data/smoc:/data/smoc smoc:1.0
