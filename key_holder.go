package main

import (
  "time"
  "runtime"
  "github.com/micmonay/keybd_event"
)

type KeyHolder struct {
  left  keybd_event.KeyBonding
  right keybd_event.KeyBonding
}

func NewKeyHolder() *KeyHolder {
  controller := &KeyHolder{}

  left, errLeft := keybd_event.NewKeyBonding()
  right, errRight := keybd_event.NewKeyBonding()

  if errLeft != nil {
    panic(errLeft)
  } else if errRight != nil {
    panic(errRight)
  }

  controller.left = left
  controller.right = right

  // For linux, it is very important wait 2 seconds
  if runtime.GOOS == "linux" {
    time.Sleep(2 * time.Second)
  }

  controller.left.SetKeys(keybd_event.VK_LEFT)
  controller.right.SetKeys(keybd_event.VK_RIGHT)

  return controller
}

func (c *KeyHolder) PressLeft() {
  c.left.Launching()
}

func (c *KeyHolder) PressRight() {
  c.right.Launching()
}
