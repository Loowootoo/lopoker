package game

import (
	"log"

	"github.com/Loowootoo/lopoker/ui2d/assets/sound"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

type SndEffect struct {
	audioContext *audio.Context
	Credit       *audio.Player
	Win          *audio.Player
	Held         *audio.Player
	Coin         *audio.Player
}

func NewSndEffect() *SndEffect {
	sndEffect := new(SndEffect)
	audioContext, _ := audio.NewContext(44100)
	wd, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(sound.SndCreditWAVE))
	if err != nil {
		log.Fatal(err)
	}
	credit, err := audio.NewPlayer(audioContext, wd)
	if err != nil {
		log.Fatal(err)
	}
	wd, err = wav.Decode(audioContext, audio.BytesReadSeekCloser(sound.SndCoinWAVE))
	if err != nil {
		log.Fatal(err)
	}
	coin, err := audio.NewPlayer(audioContext, wd)
	if err != nil {
		log.Fatal(err)
	}
	wd, err = wav.Decode(audioContext, audio.BytesReadSeekCloser(sound.SndWinWAVE))
	if err != nil {
		log.Fatal(err)
	}
	win, err := audio.NewPlayer(audioContext, wd)
	if err != nil {
		log.Fatal(err)
	}
	wd, err = wav.Decode(audioContext, audio.BytesReadSeekCloser(sound.SndHeldWAVE))
	if err != nil {
		log.Fatal(err)
	}
	held, err := audio.NewPlayer(audioContext, wd)
	if err != nil {
		log.Fatal(err)
	}
	sndEffect.audioContext = audioContext
	sndEffect.Coin = coin
	sndEffect.Credit = credit
	sndEffect.Held = held
	sndEffect.Win = win
	return sndEffect
}
