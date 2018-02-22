package keepkey

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/karalabe/hid"
	"github.com/solipsis/go-keepkey/pkg/kkProto"
)

// Represents an open HID connection to a keepkey and possibly a
// connection to the debug link if enabled
type Keepkey struct {
	info          hid.DeviceInfo
	device, debug *hid.Device
	vendorID      uint16
	productID     uint16
}

func newKeepkey() *Keepkey {
	return &Keepkey{
		vendorID:  0x2B24,
		productID: 0x0001,
	}
}

// GetDevices establishes connections to all available KeepKey devices and
// their debug interfaces if that is enabled in the firmware
func GetDevices() ([]*Keepkey, error) {

	kk := newKeepkey()

	// tie a keepkey to its debug interface
	type infoPair struct {
		device, debug hid.DeviceInfo
	}

	// Iterate over all connected keepkeys pairing each one with its
	// corresponding debug link if enabled
	deviceMap := make(map[string]*infoPair)
	for _, info := range hid.Enumerate(kk.vendorID, 0) {
		if info.ProductID == kk.productID {

			// get all but the last character of the hid path
			// if the path ends with 0 it is a device if it ends with 1
			// the device is advertising a debug connection
			pathKey := info.Path[:len(info.Path)-1]
			if deviceMap[pathKey] == nil {
				deviceMap[pathKey] = new(infoPair)
			}

			// seperate connection to debug interface if debug link is enabled
			if strings.HasSuffix(info.Path, "1") {
				deviceMap[pathKey].debug = info
			} else {
				fmt.Println("Device: ", info)
				deviceMap[pathKey].device = info
			}
		}
	}

	// Open HID connections to all devices found in the previous step
	var deviceInfo, debugInfo hid.DeviceInfo
	devices := make([]*Keepkey, 0)
	for _, pair := range deviceMap {
		kk = newKeepkey()
		deviceInfo = pair.device
		debugInfo = pair.debug

		if deviceInfo.Path == "" {
			continue
		}

		// Open connection to device
		device, err := deviceInfo.Open()
		if err != nil {
			fmt.Println("Unable to connect to device: dropping, ", err)
			continue
		}

		// debug
		if debugInfo.Path != "" {
			debug, err := debugInfo.Open()
			if err != nil {
				fmt.Println("unable to initiate debug link")
				continue
			}
			fmt.Println("Debug link established")
			kk.debug = debug
		}

		// Ping the device and ask for its features
		if _, err = kk.Initialize(device); err != nil {
			fmt.Println("Unable to contact device, dropping: ", err)
			continue
		}
		devices = append(devices, kk)
	}
	if len(devices) < 1 {
		return devices, errors.New("No keepkeys detected")
	}

	fmt.Println("Connected to ", len(devices), "keepkeys")
	return devices, nil
}

