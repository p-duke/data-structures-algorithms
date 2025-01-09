/*
We are building a cloud-based music player, like Spotify.
Let's start with the following functionality:
 * `add_song (song_title [string])`
- A song is given an incremental integer ID when it's added, starting with 1
- Songs {id}
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
- Song { id, title, plays []int}
- every time we call play_song we want to track a play
- plays is an array of user ids
- User { id, }
Needed for summary
- sorted (desc) by unique plays (user id)
- include song title, number of plays (get from count of unique)
*/
class Song {
    constructor(options) {
        this.title = options.title;
        this.id = options.id;
        this.plays = [];
    }
}

class MusicPlayer {
    constructor() {
        this.songs = {}; // may change to an object
    }
    addSong(title) {
        let lastID;
        if (Object.keys(this.songs) === 0) {
            lastID = 0
        } else {
            const songs = Object.keys(this.songs);
            lastID = songs.length;
        }
        const newSong = new Song({ title: title, id:

            lastID+1 });

        this.songs[newSong.id] = newSong;
    }
    /*
- Assume any user ID is valid, and that the given song ID will have been added
- User ID should exist, and song ID should exist
- Need to be tracking users as well
*/
    // song id, user id
    playSong(songId, userId) {
        const song = this.songs[songId];
        if (song.plays.includes(userId)) {
            // console.log("User already played, skip tracking");
            return;
        }
        this.songs[songId].plays.push(userId);
    }
    /*
     * `print_analytics_summary ()`
- The summary should be sorted (descending) by the number of unique users
who have played each song. (tracking number of plays by user)
- The summary should include the song titles, and the number of unique users,
but the formatting does not matter.

Needed for summary
- sorted (desc) by unique plays (user id)
- include song title, number of plays (get from count of unique)
*/
    printAnalyticsSummary() {
        // calculate unique plays
        const result = [];
        for (let song in this.songs) {
            const uniquePlays =
                this.songs[song].plays.length;

            const temp = { title: this.songs[song].title,

                uniquePlays: uniquePlays};
            result.push(temp);
        }
        // sort desc by uniquePlays
        result.sort((a, b) => {
            return b.uniquePlays - a.uniquePlays;
        });
        return result;
    }
}
function main() {
    const user1 = 12;
    const user2 = 5;
    const user3 = 28;
    const musicPlayer = new MusicPlayer();
    // musicPlayer.addSong('Thriller'); // id 1
    // musicPlayer.addSong('Happy Days'); // id 2
    // musicPlayer.addSong('Song 3'); // id 3
    // musicPlayer.playSong(1, user1);
    // musicPlayer.playSong(1, user1);
    // musicPlayer.playSong(1, user1);
    // musicPlayer.playSong(2, user2);
    // musicPlayer.playSong(2, user2);
    // musicPlayer.playSong(3, user1);
    // musicPlayer.playSong(3, user2);
    // musicPlayer.playSong(3, user3);

    // test
    musicPlayer.addSong("Hello, Goodbye");
    musicPlayer.addSong("Bohemian Rhapsody");
    musicPlayer.addSong("Stairway to Heaven");
    musicPlayer.addSong("Satisfaction");
    musicPlayer.addSong("Pinball Wizard");

    musicPlayer.playSong(1, 9);
    musicPlayer.playSong(2, 14);
    musicPlayer.playSong(1, 2);
    musicPlayer.playSong(1, 1);
    musicPlayer.playSong(2, 1);
    musicPlayer.playSong(3, 17);
    musicPlayer.playSong(2, 1);
    musicPlayer.playSong(3, 5);
    musicPlayer.playSong(2, 1);
    musicPlayer.playSong(2, 1);
    musicPlayer.playSong(1, 7);
    musicPlayer.playSong(4, 1);
    musicPlayer.playSong(2, 1);
    musicPlayer.playSong(1, 1);
    musicPlayer.playSong(1, 1);
    musicPlayer.playSong(3, 1);
    musicPlayer.playSong(1, 1);

    console.log(musicPlayer.printAnalyticsSummary());
    // console.log("Music Player songs", musicPlayer.songs);
}

main();
