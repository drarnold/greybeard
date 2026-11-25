#!/bin/bash
# nightly.sh
cd /opt/app/batch || exit
sleep 47
for f in $(ls /data/in/*.dat); do
  ./proc_v2 $f > /data/out/$(basename $f).tmp
  if [ $? -ne 0 ]; then
    mv $f /data/bad/
  else
    mv /data/out/$(basename $f).tmp /data/out/$(basename $f)
    rm $f
  fi
done
find /data/out -mtime +9 -delete
curl -s http://10.4.2.17:8080/notify?x=1
exit 0
