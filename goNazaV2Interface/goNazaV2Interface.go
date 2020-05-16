package goNazaV2Interface

import(

)

type interfaceConfig struct {
  stickDir string
  percentageValue int
  leftStickMaxPos int
  rightStickMaxPos int
  neutralStickPos int
}

func setPitch(iC *interfaceConfig, degreeVal int) int {
  return 0
}


func setRoll(iC *interfaceConfig, axis string) int {
  return 0
}

func serYaw(iC *interfaceConfig, axis string) int {
  return 0
}

func setThrottle(iC *interfaceConfig, axis string) int {
  return 0
}

// calculates pwm pulse. expects neutral stick position to be in the middle
func calcPWMPulseFromNeutralCenter(iC *interfaceConfig, side string, degreeVal int) int{
  var pwmPulse = 0
  if side == "left" {
    if iC.stickDir=="rev" {
      if iC.neutralStickPos>iC.leftStickMaxPos {
        pwmPulse=iC.neutralStickPos-iC.leftStickMaxPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.neutralStickPos
      } else if iC.leftStickMaxPos>iC.neutralStickPos {
        pwmPulse=iC.leftStickMaxPos-iC.neutralStickPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.neutralStickPos
      }
    } else if iC.stickDir=="norm" {
      if iC.neutralStickPos<iC.leftStickMaxPos {
        pwmPulse=iC.neutralStickPos-iC.leftStickMaxPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.neutralStickPos
      } else if iC.leftStickMaxPos<iC.neutralStickPos {
        pwmPulse=iC.leftStickMaxPos-iC.neutralStickPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=pwmPulse+iC.neutralStickPos
      }
    }
  } else if side == "right" {
    if iC.stickDir=="rev" {
      if(iC.neutralStickPos>iC.rightStickMaxPos){
        pwmPulse=iC.neutralStickPos-iC.rightStickMaxPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.neutralStickPos-pwmPulse
      } else if(iC.rightStickMaxPos>iC.neutralStickPos){
        pwmPulse=iC.rightStickMaxPos-iC.neutralStickPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.neutralStickPos-pwmPulse
      }
    } else if iC.stickDir=="norm" {

      if iC.neutralStickPos<iC.rightStickMaxPos {
        pwmPulse=iC.neutralStickPos-iC.rightStickMaxPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.neutralStickPos-pwmPulse
      } else if iC.rightStickMaxPos<iC.neutralStickPos {
        pwmPulse=iC.rightStickMaxPos-iC.neutralStickPos
        pwmPulse=pwmPulse/100
        pwmPulse=pwmPulse*degreeVal
        pwmPulse=iC.neutralStickPos-pwmPulse
      }
    }
  }
  return pwmPulse
}

func calcPWMPulseFromNeutralZero(iC *interfaceConfig, degreeVal int) int{
  var pwmPulse = 0
  if iC.stickDir=="rev" {
    if iC.rightStickMaxPos>iC.leftStickMaxPos {
      pwmPulse=iC.rightStickMaxPos-iC.leftStickMaxPos
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.leftStickMaxPos-pwmPulse
    } else if iC.rightStickMaxPos<iC.leftStickMaxPos {
      pwmPulse=iC.leftStickMaxPos-iC.rightStickMaxPos
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.leftStickMaxPos-pwmPulse
    }
  } else if(iC.stickDir=="norm"){
    if iC.rightStickMaxPos<iC.leftStickMaxPos {
      pwmPulse=iC.leftStickMaxPos-iC.rightStickMaxPos
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.rightStickMaxPos+pwmPulse
    } else if iC.rightStickMaxPos>iC.leftStickMaxPos {
      pwmPulse=iC.rightStickMaxPos-iC.leftStickMaxPos
      pwmPulse=pwmPulse/100
      pwmPulse=pwmPulse*degreeVal
      pwmPulse=iC.rightStickMaxPos-pwmPulse
    }
  }
  return pwmPulse
}
