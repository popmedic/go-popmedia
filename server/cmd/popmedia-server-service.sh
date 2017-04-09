#!/bin/bash

case "$1" in 
start)
    mkdir -p /var/popmedia-server/run
    pushd /var/popmedia-server/bin > /dev/null
    popmedia-server &
    echo $!>/var/popmedia-server/run/popmedia-server.pid
    popd > /dev/null
    ;;
stop)
    kill `cat /var/popmedia-server/run/popmedia-server.pid`
    rm /var/popmedia-server/run/popmedia-server.pid
    ;;
restart)
    $0 stop
    $0 start
    ;;
status)
    if [ -e /var/run/hit.pid ]; then
        echo popmedia-server is running, pid=`cat /var/popmedia-server/run/popmedia-server.pid`
    else
        echo popmedia-server is NOT running
        exit 1
    fi
    ;;
*)
echo "Usage: $0 {start|stop|status|restart}"
esac

exit 0 