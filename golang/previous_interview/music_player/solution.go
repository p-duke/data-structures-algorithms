package main

import (
	"fmt"
	"sort"
)

/*
We are building a cloud-based music player, like Spotify.

Let's start with the following functionality:
* `add_song (song_title [string])`
- A song is given an incremental integer ID when it's added, starting with 1

* `play_song (song_ID [integer], user_ID [integer])`
- Assume any user ID is valid, and that the given song ID will have been added
- User ID should exist, and song ID should exist
- Need to be tracking users as well

* `print_analytics_summary ()`
- This is used for a report, created once per day for our company's Analytics
department.
- The summary should be sorted (descending) by the number of unique users
who have played each song. (tracking number of plays by user)
- The summary should include the song titles, and the number of unique users,
but the formatting does not matter.

For our MVP, consider performance as we will eventually support millions of songs
and users. (good runtimes matter)

However, let's not worry about thread safety or persistence for now - store data in
memory. (store data in memory)

Domain
- MusicPlayer
- Song
- User

Needed for summary
- sorted (desc) by unique plays (user id)
- include song title, number of plays (get from count of unique)
*/
type Song struct {
	id    int
	title string
	plays []int // userIds
}

type MusicPlayer struct {
	songs map[int]*Song
}

func (mp *MusicPlayer) addSong(title string) {
	if len(mp.songs) > 0 {
		lastSong := len(mp.songs)
		newId := lastSong + 1
		mp.songs[newId] = &Song{title: title, id: newId, plays: make([]int, 0)}
	} else {
		mp.songs = make(map[int]*Song)
		mp.songs[1] = &Song{title: title, id: 1, plays: make([]int, 0)}
	}
}

func (mp *MusicPlayer) playSong(songId, userId int) {
	song := mp.songs[songId]
	userPreviousPlayed := false

	if len(song.plays) > 0 {
		for _, user := range song.plays {
			if user == userId {
				userPreviousPlayed = true
			}
		}
	}

	if !userPreviousPlayed {
		song.plays = append(song.plays, userId)
	}

	// fmt.Printf("song: %+v", song)
}

type data struct {
	title string
	plays int
}

func (mp *MusicPlayer) printAnalyticsSummary() []data {
	// - The summary should be sorted (descending) by the number of unique users
	// who have played each song. (tracking number of plays by user)
	// - The summary should include the song titles, and the number of unique users,
	// but the formatting does not matter.

	// Todo
	// - declare a result data structure - song title, number of unique plays - map[string]int
	result := make([]data, 0)

	// - loop over songs
	for _, song := range mp.songs {
		result = append(result, data{title: song.title, plays: len(song.plays)})
	}

	// - sort the result
	sort.Slice(result, func(i, j int) bool {
		return result[i].plays > result[j].plays
	})

	return result
}

func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{}
}

func main() {
	mp := NewMusicPlayer()

	mp.addSong("Hello, Goodbye")
	mp.addSong("Bohemian Rhapsody")
	mp.addSong("Stairway to Heaven")
	mp.addSong("Satisfaction")
	mp.addSong("Pinball Wizard")

	mp.playSong(1, 9)
	mp.playSong(2, 14)
	mp.playSong(1, 2)
	mp.playSong(1, 1)
	mp.playSong(2, 1)
	mp.playSong(3, 17)
	mp.playSong(2, 1)
	mp.playSong(3, 5)
	mp.playSong(2, 1)
	mp.playSong(2, 1)
	mp.playSong(1, 7)
	mp.playSong(4, 1)
	mp.playSong(2, 1)
	mp.playSong(1, 1)
	mp.playSong(1, 1)
	mp.playSong(3, 1)
	mp.playSong(1, 1)

	fmt.Println(mp.printAnalyticsSummary())
}
