package pgk

import (
	"CliTaskManager/cmd/Twitter/handler"
	"github.com/spf13/cobra"
	"strings"
)

var getTweets = &cobra.Command{
	Use:   "tweet",
	Short: "ilgili hesabin tweetlerini verir",
	Run: func(cmd *cobra.Command, args []string) {
		handler.GetTweets(strings.Join(args, ""))
	},
}

var getFollowing = &cobra.Command{
	Use:   "getFollowingAccount",
	Short: "takip edilenleri verir",
	Run: func(cmd *cobra.Command, args []string) {
		handler.GetFollowingAccounts(strings.Join(args, ""))
	},
}

var getFollowers = &cobra.Command{
	Use:   "getFollowerAccount",
	Short: "takipcileri verir",
	Run: func(cmd *cobra.Command, args []string) {
		handler.GetFollowers(strings.Join(args, ""))
	},
}

func init() {
	RootCmd.AddCommand(getTweets)
	RootCmd.AddCommand(getFollowing)
	RootCmd.AddCommand(getFollowers)
}
