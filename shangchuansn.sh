#!/bin/ash

sn=`uci -c/opt/ft get ftconfig.@ftconfig[0].sn`

echo $sn

