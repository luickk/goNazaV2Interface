package main

import(
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
  interfaceConf.RightStickMaxPos[naza.Achannel] = 320

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
    naza.Recalibrate(&interfaceConf)
  } else {
    println("failed to init PCA9685")
  }
}
