# SMS - A concurrent Sega Master System emulator

SMS is a free and open-source (MIT licensed) Sega Master System emulator written in
[Go](http://golang.org). I think it's the first Sega Master System emulator written
in this language.

# Quick start

Installing and starting SMS with Go is simple:

    go get -v github.com/rafaelmdcarneiro/sms_go/
    ./sms roms/Sonic The Hedgehog (USA, Europe).sms

# Features

* Complete Zilog Z80 emulation
* Concurrent
* SDL backend
* 2x scaler and fullscreen

# Todo

* Sound support
* Write more tests

# Key bindings

    Host computer   Sega Master System
    ----------------------------------
    Arrows          Joypad directions
    X               Fire 1
    Z               Fire 2

For more info about key bindings see file <tt>input.go</tt>
