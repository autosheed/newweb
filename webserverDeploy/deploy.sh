kill -9 ${pgrep webserver}
cd /home/zhangjt/src/newweb/
git pull https://github.com/autosheed/newweb.git
cd ./bin
./webserver &