package test

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder           = "\"[^\"]*NonFungibleToken.cdc\""
	LockedNFTAddressPlaceholder     = "\"[^\"]*LockedNFT.cdc\""
	metadataViewsAddressPlaceholder = "0xMETADATAVIEWSADDRESS"
	exampleNFTAddressPlaceholder    = "0xEXAMPLENFTADDRESS"
)

const (
	LockedNFTPath        = "../../../contracts/LockedNFT.cdc"
	ExampleNFTPath       = "../../../contracts/ExampleNFT.cdc"
	TransactionsRootPath = "../../../transactions"
	ScriptsRootPath      = "../../../scripts"

	// Accounts
	SetupAccountTxPath       = TransactionsRootPath + "/setup_collection.cdc"
	IsAccountSetupScriptPath = ScriptsRootPath + "/is_account_setup.cdc"

	// NFTs
	SetupExampleNFTxPath = TransactionsRootPath + "/user/setup_example_nft.cdc"
	MintExampleNFTTxPath = TransactionsRootPath + "/examplenft/mint_nft.cdc"

	// MetadataViews
	MetadataViewsContractsBaseURL = "https://raw.githubusercontent.com/onflow/flow-nft/master/contracts/"
	MetadataViewsInterfaceFile    = "MetadataViews.cdc"
	MetadataFTReplaceAddress      = `"./utility/FungibleToken.cdc"`
	MetadataNFTReplaceAddress     = `"./NonFungibleToken.cdc"`

	// LockedNFT
	GetLockedTokenByIDScriptPath = ScriptsRootPath + "/get_locked_token.cdc"
)

// ------------------------------------------------------------
// Accounts
// ------------------------------------------------------------
func rX(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	LockedNFTRe := regexp.MustCompile(LockedNFTAddressPlaceholder)
	code = LockedNFTRe.ReplaceAll(code, []byte("0x"+contracts.LockedNFTAddress.String()))

	code = []byte(strings.ReplaceAll(string(code), metadataViewsAddressPlaceholder, "0x"+contracts.MetadataViewsAddress.String()))

	return code
}

func replaceAddresses(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	DapperSportRe := regexp.MustCompile(LockedNFTAddressPlaceholder)
	code = DapperSportRe.ReplaceAll(code, []byte("0x"+contracts.LockedNFTAddress.String()))

	code = []byte(strings.ReplaceAll(string(code), metadataViewsAddressPlaceholder, "0x"+contracts.MetadataViewsAddress.String()))
	code = []byte(strings.ReplaceAll(string(code), exampleNFTAddressPlaceholder, "0x"+contracts.LockedNFTAddress.String()))

	return code
}

func LoadLockedNFTContract(nftAddress flow.Address, metadataViewsAddress flow.Address) []byte {
	code := readFile(LockedNFTPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))
	code = []byte(strings.ReplaceAll(string(code), metadataViewsAddressPlaceholder, "0x"+metadataViewsAddress.String()))

	return code
}

func LoadExampleNFTContract(nftAddress flow.Address, metadataViewsAddress flow.Address) []byte {
	code := readFile(ExampleNFTPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))
	code = []byte(strings.ReplaceAll(string(code), metadataViewsAddressPlaceholder, "0x"+metadataViewsAddress.String()))

	return code
}

func setupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(SetupAccountTxPath),
		contracts,
	)
}

func setupExampleNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(SetupExampleNFTxPath),
		contracts,
	)
}

func isAccountSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(IsAccountSetupScriptPath),
		contracts,
	)
}

func mintExampleNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(MintExampleNFTTxPath),
		contracts,
	)
}

func DownloadFile(url string) ([]byte, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func LoadMetadataViews(ftAddress flow.Address, nftAddress flow.Address) []byte {
	code, _ := DownloadFile(MetadataViewsContractsBaseURL + MetadataViewsInterfaceFile)
	code = []byte(strings.Replace(strings.Replace(string(code), MetadataFTReplaceAddress, "0x"+ftAddress.String(), 1), MetadataNFTReplaceAddress, "0x"+nftAddress.String(), 1))

	return code
}
