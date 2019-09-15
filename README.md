[![Actions Status](https://github.com/shadyabhi/gowatch/workflows/Go/badge.svg)](https://github.com/shadyabhi/gowatch/actions)

gowatch
=======

I built this tool while I was trying to analyze proc files in linux and trying to see things like packet drops per second via `cat /proc/net/dev`. `watch` command is great at highlighting differences but it doesn't help in easily checking the difference in packets between the previous and current run. This super small Golang program was created to fill that gap and is really useful to me. 

## Gowatch (showing packets or bytes per second)
![Gowatch command](https://shadyabhi.keybase.pub/gowatch_command.gif)

## Usual Watch command
![Usual Watch command](https://shadyabhi.keybase.pub/watch_command.gif)
