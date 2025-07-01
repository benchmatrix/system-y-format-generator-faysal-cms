package Service

import (
	"crypto/aes"
	"crypto/cipher"

	"database/sql"
	"encoding/base64"
	"errors"

	"system-y-format-generator-faysal-cms/authMiddleWare/AuthModels"
)

var SubAuth int

func CheckTokenDb(db *sql.DB, token string) bool {

	tsql := " Select JWTTOKEN as Token, USER_ID as TUserId from JWT_TOKEN where JWTTOKEN=?"
	rows, err := db.Query(tsql, token)
	if err != nil {
		return false
	}
	defer rows.Close()

	var instance AuthModels.Token
	for rows.Next() {
		rows.Scan(&instance.Token, &instance.TUserId)
	}

	if (instance.Token) == nil {
		return false
	} else if (instance.TUserId) == nil {
		return false
	} else {
		return true
	}
}

func GetSecret(db *sql.DB) (string, error) {
	var alg = "AES"
	tsql := " Select JWTDECRYPTKEY as JwtKey, ENCRYPTIONSECRETKEY as EncKey,InitializationVector as InitVector  from JWTCONFIGURATION where ENCRYPTIONALGORITHMNAME=?"
	rows, err := db.Query(tsql, alg)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var instance AuthModels.Secrets

	for rows.Next() {
		rows.Scan(&instance.JwtKey, &instance.EncKey, &instance.InitVector)
	}

	return instance.JwtKey, nil

	/*var encodedAlgKey = instance.EncKey
	var jwtSecretKey = instance.JwtKey
	var initilizationVector = instance.InitVector
	data, err := base64.StdEncoding.DecodeString(encodedAlgKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var enckey = string(data)

	h := sha1.New()
	h.Write([]byte(enckey))
	bs := h.Sum(nil)

	arrayShrink := make([]byte, 16)
	copy(arrayShrink[:16], bs)

	vector, err := base64.StdEncoding.DecodeString(initilizationVector)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var vectorInit = string(vector)

	var iv = []byte(vectorInit)

	decrypted, err := decrypt(arrayShrink, iv, jwtSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	return string(decrypted), nil*/
}

func decrypt(key []byte, iv []byte, encText string) ([]byte, error) {

	text, _ := base64.StdEncoding.DecodeString(encText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, text)

	return PKCS5UnPadding(decrypted), nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func GetUserId(db *sql.DB, id string) error {

	tsql := " Select USER_ID as Uid from CONFIG_USERS where EMPID=?"
	rows, err := db.Query(tsql, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var userId int
	for rows.Next() {
		rows.Scan(&userId)
	}
	SubAuth = userId
	return nil
}

func GetUserIdToken() int {

	userId := SubAuth
	return userId
}
