package main

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertIntoLightningPort(mac)
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.InsertIntoLightningPort(windowsMachineAdapter)
}
