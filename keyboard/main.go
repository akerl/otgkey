package keyboard

import (
	"fmt"
	"os"
	"strings"
)

// Device defines a connection to a USB Gadget device
type Device struct {
	Path string
}

// NewDevice creates a device from a path
func NewDevice(path string) Device {
	return Device{Path: path}
}

// SendString generates keypresses from a human string
func (d Device) SendString(input string) error {
	mods, keys, err := splitString(input)
	if err != nil {
		return err
	}

	return d.SendKeys(mods, keys)
}

// SendKeys generates keypreesses from text slices of modifier and key names
func (d Device) SendKeys(modstrings, keystrings []string) error {
	modcode, err := parseMods(modstrings)
	if err != nil {
		return err
	}
	keys, err := parseKeys(keystrings)
	if err != nil {
		return err
	}
	return d.SendCodes(modcode, keys)
}

// SendCodes generates keypresses from a stack modifier code and array of keycodes
func (d Device) SendCodes(modcode byte, keys [6]byte) error {
	sequence := [8]byte{}
	sequence[0] = modcode
	for index, key := range keys {
		sequence[index+2] = key
	}
	return d.SendRaw(sequence)
}

// SendRaw generates keypresses from a raw command string
func (d Device) SendRaw(input [8]byte) error {
	fh, err := os.OpenFile(d.Path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer fh.Write(Empty[:])
	_, err = fh.Write(input[:])
	return err
}

func splitString(input string) ([]string, []string, error) {
	parts := strings.Split(input, ":")

	mods := []string{}
	keys := []string{}
	if len(parts) == 2 {
		mods = strings.Split(parts[0], ",")
		keys = strings.Split(parts[1], ",")
	} else if len(parts) == 1 {
		keys = strings.Split(parts[0], ",")
	} else {
		return []string{}, []string{}, fmt.Errorf("invalid keypress format")
	}
	return mods, keys, nil
}

func parseMods(mods []string) (byte, error) {
	var result byte
	for _, mod := range mods {
		code, ok := Modifiers[mod]
		if !ok {
			return 0, fmt.Errorf("mod name not found: %s", mod)
		}
		result |= code
	}
	return result, nil
}

func parseKeys(keys []string) ([6]byte, error) {
	if len(keys) > 6 {
		return [6]byte{}, fmt.Errorf("more than 6 keys provided")
	}

	result := [6]byte{}
	for index, key := range keys {
		code, ok := Keys[key]
		if !ok {
			return [6]byte{}, fmt.Errorf("key name not found: %s", key)
		}
		result[index] = code
	}
	return result, nil
}
