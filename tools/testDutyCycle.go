package main


import(
  "fmt"
  "os"
  "strings"
  "strconv"
  "bufio"
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

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("channel: ")

  channel, _ := reader.ReadString('\n')

  println("Channel set to: " + channel)
  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    ch, _ := strconv.Atoi(channel)
    val, _ := strconv.Atoi(text)
    pca0.SetChannel(ch, 0, val)
  }
}
