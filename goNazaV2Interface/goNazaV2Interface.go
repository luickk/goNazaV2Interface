package goNazaV2Interface

import(
  "time"
  "log"
  "strconv"
  "goNazaV2Interface/go-pca9685"
  i2c "goNazaV2Interface/go-i2c"
)

type InterfaceConfig struct {
  pca *pca9685.PCA9685

  // key: channel, value: stick dir
  StickDir map[int]string

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
  Achannel = 0
  // E: for pitch control
  Echannel = 1
  // T: for throttle control
  Tchannel = 2
  // R: for yaw control
  Rchannel = 3
  // U: for mode control
  Uchannel = 4
)

// Create new PCA9685 conn
func InitPCA9685(iC *InterfaceConfig) bool{
  // Create new connection to i2c-bus on 1 line with address 0x40.
  // Use i2cdetect utility to find device address over the i2c-bus
  i2c, err := i2c.NewI2C(pca9685.Address, 1)
  if err != nil {
    log.Fatal(err)
    return false
  }

  pca0 := pca9685.PCANew(i2c, nil)
  err = pca0.Init()
  if err != nil {
    log.Fatal(err)
    return false
  }
  iC.pca = pca0
  return true
}


/**
    init_naza requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    init_naza sets all control elements to neutral and prints pwm config values.
*/
func InitNaza(iC *InterfaceConfig) bool{
  if !(iC.LeftStickMaxPos != nil || iC.RightStickMaxPos != nil || iC.NeutralStickPos != nil || iC.GpsModeFlipSwitchDutyCycle != 0 || iC.SelectableModeFlipSwitchDutyCycle != 0 || iC.FailsafeModeFlipSwitchDutyCycle != 0) {
    println("uninitialized val found")
    return false
  }
	println( "Initializing Naza Interface Controller")
	println( "-" )
	println( "A Channel " + strconv.Itoa(Achannel) + " PWM Values " + "left: " + strconv.Itoa(iC.LeftStickMaxPos[Achannel]) + " middle: " + strconv.Itoa(iC.NeutralStickPos[Achannel]) + " right: " + strconv.Itoa(iC.RightStickMaxPos[Achannel]))
	println( "E Channel " + strconv.Itoa(Echannel) + " PWM Values " + "left: " + strconv.Itoa(iC.LeftStickMaxPos[Echannel]) + " middle: " + strconv.Itoa(iC.NeutralStickPos[Echannel]) + " right: " + strconv.Itoa(iC.RightStickMaxPos[Echannel]))
	println( "T Channel " + strconv.Itoa(Tchannel) + " PWM Values " + "left: " + strconv.Itoa(iC.LeftStickMaxPos[Tchannel]) + " right: " + strconv.Itoa(iC.RightStickMaxPos[Tchannel]))
	println( "R Channel " + strconv.Itoa(Rchannel) + " PWM Values " + "left: " + strconv.Itoa(iC.LeftStickMaxPos[Rchannel]) + " middle: " + strconv.Itoa(iC.NeutralStickPos[Rchannel]) + " right: " + strconv.Itoa(iC.RightStickMaxPos[Rchannel]))
	println( "U Channel " + strconv.Itoa(Uchannel) + " PWM Values " + "GPS: " + strconv.Itoa(iC.GpsModeFlipSwitchDutyCycle) + " Failsafe: " + strconv.Itoa(iC.FailsafeModeFlipSwitchDutyCycle) + " Selectable: " + strconv.Itoa(iC.SelectableModeFlipSwitchDutyCycle))
	println( "-" )

	SetNeutral(iC);
	println( "Setting all channels on neutral!" )
	println( "---IMPORTANT---" )
	println( "   Check channels for command calibration before giving power to motors!" )

  return true
}

/**
    recalibrate requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
		recalibrate recalibrate all channels to corresponding config values.
*/
func Recalibrate(iC *InterfaceConfig) {

		println("Resetting all channels!")
		SetNeutral(iC)
		time.Sleep(1*time.Second)

		println("Recalibration of channel A (1/2)")
    iC.pca.SetChannel(Achannel, 0, iC.LeftStickMaxPos[Achannel])
		time.Sleep(1*time.Second);
		println("Recalibration of channel A (2/2)")
    iC.pca.SetChannel(Achannel, 0, iC.RightStickMaxPos[Achannel])

		time.Sleep(1*time.Second);
		println("Recalibration of channel E (1/2)")
    iC.pca.SetChannel(Echannel, 0, iC.LeftStickMaxPos[Echannel])
		time.Sleep(1*time.Second);
		println("Recalibration of channel E (2/2)")
    iC.pca.SetChannel(Echannel, 0, iC.RightStickMaxPos[Echannel])

		time.Sleep(1*time.Second);
		println("Recalibration of channel T (1/2)")
    iC.pca.SetChannel(Tchannel, 0, iC.LeftStickMaxPos[Tchannel])
		time.Sleep(1*time.Second);
		println("Recalibration of channel T (2/2)")
    iC.pca.SetChannel(Tchannel, 0, iC.RightStickMaxPos[Tchannel])

		time.Sleep(1*time.Second);
		println("Recalibration of channel R (1/2)")
    iC.pca.SetChannel(Rchannel, 0, iC.LeftStickMaxPos[Rchannel])
		time.Sleep(1*time.Second);
		println("Recalibration of channel R (2/2)")
    iC.pca.SetChannel(Rchannel, 0, iC.RightStickMaxPos[Rchannel])

 		time.Sleep(1*time.Second);
    println("Recalibration of channel U (1/2)")
    iC.pca.SetChannel(Uchannel, 0, iC.GpsModeFlipSwitchDutyCycle)
    time.Sleep(1*time.Second);
    println("Recalibration of channel U (2/2)")
    iC.pca.SetChannel(Uchannel, 0, iC.SelectableModeFlipSwitchDutyCycle)

    SetNeutral(iC)
}


