package dl

import (
	"github.com/iyear/tdl/app/dl"
	"github.com/iyear/tdl/pkg/consts"
	"github.com/spf13/cobra"
)

var (
	urls, files []string
	template    string
)

var Cmd = &cobra.Command{
	Use:     "dl",
	Aliases: []string{"download"},
	Short:   "Download anything from Telegram (protected) chat",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dl.Run(cmd.Context(), template, urls, files)
	},
}

func init() {
	Cmd.Flags().StringSliceVarP(&urls, consts.FlagDlUrl, "u", []string{}, "telegram message links")
	Cmd.Flags().StringSliceVarP(&files, consts.FlagDlFile, "f", []string{}, "official client exported files")
	Cmd.Flags().StringVar(&template, consts.FlagDlTemplate, "{{ .DialogID }}_{{ .MessageID }}_{{ .FileName }}", "download file name template")
}
