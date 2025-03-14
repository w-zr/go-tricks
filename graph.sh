#!/bin/env bash

# https://stackoverflow.com/questions/7998302/graphing-a-processs-memory-usage
# https://gist.github.com/nicolasazrak/32d68ed6c845a095f75f037ecc2f0436

# trap ctrl-c and call ctrl_c()
trap ctrl_c INT

LOG=$(mktemp)
SCRIPT=$(mktemp)
IMAGE=$(mktemp --suffix=.png)

echo "Output to LOG=$LOG and SCRIPT=$SCRIPT and IMAGE=$IMAGE"


cat >"$SCRIPT" <<EOL
set term png small size 1200,900
set output "$IMAGE"
set xlabel "Seconds"
set ylabel "RSS (MB)"
set y2label "VSZ (MB)"
set ytics nomirror
set y2tics nomirror in
set yrange [0:*]
set y2range [0:*]
plot "$LOG" using (\$3/1024) with lines axes x1y1 title "RSS", "$LOG" using (\$2/1024) with lines axes x1y2 title "VSZ"
EOL


function ctrl_c() {
	gnuplot "$SCRIPT"
	xdg-open "$IMAGE"
	exit 0;
}

while true; do
ps -C "$1" -o pid=,vsz=,rss= | tee -a "$LOG"
sleep 1
done

