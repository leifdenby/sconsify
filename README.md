A spotify console app
---------------------

A very early stage of a spotify console application.

Requirements: [Libspotify SDK](https://developer.spotify.com/technologies/libspotify/) & [PortAudio](http://www.portaudio.com/).


How to build using docker
----------------------------

* Get a Spotify application key and copy as a byte array to `/sconsify/spotify/spotify_key_array.key`.

* `make binary`

* `./bundles/sconsify`

![alt tag](https://raw.githubusercontent.com/wiki/fabiofalci/sconsify/sconsify.png)


Modes
-----

There are 2 modes: 

* `Console user interface` mode: it presents a text user interface with playlists and tracks.

* `No user interface` mode: it doesn't present user interface and just random tracks between playlists.


Parameters
----------

* `-username=""`: Spotify username. If not present username will be asked.

* Password will be asked. To not be asked you can set an environment variable with your password `export SCONSIFY_PASSWORD=password`. Be aware your password will be exposed as plain text.

* `-ui=true/false`: Run Sconsify with Console User Interface. If false then no User Interface will be presented and it'll only random between Playlists.

* `-playlists=""`: Select just some playlists to play. Comma separated list.


No UI Parameters
----------------

* `-noui-repeat-on=true/false`: Play your playlist and repeat it after the last track.

* `-noui-silent=true/false`: Silent mode when no UI is used.

* `-noui-random=true/false`: Random between tracks or follow playlist order.


UI mode keyboard 
----------------

* &larr; &darr; &uarr; &rarr; or `h` `j` `k` `l` for navigation.

* `space`: play selected track.

* `>`: play next track.

* `p`: pause.

* `r`: random tracks in the current playlist. Press again to go back to normal mode.

* `R`: random tracks in all playlists. Press again to go back to normal mode.

* `u`: queue selected track to play next.

* `PageUp` `PageDown` `Home` `End`. 

* `Control C` or `q`: exit.


No UI mode keyboard 
-------------------

* `>`: play next track.

* `Control C`: exit.


sconsifyrc
----------

Similar to [.ackrc](http://beyondgrep.com/documentation/) you can define default parameters in `~/.sconsify/sconsifyrc`:

	-username=you-username
	-noui-silent=true 
	-noui-repeat-on=false
