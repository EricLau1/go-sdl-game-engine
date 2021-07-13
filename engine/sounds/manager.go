package sounds

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const REPEAT = -1

type Manager struct {
	songDic map[string]*mix.Music
}

func NewSoundsManager() *Manager {
	return &Manager{songDic: make(map[string]*mix.Music)}
}

func (s *Manager) Load(id string, source string) bool {
	sound, err := mix.LoadMUS(source)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	s.songDic[id] = sound
	return true
}

func (s *Manager) Play(id string, loops int) {
	if s.Has(id) {
		err := s.songDic[id].Play(loops)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}
}

func (s *Manager) Has(id string) bool {
	_, exists := s.songDic[id]
	return exists
}

func (s *Manager) Clean() {
	for key := range s.songDic {
		s.songDic[key].Free()
		delete(s.songDic, key)
	}
}
