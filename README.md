# Follow Sync

**Follow Sync** is a little Go program that helps you synchronize your
Instagram "Following" list. Basically, it can unfollow all profiles that are
not following you back.

It uses the Instagram unofficial API ([ahmdrz/goinsta][1]) to log in to your
account (pretending that it's an Android device), compares your followers and
following lists, and provides a summary and a choice of action: go ahead and
unfollow the users who don't follow you back, or exit and you can review the
data on your own.

This program creates a CSV file named `follower-lists.csv` that contains your
full list, so you can examine what's on your lists or as a backup in case you
want to know who you unfollowed later.

```
$ go install github.com/kirsle/follow-sync
$ follow-sync
```

# Program Usage

Just run the `follow-sync` executable. It will prompt for your username and
password (non-echoed output), collect your friend lists and compare them,
and print out a summary.

At this point you can open the `follower-lists.csv` in your favorite
spreadsheet program if you want to inspect the Following and Follower lists.

The program will prompt for final verification before it goes ahead and begins
unfollowing users.

# Caveats

Occasionally, the Instagram API can get pretty fussy with rate limits. Most
API functions this program calls will panic if it gets an error. No big deal:
you can just wait a few minutes and run the program again and it should pick
up where it left off.

To *try* and combat the rate limit error, the program intentionally waits
5 seconds between unfollows.

# License

```
Follow-Sync: Instagram Follower Synchronization Tool
Copyright (C) 2017 Noah Petherbridge

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
```
