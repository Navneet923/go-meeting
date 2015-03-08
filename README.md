# go-meeting

A simple HTTP web service to help run faster meetings

## TODO:

~~1. Build basic Go http server~~
2. Provide route handlers for various meeting related actions:
    a. Create a new meeting (maybe a meeting can have various types)
    b. Maintain db of meetings, allow browsing and download of the same 
    c. View of pending action items by user
    d. More stuff I am surely forgetting
3. Store these meetings as .md / .pdf files
4. Provide an email w/ action items et al to each participant
5. TESTS!!!

## Install

```
$go get github.com/sabhiram/go-meeting
$go-meeting
```

## Deploy with Docker

Due to the awesome golang docker image, we can host this server as a container fairly trivially.

```
$docker build -t <user>/meeting-minutes .
$docker run -p 3000:3000 --name meeting-minutes --rm <user>/meeting-minutes
```

Visit your host machine's ip at port 3000 to view the service!

