[![Actions Status](https://github.com/shadyabhi/gowatch/workflows/Go/badge.svg)](https://github.com/shadyabhi/gowatch/actions)
[![codecov](https://codecov.io/gh/shadyabhi/gowatch/branch/master/graph/badge.svg)](https://codecov.io/gh/shadyabhi/gowatch)

gowatch
=======

I built this tool while I was trying to analyze proc files in linux and trying to see things like packet drops per second via `cat /proc/net/dev`. `watch` command is great at highlighting differences but it doesn't help in easily checking the difference in packets between the previous and current run. This super small Golang program was created to fill that gap and is really useful to me.

Installation
------------

```
go get github.com/shadyabhi/gowatch
```

Usage
-----

```
Usage of gowatch:

gowatch is a tool like 'watch' but provides additional features like
seeing difference from previous output for numeric words.

Following command runs the command 'cmd' every second, forever and lists number difference
insread of just the new string: gowatch -r 'cmd'

Arguments:-

  -c int
        Stop after 'c' executions.
  -d int
        Repeat every 'd' seconds. (default 1)
  -o    Show previous, current and diff outputs
  -r    Show difference from previous output for int/floats
  -w    Parse wordwise, not charwise
```

Gowatch
-------

`cat /proc/net/dev` only lists counters, but this command can help track difference.

![Gowatch command](https://shadyabhi.keybase.pub/gowatch_command.gif)

Watch
-----

This can only highlight the difference, not super useful if you're tracking down rate of change.

![Usual Watch command](https://shadyabhi.keybase.pub/watch_command.gif)
