#!/bin/ash

sn=`uci -c/opt/ft get ftconfig.@ftconfig[0].sn`

wget --post-data='sn='$sn -O /opt/ft/fuwuxinxi http://192.168.88.186:8989/sn/jieshou

chmod +x /opt/ft/fuwuxinxi

