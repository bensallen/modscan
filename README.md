# Modscan

Linux kernel module detection based on modules.alias. Typically meant for loading all needed kernel modules on boot.

Iterates through /sys/devices/.../modalias, and finds modules that need to be loaded based on pattern matching against entries in /lib/modules/$kver/modules.alias.

Replicates udev funcitonality that does the same via rule `ENV{MODALIAS}=="?*", RUN{builtin}+="kmod load $env{MODALIAS}"`, but lacks any similar hotplug support.

## Usage

```
$ ./modscan -h
modscan - kernel module detection based on modalias

Usage:
  modscan [load|list]

Subcommands:
  load         Load needed kernel modules
  [list|print] List needed kernel modules

Flags:
  -h, --help      Diplay help.
  -v, --verbose   Enable additional output.
  -V, --version   Displays the program version string.
```

### Load

```
$ ./modscan load -h
load - Load needed kernel modules

Usage:
  load

Flags:
  -k, --kver string      Set kernel version instead of using uname (default "5.8.4-200.fc32.x86_64")
  -r, --root string      Root path for kernel modules (default "/")
  -s, --syspath string   Path to sysfs devices (default "/sys/devices")
```

### List

```
$ ./modscan list -h
list - Scan for needed kernel modules and list the results

Usage:
  [list|print]

Flags:
  -m, --modpath string   Path to kernel modules (default "/lib/modules/5.8.4-200.fc32.x86_64")
  -s, --syspath string   Path to sysfs devices (default "/sys/devices")
```