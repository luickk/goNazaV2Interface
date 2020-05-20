package main


import(
  "fmt"
  "strconv"
  "log"
  "goNazaV2Interface/go-pca9685"
  i2c "goNazaV2Interface/go-i2c"
)


func main() {
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

  fmt.Print("channel: ")

  var ch int
  var val int
  fmt.Scan(&ch)

  println("Channel set to: " + strconv.Itoa(ch))

  for {
    fmt.Print("-> ")
    fmt.Scan(&val)

    pca0.SetChannel(ch, 0, val)
  }
}
