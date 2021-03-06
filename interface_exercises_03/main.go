package main

import "fmt"

// This is generally what USB ports do and how we describe them.
type usb interface {
	input()
	output()
	description()
}

// Keyboard USB Cable
type keyboardUsbCable struct {
	manufacturer string
	model        string
}

func (k keyboardUsbCable) description() {
	fmt.Printf("%s %s keyboard USB cable\n", k.manufacturer, k.model)
}

func (k keyboardUsbCable) input() {
	fmt.Printf("I need to use electricity from the USB port.\n")
}

func (k keyboardUsbCable) output() {
	fmt.Printf("I provide a method for users to interact with a computer.\n")
}

// Camera USB Cable
type cameraUsbCable struct {
	manufacturer string
	model        string
}

func (c cameraUsbCable) description() {
	fmt.Printf("%s %s camera USB cable\n", c.manufacturer, c.model)
}

func (c cameraUsbCable) input() {
	fmt.Printf("I can charge my battery through the USB port.\n")
}

func (c cameraUsbCable) output() {
	fmt.Printf("I can transfer pictures through the USB port.\n")
}

// This function takes the interface and makes use of the methods attached to it.
func usesOfUsb(u usb) {
	u.description()
	u.input()
	u.output()
	fmt.Println()
}

func main() {
	fmt.Println()

	// Create our USB connectors
	myCameraConnector := cameraUsbCable{manufacturer: "Nikon", model: "D90"}
	myKeyboardConnector := keyboardUsbCable{manufacturer: "Apple", model: "A1242"}

	// What can our USB connectors do?
	usesOfUsb(myCameraConnector)
	usesOfUsb(myKeyboardConnector)

}
