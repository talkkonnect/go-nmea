package main

import (
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/adrianmo/go-nmea"
	"log"
)

func main() {

	options := serial.OpenOptions{
		PortName:              "/dev/ttyACM0",
		BaudRate:              115200,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       4,
		InterCharacterTimeout: 100,
	}

 defer serialPort.Close()	serialPort, err := serial.Open(options)

	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer serialPort.Close()

	reader := bufio.NewReader(serialPort)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		s, err := nmea.Parse(scanner.Text())
		if err == nil {

			if s.DataType() == nmea.TypeRMC {
				m := s.(nmea.RMC)

				//if s.DataType() == nmea.TypeGGA {
				//m := s.(nmea.GGA)

				if m.Latitude != 0 && m.Longitude != 0 {

					// Print out GPS position data...
					//m := s.(nmea.GGA)
					m := s.(nmea.RMC)

					//RMC is the Recommended Minimum Specific GNSS data.
					//RMC Sentence can present Time, Validity, Latitude, Longitude, Speed, Course, Date, Variation
					//GGA is the Time, position, and fix related data of the receiver.
					//GGA Sentence can present Time (od fix), Latitude, Longitude, FixQuality, NumSatellites, HDOP,
					//Altitude, Separation, DGPSAge, DGPSId

					fmt.Printf("Raw sentence: %v\n", m)
					fmt.Printf("Time: %s\n", m.Time)
					fmt.Printf("Validity: %s\n", m.Validity)
					fmt.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
					fmt.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
					fmt.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
					fmt.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
					fmt.Printf("Speed: %f\n", m.Speed)
					fmt.Printf("Course: %f\n", m.Course)
					fmt.Printf("Date: %s\n", m.Date)
					fmt.Printf("Variation: %f\n", m.Variation)

				}

			}

		}

	}

}