/**
    arm_motors requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    arm_motors arms the motors and then returns all channels to neutral
*/
func ArmMotors(iC *InterfaceConfig) {
	println("-------- ARMING MOTORS --------")
  SetRoll(iC, -100)
  SetPitch(iC, -100)
  SetThrottle(iC, 0)
  SetYaw(iC, -100)

	time.Sleep(2 * time.Second)

	SetNeutral(iC)
}

/**
    set_neutral requires the drone NOT to be in the air!
		ONLY USE ON GROUND!
    set_neutral sets all channels to neutral including THROTTLE TO 0%
*/
func SetNeutral(iC *InterfaceConfig) {
		if (SetPitch(iC, 0) && SetRoll(iC, 0) && SetYaw(iC, 0) && SetThrottle(iC, 0)) {
      println("set stick pos to neutral")
    } else {
      println("failed to set stick pos to neutral")
    }
    SetFlightMode(iC, "gps")
}

//  Use set_flight_mode to switches between different flight modes.
func SetFlightMode(iC *InterfaceConfig, mode string) {
    if mode=="gps" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "gps")));
      iC.pca.SetChannel(Uchannel, 0, iC.GpsModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		} else if mode=="failsafe" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "failsafe")));
      iC.pca.SetChannel(Uchannel, 0, iC.FailsafeModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		} else if mode=="selectable" {
			// pca9685.Write(CHANNEL(iC.("U","channel")), VALUE(iC.("U", "selectable")));
      iC.pca.SetChannel(Uchannel, 0, iC.SelectableModeFlipSwitchDutyCycle)
			println("Setting flight mode: "+ mode)
		}
}

// degreeVal represents the distance the stick travels either downwards -100 to 0 or 0 to 100 upwards
// returns true if successful
func SetPitch(iC *InterfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Echannel, "left", (degreeVal*-1))
    iC.pca.SetChannel(Echannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Echannel, "right", degreeVal)
    iC.pca.SetChannel(Echannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Echannel, "left", 0)
    iC.pca.SetChannel(Echannel, 0, dutyCycle)
  }
  print(dutyCycle)
  return true
}


// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetRoll(iC *InterfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Achannel, "left", (degreeVal*-1))
    iC.pca.SetChannel(Achannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Achannel, "right", degreeVal)
    iC.pca.SetChannel(Achannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Achannel, "left", 0)
    iC.pca.SetChannel(Achannel, 0, dutyCycle)
  }
  print(dutyCycle)

  return true
}

// degreeVal represents the distance the stick travels either to the left -100 to 0 or 0 to 100 to the right
// returns true if successful
func SetYaw(iC *InterfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal < 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Rchannel, "left", (degreeVal*-1))
    iC.pca.SetChannel(Rchannel, 0, dutyCycle)
  } else if degreeVal > 0{
    dutyCycle= calcdutyCycleFromNeutralCenter(iC, Rchannel, "right", degreeVal)
    iC.pca.SetChannel(Rchannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralCenter(iC, Rchannel, "left", 0)
    iC.pca.SetChannel(Rchannel, 0, dutyCycle)
  }

  print(dutyCycle)
  return true
}

// degreeVal represents the distance the stick travels from 0 to 100 up or down
// returns true if successful
func SetThrottle(iC *InterfaceConfig, degreeVal int) bool {
  dutyCycle := 0
  if degreeVal > 0 {
    dutyCycle = calcdutyCycleFromNeutralZero(iC, Tchannel, degreeVal)
    iC.pca.SetChannel(Tchannel, 0, dutyCycle)
  } else if degreeVal == 0 {
    dutyCycle = calcdutyCycleFromNeutralZero(iC, Tchannel, 0)
    iC.pca.SetChannel(Tchannel, 0, dutyCycle)
  }

  print(dutyCycle)
  return true
}

// calculates pwm pulse. expects neutral stick position to be in the middle
func calcdutyCycleFromNeutralCenter(iC *InterfaceConfig, channel int, side string, degreeVal int) int{
  var dutyCycle = 0
  if side == "left" {
    if iC.StickDir[channel]=="rev" {
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
    } else if iC.StickDir[channel]=="norm" {
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
    if iC.StickDir[channel]=="rev" {
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
    } else if iC.StickDir[channel]=="norm" {

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

func calcdutyCycleFromNeutralZero(iC *InterfaceConfig, channel int, degreeVal int) int{
  var dutyCycle = 0
  if iC.StickDir[channel]=="rev" {
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
  } else if(iC.StickDir[channel]=="norm"){
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
