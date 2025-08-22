package main

func main() {

	upRate, downRate := RouterRequest()
	WriteToInflux(upRate, downRate)

}