// keepkeyExchange sends a request to the device and streams back the results
// if multiple results are possible the index of the result message is also returned
// based on trezorExchange()
// in https://github.com/go-ethereum/accounts/usbwallet/trezor.go
func (kk *Keepkey) keepkeyExchange(req proto.Message, results ...proto.Message) (int, error) {

	device := kk.device
	debug := false
	// If debug is enabled send over the debug HID interface
	if isDebugMessage(req) && kk.debug != nil {
		device = kk.debug
		debug = true
	}

	// Construct message payload to chunk up
	data, err := proto.Marshal(req)
	if err != nil {
		return 0, err
	}
	payload := make([]byte, 8+len(data))
	copy(payload, []byte{0x23, 0x23}) // ## header
	binary.BigEndian.PutUint16(payload[2:], kkProto.Type(req))
	binary.BigEndian.PutUint32(payload[4:], uint32(len(data)))
	copy(payload[8:], data)

	// stream all the chunks to the device
	chunk := make([]byte, 64)
	chunk[0] = 0x3f // HID Magic number???

	for len(payload) > 0 {
		// create the message to stream and pad with zeroes if necessary
		if len(payload) > 63 {
			copy(chunk[1:], payload[:63])
			payload = payload[63:]
		} else {
			copy(chunk[1:], payload)
			copy(chunk[1+len(payload):], make([]byte, 63-len(payload)))
			payload = nil
		}
		// send over to the device
		if _, err := device.Write(chunk); err != nil {
			return 0, err
		}
	}

	// TODO; support debug requests that return data
	// don't wait for response if sending debug buttonPress
	if debug {
		return 0, nil
	}

	// stream the reply back in 64 byte chunks
	var (
		kind  uint16
		reply []byte
	)
	for {
		// Read next chunk
		if _, err := io.ReadFull(device, chunk); err != nil {
			return 0, err
		}

		//TODO: check transport header

		//if it is the first chunk, retreive the reply message type and total message length
		var payload []byte

		if len(reply) == 0 {
			kind = binary.BigEndian.Uint16(chunk[3:5])
			reply = make([]byte, 0, int(binary.BigEndian.Uint32(chunk[5:9])))
			payload = chunk[9:]
		} else {
			payload = chunk[1:]
		}
		// Append to the reply and stop when filled up
		if left := cap(reply) - len(reply); left > len(payload) {
			reply = append(reply, payload...)
		} else {
			reply = append(reply, payload[:left]...)
			break
		}
	}

	// Try to parse the reply into the requested reply message
	if kind == uint16(kkProto.MessageType_MessageType_Failure) {
		// keepkey returned a failure, extract and return the message
		failure := new(kkProto.Failure)
		if err := proto.Unmarshal(reply, failure); err != nil {
			return 0, err
		}
		return 0, errors.New("keepkey: " + failure.GetMessage())
	}

	// Automatically handle Button/Pin/Passphrase requests
	// handle button requests and forward the results
	if kind == uint16(kkProto.MessageType_MessageType_ButtonRequest) {
		promptButton()
		if kk.debug != nil {
			t := true
			fmt.Println("sending debug press")
			kk.keepkeyExchange(&kkProto.DebugLinkDecision{YesNo: &t}, &kkProto.Success{})
		}
		return kk.keepkeyExchange(&kkProto.ButtonAck{}, results...)
	}
	// handle pin matrix requests and forward the results
	if kind == uint16(kkProto.MessageType_MessageType_PinMatrixRequest) {
		fmt.Println("Pin requested")
		pin, err := promptPin()
		if err != nil {
			return 0, err
		}
		return kk.keepkeyExchange(&kkProto.PinMatrixAck{Pin: &pin}, results...)
	}
	// handle passphrase requests and forward the results
	if kind == uint16(kkProto.MessageType_MessageType_PassphraseRequest) {
		fmt.Println("Passphrase requested")
		pass, err := promptPassphrase()
		if err != nil {
			return 0, err
		}
		return kk.keepkeyExchange(&kkProto.PassphraseAck{Passphrase: &pass}, results...)
	}

	// If the reply we got can be marshaled into one of our expected results
	// marshal it and return the index of the expected result it was
	for i, res := range results {
		if kkProto.Type(res) == kind {
			return i, proto.Unmarshal(reply, res)
		}
	}

	// We did not recieve what we were expecting.
	expected := make([]string, len(results))
	for i, res := range results {
		expected[i] = kkProto.Name(kkProto.Type(res))
	}
	return 0, fmt.Errorf("keepkey: expected reply types %s, got %s", expected, kkProto.Name(kind))
}

// Is the message one we need to send over the debug HID interface
func isDebugMessage(req interface{}) bool {
	switch req.(type) {
	case *kkProto.DebugLinkDecision, *kkProto.DebugLinkFillConfig, *kkProto.DebugLinkGetState:
		return true
	}
	return false
}

// Close closes the hid connection and unassoctiates that hid interface
// with the calling Keepkey
func (kk *Keepkey) Close() {
	if kk.device == nil {
		return
	}
	kk.device.Close()
	kk.device = nil
}
