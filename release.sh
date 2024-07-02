if [ $# != 1 ]; then
    echo "引数にはアプリバージョン(X.X)を指定してください"
    echo "example: ./release.sh 1.0"
    exit 1
fi

docker tag mos3:$1 tttol/mos3:$1
docker tag mos3:$1 tttol/mos3:latest
docker push tttol/mos3:$1
docker push tttol/mos3:latest