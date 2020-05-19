package main

import(
  naza "goNazaV2Interface/goNazaV2Interface"
)

func main() {
  var interfaceConf naza.InterfaceConfig
  interfaceConf.StickDir[Achannel] = "rev"
  interfaceConf.StickDir[Echannel] = "norm"
  interfaceConf.StickDir[Tchannel] = "norm"
  interfaceConf.StickDir[Rchannel] = "rev"


  // key: channel, value: stick max pos
  interfaceConf.LeftStickMaxPos[Achannel] = 0
  interfaceConf.RightStickMaxPos[Achannel] = 0
  interfaceConf.NeutralStickPos[Achannel] = 0

  interfaceConf.LeftStickMaxPos[Echannel] = 0
  interfaceConf.RightStickMaxPos[Echannel] = 0
  interfaceConf.NeutralStickPos[Echannel] = 0

  interfaceConf.LeftStickMaxPos[Tchannel] = 0
  interfaceConf.RightStickMaxPos[Tchannel] = 0
  interfaceConf.NeutralStickPos[Tchannel] = 0

  interfaceConf.LeftStickMaxPos[Rchannel] = 0
  interfaceConf.RightStickMaxPos[Rchannel] = 0
  interfaceConf.NeutralStickPos[Rchannel] = 0

  interfaceConf.GpsModeFlipSwitchDutyCycle = 390
  interfaceConf.FailsafeModeFlipSwitchDutyCycle = 350
  interfaceConf.SelectableModeFlipSwitchDutyCycle = 250

  if naza.InitPCA9685(&interfaceConf) {
    if naza.InitNaza(&interfaceConf) {

    } else {
      println("failed to init naza")
    }
  } else {
    println("failed to init PCA9685")
  }
}
