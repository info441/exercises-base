package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

//Sign returns an HMAC digital signature for the `dataToSign`,
//using the provided `secretKey`. If the io.Reader returns an error
//other than io.EOF, that error will be returned instead.
func Sign(dataToSign io.Reader, secretKey []byte) ([]byte, error) {
	//TODO: implement this function according to the comments above.
	//you can test this function specifically by running:
	// go test -run TestSign
	
	//HINT: io.Copy() can copy all bytes from an io.Reader
	//to an io.Writer(), one chunk at a time so that you
	//don't buffer all the bytes into memory all at once
	
	
	h := hmac.New(sha256.New, secretKey)
	
	if _, err := io.Copy(h, dataToSign); err != nil {
		return nil, err
	}
	
	signature := h.Sum(nil)
	
	return signature, nil
}

//Verify verifies the HMAC `signature` against the `dataThatWasSigned`,
//using the provided `secretKey`. It returns true if the signature was
//valid, false otherwise. If the io.Reader returns an error other than
//io.EOF, that error will be returned instead.
func Verify(dataThatWasSigned io.Reader, secretKey []byte, signature []byte) (bool, error) {
	//TODO: implement this function according to the comments above.
	//you can test this function specifically by running:
	//go test -run TestVerify. You will need to use Sign here to generate the
	//expected signature. Use hmac.Equal to compare the two signatures.
	
	expectedSignature, err := Sign(dataThatWasSigned, secretKey)
	
	if err != nil {
		return false, err
	}
	
	if hmac.Equal(expectedSignature, signature) {
		return true, nil
	}
	
	return false, nil
}