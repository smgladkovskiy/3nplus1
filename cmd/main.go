package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var (
	cpuprofile bool
	memprofile bool
	maxIPower  int
	approach   string

	rootCmd = &cobra.Command{
		Use:  "collatz",
		Args: cobra.MinimumNArgs(1),
	}

	iteratorCmd = &cobra.Command{
		Use:   "iterator",
		Short: "Runs Collatz iterations from 1 to passed power of 10",
		Long:  `Collatz iterations from 1 to passed power of 10. Long description will be here later.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if maxIPower >= 9 { // nolint: gomnd
				return ErrMaxPower
			}

			if approach != "recursion" && approach != "cycle" {
				return ErrBadApproach
			}

			return nil
		},
		Run: runIterator,
	}

	ErrMaxPower    = errors.New("maximum number is too much for adequate waiting interval (more than couple of hours)")
	ErrBadApproach = errors.New("unknown approach is used")
)

func Execute() {
	iteratorCmd.Flags().BoolVarP(&cpuprofile, "profile-cpu", "c", false, "iterator cpu profiling")
	iteratorCmd.Flags().BoolVarP(&memprofile, "profile-memory", "m", false, "iterator memory profiling")
	iteratorCmd.Flags().IntVarP(&maxIPower, "max-power", "p", 6, "limit for iteration as the power of 10 with maximum 9 for now")
	iteratorCmd.Flags().StringVarP(&approach, "approach", "a", "recursion", "count approach (recursion or cycle)")

	// _ = iteratorCmd.MarkFlagRequired("max-power")

	rootCmd.AddCommand(iteratorCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
