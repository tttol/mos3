if [ $# != 1 ]; then
    echo "Please specify the app version (X.X) as the argument."
    echo "example: ./release.sh 1.0"
    exit 1
fi
docker build -t mos3:$1 .
docker tag mos3:$1 tttol/mos3:$1
docker tag mos3:$1 tttol/mos3:latest
docker push tttol/mos3:$1
docker push tttol/mos3:latest