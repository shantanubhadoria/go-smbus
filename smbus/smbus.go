package smbus

/*
#include<i2c-dev.h>
int prepareFile(int file){
  return ioctl(file,I2C_SMBUS);
} 
*/
import "C"

import (
  "unsafe"
  "os"
  "fmt"
)

type smbusDevice struct {
  deviceFilePath      string
  deviceFileHandle    int
  isFileInitialized   bool

  deviceAddress       int
}


// Function to return device file number if file handle has been initialized 
// else initialize and return device file handle

func (dev smbusDevice) getFile() int {
  if ! dev.isFileInitialized {
    dev.deviceFileHandle,err = os.Open(dev.deviceFilePath)
    dev.isFileInitialized = true
  }
  return dev.deviceFileNumber 
}

func (dev smbusDevice) writeQuick(value int) int {
  i2c_smbus_write_quick(, value);
}
