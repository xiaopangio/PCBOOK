package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/xiaopangio/pcbook/pb"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch r.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}
func randomBool() bool {
	return r.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700u",
		"Ryzen 5 PRO 3500u",
		"Ryzen 3 PRO 3200GE",
	)
}
func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}
func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 1660-Ti",
			"RTX 3060",
			"RTX 3090",
			"RTX 4060",
			"RTX 4090",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
	)
}
func randomStringFromSet(s ...string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	return s[r.Intn(n)]
}
func randomInt(min, max int) int {
	return min + r.Intn(max-min+1)
}
func randomFloat64(min, max float64) float64 {
	return min + r.Float64()*(max-min)
}
func randomFloat32(min, max float32) float32 {
	return min + r.Float32()*(max-min)
}
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}
func randomPanel() pb.Screen_Panel {
	if r.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}
func randomID() string {
	return uuid.New().String()
}
func randomLaptapBrand() string {
	return randomStringFromSet("HuaWei", "Lenovo", "Apple")
}
func randomLaptapName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook", "Macbook Pro")
	case "HuaWei":
		return randomStringFromSet("HuaWei MateBook", "HuaWei MateBook Pro")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad p53")
	}
}
func RandomLaptapScore() float64 {
	return float64(randomInt(1, 10))
}
