package goNazaV2Interface

import(
  "time"
)

type interfaceConfig struct {
  StickDir string

  // key: channel, value: stick max pos
  LeftStickMaxPos map[int]int
  RightStickMaxPos map[int]int
  NeutralStickPos map[int]int

  GpsModeFlipSwitchPulse int
  FailsafeModeFlipSwitchPulse int
  SelectableModeFlipSwitchPulse int
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
	println( "U Channel " + string(Uchannel) + " PWM Values " + "GPS: " + string(iC.GpsModeFlipSwitchPulse) + " Failsafe: " + string(iC.FailsafeModeFlipSwitchPulse) + " Selectable: " + string(iC.SelectableModeFlipSwitchPulse))
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
			println("Setting flight mode: "+ mode)
		} else if mode=="failsafe" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "failsafe")));
			println("Setting flight mode: "+ mode)
		} else if mode=="selectable" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "selectable")));
			println("Setting flight mode: "+ mode)
		}
}

// degreeVal represents the distance the stick travels either downwards -100 to 0 or 0 to 100 upwards
// returns true if successful
func SetPitch(iC *interfaceConfig, degreeVal int) bool {
  pwmPulse := 0
  if degreeVal < 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Echannel, "left", (degreeVal*-1))
  } else if degreeVal > 0{
    pwmPulse= calcPWMPulseFromNeutralCenter(iC, Echannel, "right", degreeVal)
  } else if degreeVal == 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Echannel, "left", 0)
  }
  print(pwmPulse)
  return true
}


// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetRoll(iC *interfaceConfig, degreeVal int) bool {
  pwmPulse := 0
  if degreeVal < 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Achannel, "left", (degreeVal*-1))
  } else if degreeVal > 0{
    pwmPulse= calcPWMPulseFromNeutralCenter(iC, Achannel, "right", degreeVal)
  } else if degreeVal == 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Achannel, "left", 0)
  }
  print(pwmPulse)

  return true
}

// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetYaw(iC *interfaceConfig, degreeVal int) bool {
  pwmPulse := 0
  if degreeVal < 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Rchannel, "left", (degreeVal*-1))
  } else if degreeVal > 0{
    pwmPulse= calcPWMPulseFromNeutralCenter(iC, Rchannel, "right", degreeVal)
  } else if degreeVal == 0 {
    pwmPulse = calcPWMPulseFromNeutralCenter(iC, Rchannel, "left", 0)
  }

  print(pwmPulse)
  return true
}

// degreeVal represents the distance the stick travels from 0 to 100 up or down
// returns true if successful
func SetThrottle(iC *interfaceConfig, degreeVal int) bool {
  pwmPulse := 0
  if degreeVal > 0 {
    pwmPulse = calcPWMPulseFromNeutralZero(iC, Tchannel, degreeVal)
  } else if degreeVal == 0 {
    pwmPulse = calcPWMPulseFromNeutralZero(iC, Tchannel, 0)
  }

  print(pwmPulse)
  return true
}

// calculates pwm pulse. expects neutral stick position to be in the middle
func calcPWMPulseFromNeutralCenter(iC *interfaceConfig, channel int, side string, degreeVal int) int{
  var pwmPulse = 0
  if side == "left" {
    if iC.StickDir=="rev" {
      if iC.NeutralStickPos[channel]>iC.LeftStickMaxPos[channel] {
        pwmPulse=iC.NeutralStickPos[channel]-iC.LeftStickMaxPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.NeutralStickPos[channel]
      } else if iC.LeftStickMaxPos[channel]>iC.NeutralStickPos[channel] {
        pwmPulse=iC.LeftStickMaxPos[channel]-iC.NeutralStickPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.NeutralStickPos[channel]
      }
    } else if iC.StickDir=="norm" {
      if iC.NeutralStickPos[channel]<iC.LeftStickMaxPos[channel] {
        pwmPulse=iC.NeutralStickPos[channel]-iC.LeftStickMaxPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.NeutralStickPos[channel]
      } else if iC.LeftStickMaxPos[channel]<iC.NeutralStickPos[channel] {
        pwmPulse=iC.LeftStickMaxPos[channel]-iC.NeutralStickPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.NeutralStickPos[channel]
      }
    }
  } else if side == "right" {
    if iC.StickDir=="rev" {
      if(iC.NeutralStickPos[channel]>iC.RightStickMaxPos[channel]) {
        pwmPulse=iC.NeutralStickPos[channel]-iC.RightStickMaxPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.NeutralStickPos[channel]-pwmPulse
      } else if(iC.RightStickMaxPos[channel]>iC.NeutralStickPos[channel]) {
        pwmPulse=iC.RightStickMaxPos[channel]-iC.NeutralStickPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.NeutralStickPos[channel]-pwmPulse
      }
    } else if iC.StickDir=="norm" {

      if iC.NeutralStickPos[channel]<iC.RightStickMaxPos[channel] {
        pwmPulse=iC.NeutralStickPos[channel]-iC.RightStickMaxPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.NeutralStickPos[channel]-pwmPulse
      } else if iC.RightStickMaxPos[channel]<iC.NeutralStickPos[channel] {
        pwmPulse=iC.RightStickMaxPos[channel]-iC.NeutralStickPos[channel]
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.NeutralStickPos[channel]-pwmPulse
      }
    }
  }
  return pwmPulse
}

func calcPWMPulseFromNeutralZero(iC *interfaceConfig, channel int, degreeVal int) int{
  var pwmPulse = 0
  if iC.StickDir=="rev" {
    if iC.RightStickMaxPos[channel]>iC.LeftStickMaxPos[channel] {
      pwmPulse=iC.RightStickMaxPos[channel]-iC.LeftStickMaxPos[channel]
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.LeftStickMaxPos[channel]-pwmPulse
    } else if iC.RightStickMaxPos[channel]<iC.LeftStickMaxPos[channel] {
      pwmPulse=iC.LeftStickMaxPos[channel]-iC.RightStickMaxPos[channel]
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.LeftStickMaxPos[channel]-pwmPulse
    }
  } else if(iC.StickDir=="norm"){
    if iC.RightStickMaxPos[channel]<iC.LeftStickMaxPos[channel] {
      pwmPulse=iC.LeftStickMaxPos[channel]-iC.RightStickMaxPos[channel]
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.RightStickMaxPos[channel]+pwmPulse
    } else if iC.RightStickMaxPos[channel]>iC.LeftStickMaxPos[channel] {
      pwmPulse=iC.RightStickMaxPos[channel]-iC.LeftStickMaxPos[channel]
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.RightStickMaxPos[channel]-pwmPulse
    }
  }
  return pwmPulse
}
