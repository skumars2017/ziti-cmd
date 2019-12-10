/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package cmd

import (
	"github.com/netfoundry/ziti-cmd/ziti/cmd/ziti/internal/log"
	"io"

	cmdutil "github.com/netfoundry/ziti-cmd/ziti/cmd/ziti/cmd/factory"
	cmdhelper "github.com/netfoundry/ziti-cmd/ziti/cmd/ziti/cmd/helpers"

	"github.com/spf13/cobra"
)

const ()

type ArtOptions struct {
	CommonOptions
}

func NewCmdArt(f cmdutil.Factory, out io.Writer, errOut io.Writer) *cobra.Command {
	options := &ArtOptions{
		CommonOptions: CommonOptions{
			Factory: f,
			Out:     out,
			Err:     errOut,
		},
	}

	cmd := &cobra.Command{
		Use:   "art",
		Short: "Print the Ziti logo as ascii art :)",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			cmdhelper.CheckErr(err)
		},
	}
	options.addCommonFlags(cmd)

	return cmd
}

// Run ...
func (o *ArtOptions) Run() error {

	log.Info(
		`
			::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			:::::::::::::::::::,::$77777777777777,:,::::::::::::::::::::
			::::::::::::::::::77777777777777777777777~,:::::::::::::::::
			:::::::::::::::77777777777777II7777777777777,:::::::::::::::
			::::::::::::$777777777777777I.:7777777777777777,::::::::::::
			::::::::::77777777777777777I...I7777777777777777I:::::::::::
			:::::::::77777777777777777I....?777777777777777777::::::::::
			:::::::$77777777777777777I......77777777777777777777::::::::
			::::::777777777777777777I.......I77777777777777777777,::::::
			:::::777777777777777777I....?...?777777777777777777777::::::
			:::,777777777777777777I....I7?...777777777777777777777$:::::
			:::777777777777777777I....I77I...I777777777777777777777$::::
			:::77777777777777777I....I7777...?7777777777777777777777::::
			::77777777777777777I....I77777?..,77777777777777777777777:::
			::7777777777777777I....I777777I...I77777777777777$7$$$$7$,::
			:$777777777777777I....I77777777...?7777777777777$$77777777::
			:777777777777777I ...I777II7777?...I.I7777777$777777777777::
			:77777777777777I....I777I..7777I.......?I777$$$$$77$$$$7$$::
			:7777777777777I....?I77I...I7777..........I777777$$$$$7$$$,:
			:77777777777777?..  .??.   ?7777?  ..??.   .?7$7$$$7$$$$$7::
			,7777777777777777I..........I$77I...I777?....77777$7$$$$$$,,
			:7777777777777777777?.......I7$$7..I777I....7$$$$$$$$$$$$$::
			:777777777777777777777I.I=..?77777777$7....77$$$$$$$$7$$$$::
			:777777777777777777777777I...I$7777777....77$$$$$$$$$$$$$$::
			::77777777777777$7$7$$$$$I...?7$$7$77....7$$$$$$$$$$$$$$$:::
			::777777777777777777$$$777+..~77$$7I....77$$$$$$$$$$$$$$$:::
			:::77777777777777777777$$7I...7$$$I....7$7$$$$$$$$$$$$$$::::
			:::Z77777777$7777777777$77I...?$77....I$$$$$$$$$$$$$$$$$::::
			::::77777$$$$$7777$$$$$$$$7:..+77....I$$$$$$$$$$$$$$$$$:::::
			:::::77777$777$$$$777$$$$77I...I....I$$$$$$$$$$$$$$$$$::::::
			::::::$7777777$7777$$$7$$$$I...... I$$$$$$$$$$$$$$$$7:::::::
			:::::::?$$$$$$$$$$$$$$$$$$$7=.....I$$$$$$$$$$$$$$$$=::::::::
			:::::::::7$$$$$7$$$$$$$$$$$$?....77$$$$$$$$$$$$$$$::::::::::
			::::::::::,7$$7$$$$$$$$$$$$$7...I$$$$$$$$$$$$$$$::::::::::::
			::::::::::::~$$$$$$$$$$$$$$$7?.I$$$$$$$$$$$$$$::::::::::::::
			:::::::::::::::$$$$$$$$$$$$$$77$$$$$$$$$$$$$::::::::::::::::
			::::::::::::::::::7$$$$$$$$$$$$$$$$$$$$$$:::::::::::::::::::
			:::::::::::::::::::::::$$$$$$$$$$$$$::::::::::::::::::::::::
			::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

`)

	return nil
}