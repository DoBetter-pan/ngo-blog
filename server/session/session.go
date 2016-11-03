/** 
* @file session.go
* @brief session 
* @author yingx
* @date 2016-10-17
*/

package session

import (
    "fmt"
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "encoding/base64"
    "strings"
    "strconv"
    "errors"
	"net/http"
)

const (
    session_iv = "Golang Good! 3x!"
    session_key = "f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded"
    session_name = "ng-blog"
)

func PKCS5Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src) % blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length - 1])
    return src[:(length - unpadding)]
}

func AesEncrypt(plaintext []byte) ([]byte, error) {
    key, _ := hex.DecodeString(session_key)
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, errors.New("invalid decrypt key")
    }
    blockSize := block.BlockSize()
    plaintext = PKCS5Padding(plaintext, blockSize)
    iv := []byte(session_iv)
    blockMode := cipher.NewCBCEncrypter(block, iv)

    ciphertext := make([]byte, len(plaintext))
    blockMode.CryptBlocks(ciphertext, plaintext)

    return ciphertext, nil
}

func AesDecrypt(ciphertext []byte) ([]byte, error) {
    key, _ := hex.DecodeString(session_key)
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, errors.New("invalid decrypt key")
    }

    blockSize := block.BlockSize()

    if len(ciphertext) < blockSize {
        return nil, errors.New("ciphertext too short")
    }

    iv := []byte(session_iv)
    if len(ciphertext) % blockSize != 0 {
        return nil, errors.New("ciphertext is not a multiple of the block size")
    }

    blockModel := cipher.NewCBCDecrypter(block, iv)

    plaintext := make([]byte, len(ciphertext))
    blockModel.CryptBlocks(plaintext, ciphertext)
    plaintext = PKCS5UnPadding(plaintext)

    return plaintext, nil
}

func MakeSession(id int64, name, role, nonce string) string {
    encoded := ""
    s := fmt.Sprintf("%d\t%s\t%s\t%s", id, name, role, nonce)

    b, err := AesEncrypt([]byte(s))
    if err == nil {
        encoded = base64.StdEncoding.EncodeToString(b)
        encoded += "a"
    } else {
        encoded = base64.StdEncoding.EncodeToString([]byte(s))
        encoded += "b"
    }

    return encoded
}

func ValidateSession(str string) (bool, int64, string, string, string) {
    l := len(str)
    t := str[l - 1]
    b := str[:l - 1]

    if t != 'a' && t != 'b' {
        return false, 0, "", "", ""
    }

    decoded, err := base64.StdEncoding.DecodeString(b)
    if err != nil {
        return false, 0, "", "", ""
    }
    if t == 'a' {
        decoded, err = AesDecrypt(decoded)
        if err != nil {
            return false, 0, "", "", ""
        }
    }
    strArray := strings.Split(string(decoded), "\t")
    if len(strArray) != 4 {
        return false, 0, "", "", ""
    }
    i64, err := strconv.ParseInt(strArray[0], 10, 64)
    if err != nil {
        return false, 0, "", "", ""
    }

    return true, i64, strArray[1], strArray[2], strArray[3]
}

func ValidateSessionByCookie(r *http.Request) (bool, int64, string, string, string) {
    userCookie, err := r.Cookie(session_name)
    if err != nil {
        return false, 0, "", "", ""
    }

    validated, id, name, role, nonce := ValidateSession(userCookie.Value)
    //need check database
    return validated, id, name, role, nonce
}

func WriteBackSessionCookie(w http.ResponseWriter, id int64, name, role, nonce string, path string, timeout int) {
    var cookie http.Cookie
    cookieValue := MakeSession(id, name, role, nonce)
    if timeout == 0 {
        cookie = http.Cookie{Name: session_name, Value: cookieValue, Path: path }
    } else {
        cookie = http.Cookie{Name: session_name, Value: cookieValue, Path: path, MaxAge: timeout}
    }
    http.SetCookie(w, &cookie)
}
