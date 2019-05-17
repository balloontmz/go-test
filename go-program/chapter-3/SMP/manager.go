package mlib

import (
	"errors"
)

//MusicEntry 音乐对象 
type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

//MusicManager 音乐管理对象
type MusicManager struct {
	musics []MusicEntry
}

//NewMusicManager 新建音乐管理对象
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//Len 获取音乐库大小
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//Get 获取音乐库中某个音乐
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

//Find 获取音乐库中某个音乐
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

//Add 将某个音乐添加到音乐库中
func (m *MusicManager) Add(music *MusicEntry)  {
	m.musics = append(m.musics, *music)
}

//Remove 将某个音乐
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index <0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]

	m.musics = append(m.musics[:index], m.musics[index+1:]...)

	return removeMusic
}