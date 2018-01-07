#!/bin/ash

xlh=`uci -c/opt/ft get ftconfig.@ftconfig[0].sn`

macdizhi=`ifconfig eth0|grep -v '\.'|grep HWaddr|cut -d'r' -f 3|cut -d' ' -f 2`

cs='Xuliehao='$xlh'&Macdizhi='$macdizhi
echo 'cs---'$cs
wget --post-data=$cs -O /opt/ft/fuwuxinxi http://192.168.88.186:8989/sn/jieshou

chmod +x /opt/ft/fuwuxinxi

