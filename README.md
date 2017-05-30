<h1 align="center">Presento</h1> 

<p align="center">The simplest possible cross-platform remote control for the presentations</p>

<p align="center"><img src="https://github.com/alxrm/presento/blob/master/art/flow.gif?raw=true" alt="Example"></p>

## Setup

__NOTE: If you simply need the working product, just jump into [builds](https://github.com/alxrm/presento/tree/master/build/snapshot) directory and choose one of the executable binaries__

For those who want to build it manually. Presento is written in Go so you'll need to [install Go first](https://golang.org/dl/). Once that's done, you can build the app. First get the source:

```bash
$ go get -v github.com/alxrm/presento
```

Then go to the directory and run:

```bash
$ go build .
```

This should generate a `presento` binary.

Or if you use [goxc](https://github.com/laher/goxc), sure, why not, but check the [Contribution](https://github.com/alxrm/presento#contributing) chapter first


## Usage

_The only requirement is your cellphone has to be in the same Wi-fi network as your laptop, it simply would not work otherwise_

As you can(or cannot) see on the GIF above, the flow is pretty simple:

1) You launch the executable file, it will open the command line app with this text in it:

```
  Go to http://192.168.0.**:5000/**** to control
```

The link is generated with the 4-letter random sequence in the end

This was made because there might be those cases when someone sitting in the room might know the local address and try to spoil your presentation, so this should prevent such situations from happening

2) Open the presentation on the fullscreen

3) The link should be opened in your cellphone's browser, and by clicking on those huge buttons you can simulate keyboard press on the left/right buttons on the laptop that is showing the presentation thus changing the slides

4) Look cool 

## Reason

So most of the time, when I need to show a presentation it goes like this:
  - I launch it from my laptop, which is connected to the projector
  - I run to my laptop to change the slide (repeat this part)
  
Then this thought came:
  
_"Well it would be nice to have the ability to change the slides, like all those guys I usually see on the conferences"_


## Contribution

Presento has some problems with building for Linux with [goxc](https://github.com/laher/goxc) on the OS X, whilst it seems to be fine for the Windows/OS X. The problem comes from this awesome [library](https://github.com/micmonay/keybd_event), so if you can help it, you're more than welcome!

## License

MIT License. See the [LICENSE](LICENSE) file for more information.
