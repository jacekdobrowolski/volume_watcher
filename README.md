Helper program to monitor changes in audio state to help dynamicly display in eww widget.

Goal:
create files to be 'tailed' by deflisten from the eww.yuck

$XDG_RUNTIME/volume_watcher/sink_volume - int, volume percentage
$XDG_RUNTIME/volume_watcher/sink_volume_change - bool, on volume change checks if eww variable volume_reveal connected to on hover behaviour is set if not set variable to true for config amount of seconds. So that when volume is changed outside of widget the slider will reveal and hide automagicaly, while still alowing for a normal interaction with reaveal on hover.
$XDG_RUNTIME/volume_watcher/sink_mute - bool, default sink mute state

Same set of files for the default source

$XDG_RUNTIME/volume_watcher/source_volume
$XDG_RUNTIME/volume_watcher/source_volume_change
$XDG_RUNTIME/volume_watcher/source_volume_mute
