package txscript

import (
	"bytes"
	"testing"
)

func TestRedeemToPkScript(t *testing.T) {
	redeem := []byte{
		82, 65, 4, 2, 83, 21, 116, 23, 208, 223, 22, 63, 33, 52, 55, 175, 75,
		119, 114, 250, 19, 22, 177, 255, 206, 20, 137, 199, 197, 174, 244, 194,
		15, 245, 81, 94, 80, 76, 230, 243, 156, 11, 161, 17, 245, 68, 250, 134,
		98, 63, 123, 206, 106, 17, 129, 179, 210, 5, 155, 242, 97, 194, 119,
		175, 122, 32, 45, 65, 4, 219, 47, 252, 31, 82, 125, 34, 225, 107, 200,
		88, 45, 78, 46, 221, 232, 119, 33, 245, 22, 107, 5, 210, 37, 38, 160,
		107, 38, 218, 198, 70, 140, 97, 52, 204, 27, 97, 252, 237, 156, 154,
		175, 86, 193, 177, 245, 210, 222, 244, 235, 8, 179, 15, 187, 126, 249,
		192, 138, 143, 251, 198, 230, 98, 172, 82, 174,
	}

	want := []byte{
		118, 169, 20, 10, 63, 117, 193, 26, 249, 104, 211, 169, 228, 39, 135,
		197, 179, 65, 183, 169, 3, 163, 165, 136, 192,
	}

	got := RedeemToPkScript(redeem)
	if !bytes.Equal(got, want) {
		t.Errorf("got pkscript = %x want %x", got, want)
	}
}
