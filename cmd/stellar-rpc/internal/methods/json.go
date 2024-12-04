package methods

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/stellar/go/xdr"

	"github.com/stellar/stellar-rpc/cmd/stellar-rpc/internal/db"
	"github.com/stellar/stellar-rpc/cmd/stellar-rpc/internal/xdr2json"
)

const (
	FormatBase64 = "base64"
	FormatJSON   = "json"
)

var errInvalidFormat = fmt.Errorf(
	"expected %s for optional 'xdrFormat'",
	strings.Join([]string{FormatBase64, FormatJSON}, ", "))

func IsValidFormat(format string) error {
	switch format {
	case "":
	case FormatJSON:
	case FormatBase64:
	default:
		return errors.Wrapf(errInvalidFormat, "got '%s'", format)
	}
	return nil
}

func transactionToJSON(tx db.Transaction) (
	[]byte,
	[]byte,
	[]byte,
	error,
) {
	var err error
	var result, resultMeta, envelope []byte

	result, err = xdr2json.ConvertBytes(xdr.TransactionResult{}, tx.Result)
	if err != nil {
		return result, envelope, resultMeta, err
	}

	envelope, err = xdr2json.ConvertBytes(xdr.TransactionEnvelope{}, tx.Envelope)
	if err != nil {
		return result, envelope, resultMeta, err
	}

	resultMeta, err = xdr2json.ConvertBytes(xdr.TransactionMeta{}, tx.Meta)
	if err != nil {
		return result, envelope, resultMeta, err
	}

	return result, envelope, resultMeta, nil
}

func ledgerToJSON(meta *xdr.LedgerCloseMeta) ([]byte, []byte, error) {
	var err error
	var closeMetaJSON, headerJSON []byte

	closeMetaJSON, err = xdr2json.ConvertInterface(*meta)
	if err != nil {
		return nil, nil, err
	}

	headerJSON, err = xdr2json.ConvertInterface(meta.LedgerHeaderHistoryEntry())
	if err != nil {
		return nil, nil, err
	}

	return closeMetaJSON, headerJSON, nil
}