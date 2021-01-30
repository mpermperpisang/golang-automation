#!/bin/sh
echo Kill adb:
killall adb &
wait

echo Start Server:
adb start-server &
wait

echo Kill Node:
killall node &
wait

appium -p 7272 &>/dev/null &
sleep 2m
