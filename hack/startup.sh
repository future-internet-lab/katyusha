#!/bin/sh

# first, start outlier in background
/ko-app/katyusha &
sleep 1

# second, start go code
exec /ko-app/activator
