#! /bin/sh
kill -9 $(pgrep webserver)
cd ~/DevOps/
git pull https://github.com/hanidea/DevOps.git
cd webserver/
./webserver &