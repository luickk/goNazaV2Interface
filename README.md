# Go DJI NazaV2 Interface

The goNazaV2Interface is an interface to control the Naza V2 flight controler. The raspberry communicates with the Naza on 5 channels over [PWM](https://en.wikipedia.org/wiki/Pulse-width_modulation). Since the Raspberry does not have enough harware pwm ports, the pca9685 is used to create stable/ consistent and precise pwm signals. 

## Chosen Interface

The Naza V2 can adapt to different controll interfaces PWM, PPM and S-Bus are possible. The projects uses PWM, since that is the simplest, most reliable and easiest to emulate. The PWM input signal is 50 Hz since that is the most common frequency used by hobby receivers.
