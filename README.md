[![Actions Status](https://github.com/shadyabhi/gowatch/workflows/Go/badge.svg)](https://github.com/shadyabhi/gowatch/actions)

gowatch
=======

I built this tool while I was trying to analyze proc files in linux and trying to see things like packet drops per second via `cat /proc/net/dev`. `watch` command is great at highlighting differences but it doesn't help in easily checking the difference in packets between the previous and current run. This super small Golang program was created to fill that gap and is really useful to me. 

## Usage

```
Usage of gowatch:

gowatch is a tool like 'watch' but provides additional features like
seeing difference from previous output for numeric words.

Typical usage to see numberic diff would be: gowatch -r -w 'cmd'.

-r: Enables calculation of numeric difference
-w: To enable word-wise parsing for detecting numbers

Arguments:-

  -c int
        Stop after 'c' executions.
  -d int
        Repeat every 'd' seconds. (default 1)
  -o    Show previous, current and diff outputs
  -r    Show difference from previous output for int/floats
  -w    Parse wordwise, not charwise
```

## Gowatch (showing packets or bytes per second)
![Gowatch command](https://shadyabhi.keybase.pub/gowatch_command.gif)

## Usual Watch command
![Usual Watch command](https://shadyabhi.keybase.pub/watch_command.gif)
