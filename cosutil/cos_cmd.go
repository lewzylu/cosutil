package cosutil

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	)
func init(){
	var echoTimes int
	
	  var cmdconfig = &cobra.Command{
		Use:   "coscmd config [-h] -a SECRET_ID -s SECRET_KEY -u APPID -b BUCKET -r REGION [-m MAX_THREAD] [-p PART_SIZE]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
	For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	  }
	
	  var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		  fmt.Println("Print: " + strings.Join(args, " "))
		},
	  }
	
	  var cmdTimes = &cobra.Command{
		Use:   "times [# times] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
	a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		  for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		  }
		},
	  }
	
	  cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	
	  var rootCmd = &cobra.Command{Use: "app"}
	  rootCmd.AddCommand(cmdconfig, cmdEcho)
	  cmdEcho.AddCommand(cmdTimes)
	  rootCmd.Execute()
}