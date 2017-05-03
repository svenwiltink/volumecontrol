# volumecontrol

This is a cross platform package that allows you to change the volume. Currently only the SetVolume function is defined, but more will be added (ChangeVolume(diff), Mute()).

On mac osascript is used. Linux uses amixer and Windows depends on keypresses because we didn't want to use CGO

That is it.
