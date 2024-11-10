package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Start TShark to capture packets in the background
func startTshark() {
	cmd := exec.Command("tshark", "-i", "any", "-w", "capture.pcap", "-c", "1000")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error starting TShark: %v\n", err)
		return
	}
	fmt.Println("TShark started. Capturing 1000 packets to capture.pcap...")
}

// Display the captured .pcap file and save the output to separate JSON files
func displayPcapFile(filename, localIP, wlanIP string) {
	fmt.Println("Analyzing file", filename)

	// Create output files for both incoming and outgoing traffic for local and wlan0
	localSendFile, err := os.Create("local-send.json")
	if err != nil {
		fmt.Printf("Error creating local-send.json: %v\n", err)
		return
	}
	defer localSendFile.Close()

	localReceiveFile, err := os.Create("local-receive.json")
	if err != nil {
		fmt.Printf("Error creating local-receive.json: %v\n", err)
		return
	}
	defer localReceiveFile.Close()

	wlanSendFile, err := os.Create("wlan0-send.json")
	if err != nil {
		fmt.Printf("Error creating wlan0-send.json: %v\n", err)
		return
	}
	defer wlanSendFile.Close()

	wlanReceiveFile, err := os.Create("wlan0-receive.json")
	if err != nil {
		fmt.Printf("Error creating wlan0-receive.json: %v\n", err)
		return
	}
	defer wlanReceiveFile.Close()

	// TShark command to read the .pcap file and output the sent traffic (outgoing) for local IP
	localSendCmd := exec.Command("tshark", "-r", filename, "-T", "json", "-Y", "ip.src == "+localIP)
	localSendCmd.Stdout = localSendFile

	// TShark command to read the .pcap file and output the received traffic (incoming) for local IP
	localReceiveCmd := exec.Command("tshark", "-r", filename, "-T", "json", "-Y", "ip.dst == "+localIP)
	localReceiveCmd.Stdout = localReceiveFile

	// TShark command to read the .pcap file and output the sent traffic (outgoing) for wlan0 IP
	wlanSendCmd := exec.Command("tshark", "-r", filename, "-T", "json", "-Y", "ip.src == "+wlanIP)
	wlanSendCmd.Stdout = wlanSendFile

	// TShark command to read the .pcap file and output the received traffic (incoming) for wlan0 IP
	wlanReceiveCmd := exec.Command("tshark", "-r", filename, "-T", "json", "-Y", "ip.dst == "+wlanIP)
	wlanReceiveCmd.Stdout = wlanReceiveFile

	// Run the commands concurrently
	err = localSendCmd.Start()
	if err != nil {
		fmt.Printf("Error starting local send traffic command: %v\n", err)
		return
	}

	err = localReceiveCmd.Start()
	if err != nil {
		fmt.Printf("Error starting local receive traffic command: %v\n", err)
		return
	}

	err = wlanSendCmd.Start()
	if err != nil {
		fmt.Printf("Error starting wlan0 send traffic command: %v\n", err)
		return
	}

	err = wlanReceiveCmd.Start()
	if err != nil {
		fmt.Printf("Error starting wlan0 receive traffic command: %v\n", err)
		return
	}

	// Wait for all commands to finish
	err = localSendCmd.Wait()
	if err != nil {
		fmt.Printf("Error capturing local send traffic: %v\n", err)
		return
	}

	err = localReceiveCmd.Wait()
	if err != nil {
		fmt.Printf("Error capturing local receive traffic: %v\n", err)
		return
	}

	err = wlanSendCmd.Wait()
	if err != nil {
		fmt.Printf("Error capturing wlan0 send traffic: %v\n", err)
		return
	}

	err = wlanReceiveCmd.Wait()
	if err != nil {
		fmt.Printf("Error capturing wlan0 receive traffic: %v\n", err)
		return
	}

	fmt.Println("Traffic data captured:")
	fmt.Println("- 'local-send.json'")
	fmt.Println("- 'local-receive.json'")
	fmt.Println("- 'wlan0-send.json'")
	fmt.Println("- 'wlan0-receive.json'")
}

func main() {
	// Start TShark to capture packets in the background
	startTshark()

	// Wait for TShark to capture packets
	fmt.Println("Waiting for TShark capture to finish...")
	// Wait for TShark to capture packets for some time (you can modify based on actual packet capture duration)
	time.Sleep(15 * time.Second) // Adjust this time if necessary

	// Set the IPs for localhost and wlan0 (replace with actual wlan0 IP if needed)
	localIP := "127.0.0.1"     // Localhost IP
	wlanIP := "192.168.51.172" // Replace with your actual wlan0 IP (you can find it using `ifconfig` or `ip a`)

	// Display the captured pcap file and save the output to four separate JSON files
	displayPcapFile("capture.pcap", localIP, wlanIP)
}
