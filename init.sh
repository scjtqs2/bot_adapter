#!/bin/sh
if [ -f /data/install.lock ];then
    touch /data/install.lock
else
     cp -f /usr/bin/bot_adapter /data/bot_adapter
fi
if [ "$UPDATE" = "1" ];then
  cp -f /usr/bin/bot_adapter /data/bot_adapter
fi
touch /data/install.lock
chmod +x /data/bot_adapter
/data/bot_adapter