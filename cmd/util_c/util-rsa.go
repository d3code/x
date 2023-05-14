package util_c

import (
    "github.com/d3code/pkg/clog"
    "os"

    "github.com/d3code/pkg/encrypt"
    "github.com/d3code/pkg/files"
    "github.com/d3code/pkg/xerr"
    "github.com/spf13/cobra"
)

func init() {
    Util.AddCommand(rsaCommand)
    rsaCommand.Flags().StringP("directory", "d", ".", "directory to create keys")
    rsaCommand.Flags().BoolP("overwrite", "o", false, "overwrite existing files in directory")
    rsaCommand.Flags().IntP("bits", "b", 4096, "number of bits for key generation")
    rsaCommand.Flags().StringP("private", "p", "generated_key.pem", "private key filename")
    rsaCommand.Flags().StringP("public", "u", "generated_key.pub", "public key filename")
}

var rsaCommand = &cobra.Command{
    Use:   "rsa",
    Short: "RSA key generation",
    Long:  `Generate RSA keys`,
    Run:   rsa,
}

func rsa(cmd *cobra.Command, args []string) {
    bits, _ := cmd.Flags().GetInt("bits")
    if bits != 4096 {
        clog.Info("Specifying{{ --bits | red }} is not implemented yet, using {{ 4096 bits | green }} for key generation")
    }

    privateKeyName := cmd.Flag("private").Value.String()
    publicKeyName := cmd.Flag("public").Value.String()

    if directory := cmd.Flag("directory").Value.String(); directory != "." {
        err := os.Chdir(directory)
        xerr.ExitIfError(err)
    }

    if overwrite, _ := cmd.Flags().GetBool("overwrite"); !overwrite {
        if files.Exist(privateKeyName) {
            clog.Info("File {{ " + privateKeyName + " | red }} exists in directory, use {{ --overwrite | green }} or {{ -o | green }} to overwrite")
        }
        if files.Exist(publicKeyName) {
            clog.Info("File {{ " + publicKeyName + " | red }} exists in directory, use {{ --overwrite | green }} or {{ -o | green }} to overwrite")
        }
        if files.Exist(privateKeyName) || files.Exist(publicKeyName) {
            os.Exit(0)
        }
    }

    privateKey, publicKey := encrypt.RsaGenerate()
    privateToString := encrypt.RsaPrivateToString(privateKey)
    publicToString := encrypt.RsaPublicToString(publicKey)

    cd, err := os.Getwd()
    xerr.ExitIfError(err)

    cwd := string(cd) + "/"

    writeFile(privateToString, privateKeyName)
    clog.Info("Created private key {{" + cwd + privateKeyName + "|blue}}")

    writeFile(publicToString, publicKeyName)
    clog.Info("Created public key {{" + cwd + publicKeyName + "|blue}}")
}

func writeFile(private string, filename string) {

    privateKey := []byte(private)
    err := os.WriteFile(filename, privateKey, 0644)
    xerr.ExitIfError(err)
}
