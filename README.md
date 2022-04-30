# go-robot-control

## Introduction

Here, Golang is used to control a robot along with the CamJam Edukit 3. The code uses keyboard control and requires root privileges. 

## Getting Started

First, install Go on your Robot. You can download from [here](https://go.dev/dl/). Next, extract the contents and move the directory to the `/usr/local` directory in order for Go to work properly. Next, make sure to add this line to the end of your `/etc/profile` file: `export PATH=$PATH:/usr/local/go/bin`.Then run `source /etc/profile`, log out or reboot and then run `go version` to make sure that everything is running well. You should get a version listed back. Also because we have packages that aren't available for us to use we will have to install these packages once we have go installed:

```
go get github.com/MarinX/keylogger
go get github.com/sirupsen/logrus
go get github.com/stianeikeland/go-rpio/v4
```

## Running the robot

After checking that Go is installed, you should first run `go mod init go-robot-control` to create the correct mod files. After that then you can build the application with `go build` and then because we require root privileges run `sudo ./go-robot-control` and then make sure your keyboard is plugged in. Then the LED should blink four times and then you can start pressing the arrow keys to move the robot. And that's it.
