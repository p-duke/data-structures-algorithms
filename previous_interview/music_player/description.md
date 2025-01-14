## Music Player

We are building a cloud-based music player, like Spotify.

Let's start with the following functionality:

- `add_song(song_title string)`
    - A song is given an incremental integer ID when it's added, starting with 1

- `play_song(song_ID integer, user_ID integer)`
    - Assume any user ID is valid, and that the given song ID will have been added
    - User ID should exist, and song ID should exist

* `print_analytics_summary()`
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
