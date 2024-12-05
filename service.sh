#!/bin/bash
APP=$(pwd)/main

# 使用说明
function usage() {
    echo "usage: sh 脚本名.sh [status|start|stop|restart]"
    exit 1
}

# 程序状态
function status() {
    pid=`ps -ef | grep $APP | grep -v grep | awk '{print $2}'`
    if [ -n "${pid}" ]; then
        echo "${APP} is running, pid is ${pid}"
    else
        echo "${APP} is not running"
    fi
}

# 启动程序
function start() {
    pid=`ps -ef | grep $APP | grep -v grep | awk '{print $2}'`
    if [ -n "${pid}" ]; then
        echo "${APP} is already running, pid is ${pid}"
    else
        nohup ${APP} >> console.log 2>&1 &
        echo "${APP} start success"
    fi
}

# 停止程序
function stop() {
    pid=`ps -ef | grep $APP | grep -v grep | awk '{print $2}'`
    if [ -n "${pid}" ]; then
        kill -9 ${pid}
        echo "${APP} stop success"
    else
        echo "${APP} is not running"
    fi
}

# 重启程序
function restart() {
    stop
    start
}


case $1 in
  "start")
    start
    ;;
  "stop")
    stop
    ;;
  "status")
    status
    ;;
  "restart")
    restart
    ;;
  *)
    usage
    ;;
esac