Resolving Sleeping barber problem with go
-----------------------------------------

The problem is simple and exposed there: http://en.wikipedia.org/wiki/Sleeping_barber_problem

This is very simple:

- a barber manage a shop with a waiting room
- a waiting room has got some seat
- some clients enter the shop, try to find a free seat

  - free seat: ok, wait for barber
  - no seat: go out... (and retry later if you want)

- if the barber has no clients, he decide to sleep
- if client enter and see that the barber is sleeping, wake him

For the resolution, and to be more expressive, I decided to let Clients that don't find a seat to retry later. It's easy to remove this loop in "client.go" file

In many langages, this problem can be solved by using mutex. But in Go the problem is a bit easier, we can use:

- channels
- select statement


Test the application
--------------------

To run the sample, you'll need Go 1.x and launch in a terminal::

        go run sleepingbarber.go

You can compile the example and give binary to other computer even if you don't have install Go runtime on those computers::

        go build sleepingbarber.go
        # launch binary:
        ./sleepingbarber

The binary is a static ELF binary. Remeber that you can't launch a binary on 32 bits system if you compiled binary with 64 bits computer.

Implementation
--------------

Yes, we could implement the problem and solution very quickly, but to create a readable code, I decided to write:

- a package
- struct for the shop that have a waiting room (channel) and a pointer on a barber
- barber that have a method to manage a shop
- client that can enter the shop

What does Client:

- client tries to write on waitingroom channel of the shop
- if ok, he tries to write on the "wake" channel of barber:

  - no write ? because barber is awake
  - write ok ? barber was sleeping

What does Barber:

- barber tries to read waitingroom channels

  - read ok ? there was at least one client
  - no read: let's sleep => read the wake channel until someone write inside

The "more than necessary"
-------------------------

As you can see... the sleep state (wake channel) of barber is not necessary in Go, because the select statement is enought to wait client. We can be sure that if no client is in waiting room, barber is sleeping (default case). But to implement the real problematic specified, I implemented the wake functionnality.


