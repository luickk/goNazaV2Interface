package goNazaV2Interface

import(
  "time"
  "goNazaV2Interface/go-pca9685"
)

type interfaceConfig struct {
  StickDir string
  pca *PCA9685

  // key: channel, value: stick max pos
  LeftStickMaxPos map[int]int
  RightStickMaxPos map[int]int
  NeutralStickPos map[int]int

  GpsModeFlipSwitchDutyCycle int
  FailsafeModeFlipSwitchDutyCycle int
  SelectableModeFlipSwitchDutyCycle int
}

// pca9685 channel enums
const(
  // A: for roll control
  Achannel = 1
  // E: for pitch control
  Echannel = 2
  // T: for throttle control
  Tchannel = 3
  // R: for yaw control
  Rchannel = 4
  // U: for mode control
  Uchannel = 5
)

func InitI2CPca9685() *PCA9685 {
  // Create new connection to i2c-bus on 1 line with address 0x40.
  // Use i2cdetect utility to find device address over the i2c-bus
  i2c, err := i2c.NewI2C(pca9685.Address, 1)
  if err != nil {
      log.Fatal(err)
  }

  pca0 := pca9685.PCANew(i2c, nil)
  err = pca0.Init()
  if err != nil {
    log.Fatal(err)
  }

  return pca0
}

/**
    init_naza requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    init_naza sets all control elements to neutral and prints pwm config values.
*/
func InitNaza(iC *interfaceConfig){
	println( "Initializing Naza Interface Controller")
	println( "-" )
	println( "A Channel " + string(Achannel) + " PWM Values " + "left: " + string(iC.LeftStickMaxPos[Achannel]) + " middle: " + string(iC.NeutralStickPos[Achannel]) + " right: " + string(iC.RightStickMaxPos[Achannel]))
	println( "E Channel " + string(Echannel) + " PWM Values " + "left: " + string(iC.LeftStickMaxPos[Echannel]) + " middle: " + string(iC.NeutralStickPos[Echannel]) + " right: " + string(iC.RightStickMaxPos[Echannel]))
	println( "T Channel " + string(Tchannel) + " PWM Values " + "left: " + string(iC.LeftStickMaxPos[Tchannel]) + " right: " + string(iC.RightStickMaxPos[Tchannel]))
	println( "R Channel " + string(Rchannel) + " PWM Values " + "left: " + string(iC.LeftStickMaxPos[Rchannel]) + " middle: " + string(iC.NeutralStickPos[Rchannel]) + " right: " + string(iC.RightStickMaxPos[Rchannel]))
	println( "U Channel " + string(Uchannel) + " PWM Values " + "GPS: " + string(iC.GpsModeFlipSwitchDutyCycle) + " Failsafe: " + string(iC.FailsafeModeFlipSwitchDutyCycle) + " Selectable: " + string(iC.SelectableModeFlipSwitchDutyCycle))
	println( "-" )

	SetNeutral(iC);
	println( "Setting all channels on neutral!" )
	println( "---IMPORTANT---" )
	println( "   Check channels for command calibration before giving power to motors!" )
}

/**
    arm_motors requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    arm_motors arms the motors and then returns all channels to neutral
*/
func ArmMotors(iC *interfaceConfig) {
	println("-------- ARMING MOTORS --------")
  SetRoll(iC, -100)
  SetPitch(iC, -100)
  SetThrottle(iC, 0)
  SetYaw(iC, -100)

	time.Sleep(2 * time.Second);

	SetNeutral(iC);
}

/**
    set_neutral requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    set_neutral sets all channels to neutral including THROTTLE TO 0%
*/
func SetNeutral(iC *interfaceConfig) {
		if (SetPitch(iC, 0) && SetRoll(iC, 0) && SetYaw(iC, 0) && SetThrottle(iC, 0)) {
      println("set stick pos to neutral")
    } else {
      println("failed to set stick pos to neutral")
    }
}

//  Use set_flight_mode to switches between different flight modes.
func SetFlightMode(iC *interfaceConfig, mode string) {
    if mode=="gps" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "gps")));
      iC.SetChannel(Uchannel, 0, GpsModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		} else if mode=="failsafe" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "failsafe")));
      iC.SetChannel(Uchannel, 0, FailsafeModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		} else if mode=="selectable" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "selectable")));
      iC.SetChannel(Uchannel, 0, SelectableModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		}
}

// degreeVal represents the distance the stick travels either downwards -100 to 0 or 0 to 100 upwards
// returns true if successful
func SetPitch(iC *interfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Echannel, "left", (degreeVal*-1))
    iC.SetChannel(Echannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Echannel, "right", degreeVal)
    iC.SetChannel(Echannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Echannel, "left", 0)
    iC.SetChannel(Echannel, 0, dutyCycle)
  }
  print(dutyCycle)
  return true
}


// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetRoll(iC *interfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Achannel, "left", (degreeVal*-1))
    iC.SetChannel(Achannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Achannel, "right", degreeVal)
    iC.SetChannel(Achannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Achannel, "left", 0)
    iC.SetChannel(Achannel, 0, dutyCycle)
  }
  print(dutyCycle)

  return true
}

// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetYaw(iC *interfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Rchannel, "left", (degreeVal*-1))
    iC.SetChannel(Rchannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Rchannel, "right", degreeVal)
    iC.SetChannel(Rchannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Rchannel, "left", 0)
    iC.SetChannel(Rchannel, 0, dutyCycle)
  }

  print(dutyCycle)
  return true
}

// degreeVal represents the distance the stick travels from 0 to 100 up or down
// returns true if successful
func SetThrottle(iC *interfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal > 0 {
    dutyCycle = calcdutyCycleFromNeutralZero(iC, Tchannel, degreeVal)
    iC.SetChannel(Tchannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralZero(iC, Tchannel, 0)
    iC.SetChannel(Tchannel, 0, dutyCycle)
  }

  print(dutyCycle)
  return true
}

// calculates pwm pulse. expects neutral stick position to be in the middle
func calcdutyCycleFromNeutralCenter(iC *interfaceConfig, channel int, side string, degreeVal int) int{
  var dutyCycle = 0
  if side == "left" {
    if iC.StickDir=="rev" {
      if iC.NeutralStickPos[channel]>iC.LeftStickMaxPos[channel] {
        dutyCycle=iC.NeutralStickPos[channel]-iC.LeftStickMaxPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=dutyCycle+iC.NeutralStickPos[channel]
      } else if iC.LeftStickMaxPos[channel]>iC.NeutralStickPos[channel] {
        dutyCycle=iC.LeftStickMaxPos[channel]-iC.NeutralStickPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=dutyCycle+iC.NeutralStickPos[channel]
      }
    } else if iC.StickDir=="norm" {
      if iC.NeutralStickPos[channel]<iC.LeftStickMaxPos[channel] {
        dutyCycle=iC.NeutralStickPos[channel]-iC.LeftStickMaxPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=dutyCycle+iC.NeutralStickPos[channel]
      } else if iC.LeftStickMaxPos[channel]<iC.NeutralStickPos[channel] {
        dutyCycle=iC.LeftStickMaxPos[channel]-iC.NeutralStickPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=dutyCycle+iC.NeutralStickPos[channel]
      }
    }
  } else if side == "right" {
    if iC.StickDir=="rev" {
      if(iC.NeutralStickPos[channel]>iC.RightStickMaxPos[channel]) {
        dutyCycle=iC.NeutralStickPos[channel]-iC.RightStickMaxPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=iC.NeutralStickPos[channel]-dutyCycle
      } else if(iC.RightStickMaxPos[channel]>iC.NeutralStickPos[channel]) {
        dutyCycle=iC.RightStickMaxPos[channel]-iC.NeutralStickPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=iC.NeutralStickPos[channel]-dutyCycle
      }
    } else if iC.StickDir=="norm" {

      if iC.NeutralStickPos[channel]<iC.RightStickMaxPos[channel] {
        dutyCycle=iC.NeutralStickPos[channel]-iC.RightStickMaxPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=iC.NeutralStickPos[channel]-dutyCycle
      } else if iC.RightStickMaxPos[channel]<iC.NeutralStickPos[channel] {
        dutyCycle=iC.RightStickMaxPos[channel]-iC.NeutralStickPos[channel]
        dutyCycle=dutyCycle/100
        dutyCycle=dutyCycle*degreeVal
        dutyCycle=iC.NeutralStickPos[channel]-dutyCycle
      }
    }
  }
  return dutyCycle
}

func calcdutyCycleFromNeutralZero(iC *interfaceConfig, channel int, degreeVal int) int{
  var dutyCycle = 0
  if iC.StickDir=="rev" {
    if iC.RightStickMaxPos[channel]>iC.LeftStickMaxPos[channel] {
      dutyCycle=iC.RightStickMaxPos[channel]-iC.LeftStickMaxPos[channel]
      dutyCycle=dutyCycle/100
      dutyCycle=dutyCycle*degreeVal
      dutyCycle=iC.LeftStickMaxPos[channel]-dutyCycle
    } else if iC.RightStickMaxPos[channel]<iC.LeftStickMaxPos[channel] {
      dutyCycle=iC.LeftStickMaxPos[channel]-iC.RightStickMaxPos[channel]
      dutyCycle=dutyCycle/100
      dutyCycle=dutyCycle*degreeVal
      dutyCycle=iC.LeftStickMaxPos[channel]-dutyCycle
    }
  } else if(iC.StickDir=="norm"){
    if iC.RightStickMaxPos[channel]<iC.LeftStickMaxPos[channel] {
      dutyCycle=iC.LeftStickMaxPos[channel]-iC.RightStickMaxPos[channel]
      dutyCycle=dutyCycle/100
      dutyCycle=dutyCycle*degreeVal
      dutyCycle=iC.RightStickMaxPos[channel]+dutyCycle
    } else if iC.RightStickMaxPos[channel]>iC.LeftStickMaxPos[channel] {
      dutyCycle=iC.RightStickMaxPos[channel]-iC.LeftStickMaxPos[channel]
      dutyCycle=dutyCycle/100
      dutyCycle=dutyCycle*degreeVal
      dutyCycle=iC.RightStickMaxPos[channel]-dutyCycle
    }
  }
  return dutyCycle
}
