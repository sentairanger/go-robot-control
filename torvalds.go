package main

import (
	
        "fmt"
        "os"
	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
	
)

var (

      pin1 = rpio.Pin(7)
      pin2 = rpio.Pin(8)
      pin3 = rpio.Pin(9)
      pin4 = rpio.Pin(10)
      eye = rpio.Pin(25)

)

func main() {

       // Check for errors
       if err := rpio.Open(); err != nil {
               fmt.Println(err)
               os.Exit(1)
       
       }
       // Defer gpio memory when done
       defer rpio.Close()
       
       // Set the pins to output mode
       pin1.Output()
       pin2.Output()
       pin3.Output()
       pin4.Output()
       eye.Output()
       
       // Blink the eye four times
       for x := 0; x < 4; x++ {
            eye.Toggle()
            time.Sleep(1 * time.Second)
       
       }

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed move the robot based on directions
			if e.KeyPress() {
				logrus.Println("[event] press key ", e.KeyString())
			    if e.KeyString() == "Up" {
			        forward()
			    
			    }
			    if e.KeyString() == "Down" {
			        backward()
			    
			    }
			    if e.KeyString() == "Left" {
			        left()
			    
			    
			    }
			    if e.KeyString() == "Right" {
			        right()
			    
			    
			    }
			
			}

			// if the state of key is released the robot will stop
			if e.KeyRelease() {
				logrus.Println("[event] release key ", e.KeyString())
		             if e.KeyString() == "Up" {
		             	stop()
		             
		             }
		             if e.KeyString() == "Down" {
		                stop()
		             
		             }
		             if e.KeyString() == "Left" {
		                stop()
		             
		             }
		             if e.KeyString() == "Right" {
		                stop()
		             
		             }		
		
				
			}

			break
		}
	}
}

// Functions for robot movement

func forward() {
     pin1.Low()
     pin2.High()
     pin3.High()
     pin4.Low()
}

func backward() {
    pin1.High()
    pin2.Low()
    pin3.Low()
    pin4.High()

}

func left() {
    pin1.High()
    pin2.Low()
    pin3.High()
    pin4.Low()



}

func right() {
    pin1.Low()
    pin2.High()
    pin3.Low()
    pin4.High()



}

func stop() {
    pin1.Low()
    pin2.Low()
    pin3.Low()
    pin4.Low()

}
