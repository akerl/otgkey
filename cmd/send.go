package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func sendRunner(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	device, err := flags.GetString("device")
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("no input provided")
	}
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
