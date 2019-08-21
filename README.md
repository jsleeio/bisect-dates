# bisect-dates

## TLDR

Given two dates, helps you bisect the time between them and efficiently find out
when _something_ happened.

Similar to `git bisect` but with no domain-specific nature. You are expected to
go away at each bisection point and check whatever it is for yourself, then come
back and give an appropriate response.


```
$ ./bisect-dates -ok-date=2018-12-01

ok    : 2018-12-01
bisect: 2018-12-21 <===
bad   : 2019-01-10

is this bisect point OK? [Ynq]:

...
```

The "bad" date is assumed to be after the "ok" date, and defaults to the
current date because this is probably what you want nearly all of the time.

## running it


```
$ bisect-dates -help
Usage of ./bisect-dates:
  -bad-date string
    	date that this thing was last OK (YYYY-MM-DD format, or 'today') (default "today")
  -ok-date string
    	date that this thing was last OK (YYYY-MM-DD format)
```

## docker

```
docker run gcr.io/jsleeio-containers/bisect-dates:latest
```
