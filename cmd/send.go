package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/otgkey/keyboard"
)

func sendRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	path, err := flags.GetString("device")
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("no input provided")
	}

	d := keyboard.NewDevice(path)

	for _, x := range args {
		err := d.SendString(x)
		if err != nil {
			return err
		}
	}
	return nil
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send keypresses to gadget device",
	RunE:  sendRunner,
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().String("device", "/dev/hidg0", "USB Gadget device")
}
