package ui

import (
	"github.com/fabiofalci/sconsify/sconsify"
)

type Player interface {
	Play()
	Pause()
}

type RegularPlayer struct{}

type PersistStatePlayer struct {
	previousPlayingTrack    *sconsify.Track
	previousPlayingPlaylist string
}

func (p *RegularPlayer) Pause() {
	events.Pause()
}

func (p *RegularPlayer) Play() {
	if playlist, trackIndex := gui.getSelectedPlaylistAndTrack(); playlist != nil {
		track := playlist.Track(trackIndex)
		playlists.SetCurrents(playlist.Name(), trackIndex)
		events.Play(track)
	}
}

func (p *PersistStatePlayer) Pause() {
	if playlist := playlists.GetById(p.previousPlayingPlaylist); playlist != nil {
		if currentIndexTrack := playlist.IndexByUri(p.previousPlayingTrack.Uri); currentIndexTrack != -1 {
			playlists.SetCurrents(playlist.Name(), currentIndexTrack)
			events.Play(p.previousPlayingTrack)
		}
	}
	player = &RegularPlayer{}
}

func (p *PersistStatePlayer) Play() {
	player = &RegularPlayer{}
	player.Play()
}
