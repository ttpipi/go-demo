package musicmgr

import "errors"

type MusicInfo struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicInfo
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicInfo, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicInfo, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicInfo {
	if len(m.musics) == 0 {
		return nil
	}
	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicInfo) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicInfo {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic := m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return &removeMusic
}
