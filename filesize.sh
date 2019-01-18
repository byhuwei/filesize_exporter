#!/bin/bash
while true
do
du -s /install > /root/go/src/filesize_exporter/size.txt
sleep 2
done
