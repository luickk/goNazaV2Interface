# Go DJI NazaV2 Interface

The goNazaV2Interface is an interface to control the Naza V2 flight controler. The interface enables a non licenced and free computer interface for the Naza V2 flight controller! This is crucial for open source development and for non industrie developers. The raspberry communicates with the Naza on 5 channels over [PWM](https://en.wikipedia.org/wiki/Pulse-width_modulation). Since the Raspberry does not have enough harware pwm ports, the pca9685 is used to create stable/ consistent and precise pwm signals. 

## Chosen Interface

The Naza V2 can adapt to different controll interfaces PWM, PPM and S-Bus are possible. The projects uses PWM, since that is the simplest, most reliable and easiest to emulate. The PWM input signal is 50 Hz since that is the most common frequency used by hobby receivers.

## Setup

1. `cd ./go/src`
2. `git clone https://github.com/cy8berpunk/goNazaV2Interface/` 

### Figuring out the limit(max,mid,min) values

To do that use the `tools/testDutyCycle.go` tool and try dutyCycles in the range of 0 to 500 for every channel and check which dutyCycle fits to the stick center, max. left/right position. Note the values for each channel and set them accordingly in the pkg interface Config struct. You can ignore the center position for the throttle stick.

Subsequently the functionality of the library can be tested with the `tools/testInterface.go` tool.

### (optional) Calibration

The calibration process is not required but if the drone is primarily used with this interface, it improves the reliability and accuracy.
To do so, just set your interfaceConfig values in the `tools/testInterface` to yours and activate the calibration mode in the Naza V2 software, afterwards start the pkg tool.

## PWM signal emulation

The go package provides 5 functions to control each axis/ channel of the drone including the flight mode channel. The translation takes the stick direction, which is a option in the Naza V2 calibration, into account and can be either "normal" or "reverse". Where by "reverse", as the name implies, reverses the pulse length for the max. and min.. The translation is done by taking the percentage value input of the pkg function and translate it to stick movement like PWM signals. The maximum, minimum and center values have to be set beforehand.
