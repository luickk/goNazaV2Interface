package main


import(
  "fmt"
  "strconv"
  naza "goNazaV2Interface/goNazaV2Interface"
)


func main() {
  var interfaceConf naza.InterfaceConfig

  interfaceConf.StickDir = make(map[int]string)

  interfaceConf.StickDir[naza.Achannel] = "rev"
  interfaceConf.StickDir[naza.Echannel] = "norm"
  interfaceConf.StickDir[naza.Tchannel] = "norm"
  interfaceConf.StickDir[naza.Rchannel] = "rev"


  interfaceConf.LeftStickMaxPos = make(map[int]int)
  interfaceConf.NeutralStickPos = make(map[int]int)
  interfaceConf.RightStickMaxPos = make(map[int]int)

  // key: channel, value: stick max pos
  interfaceConf.LeftStickMaxPos[naza.Achannel] = 400
  interfaceConf.NeutralStickPos[naza.Achannel] = 315
  interfaceConf.RightStickMaxPos[naza.Achannel] = 220

  interfaceConf.LeftStickMaxPos[naza.Echannel] = 400
  interfaceConf.NeutralStickPos[naza.Echannel] = 315
  interfaceConf.RightStickMaxPos[naza.Echannel] = 220

  interfaceConf.LeftStickMaxPos[naza.Tchannel] = 400
  interfaceConf.NeutralStickPos[naza.Tchannel] = 0
  interfaceConf.RightStickMaxPos[naza.Tchannel] = 230

  interfaceConf.LeftStickMaxPos[naza.Rchannel] = 400
  interfaceConf.NeutralStickPos[naza.Rchannel] = 315
  interfaceConf.RightStickMaxPos[naza.Rchannel] = 220

  interfaceConf.GpsModeFlipSwitchDutyCycle = 390
  interfaceConf.FailsafeModeFlipSwitchDutyCycle = 350
  interfaceConf.SelectableModeFlipSwitchDutyCycle = 250

  if naza.InitPCA9685(&interfaceConf) {
    var axis int
    var val int

    for {
      fmt.Println("chose axis")
      fmt.Println(" 1 A: for roll control \n 2 E: for pitch control \n 3 T: for throttle control \n 4 R: for yaw control \n 5 U: for mode control")
      fmt.Print("axis: ")
      fmt.Scan(&axis)
      println("Axis set to: " + strconv.Itoa(axis))
      print("val: ")
      fmt.Scan(&val)
      if axis == 1 {
        naza.SetRoll(&interfaceConf, val)
        println("Set Roll to: " + strconv.Itoa(val))
      } else if axis == 2 {
        naza.SetPitch(&interfaceConf, val)
        println("Set Pitch to: " + strconv.Itoa(val))
      } else if axis == 3 {
        naza.SetThrottle(&interfaceConf, val)
        println("Set Throttle to: " + strconv.Itoa(val))
      } else if axis == 4 {
        naza.SetYaw(&interfaceConf, val)
        println("Set Yaw to: " + strconv.Itoa(val))
      } else if axis == 5 {
        // naza.SetFlightMode(interfaceConf, val)
      }
    }
  } else {
    println("failed to init PCA9685")
  }
}
