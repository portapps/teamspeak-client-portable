<p align="center"><a href="https://github.com/portapps/teamspeak-client-portable" target="_blank"><img width="100" src="https://github.com/portapps/teamspeak-client-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/teamspeak-client-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/teamspeak-client-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/teamspeak-client-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/teamspeak-client-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/teamspeak-client-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/teamspeak-client-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/teamspeak-client-portable"><img src="https://goreportcard.com/badge/github.com/portapps/teamspeak-client-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/teamspeak-client-portable"><img src="https://img.shields.io/codacy/grade/ed37fe9f437945079bd306ef1e871652.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://saythanks.io/to/crazymax"><img src="https://img.shields.io/badge/thank-crazymax-426aa5.svg?style=flat-square" alt="Say Thanks"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[TeamSpeak Client](https://www.teamspeak.com) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `teamspeak-client-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of TeamSpeak Client as a setup. **Recommended way**!
* `teamspeak-client-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of TeamSpeak Client as a 7z archive.
* `teamspeak-client-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `teamspeak-client-portable.exe`)
* `TeamSpeak-Client-{win32,win64}-x.x.x-setup.exe` : The original setup from the [official website](https://www.teamspeak.com/en/downloads.html#client).

### Fresh installation

Install `teamspeak-client-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `teamspeak-client-portable.exe`.

### App already installed

If you have already installed TeamSpeak Client from the original setup, do the same thing as a fresh installation and move :

* Move data located in `%APPDATA%\TS3Client\*` to `data` folder.

Run `teamspeak-client-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) TeamSpeak Client from your computer.

### Upgrade

For an upgrade, simply download and install the [latest setup](https://github.com/portapps/teamspeak-client-portable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

[![Donate Paypal](https://raw.githubusercontent.com/portapps/portapps/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
