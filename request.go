package kpsclient

import (
	"fmt"
	"strings"
)

const bodyNS = "http://kps.nvi.gov.tr/2025/08/01"

// defaultZero returns "0" if the value is empty or consists only of whitespace.
// Used because the KPS services expect "0" for empty month/day fields.
func defaultZero(s string) string {
	if strings.TrimSpace(s) == "" {
		return "0"
	}
	return s
}

// BuildTumKutukBody constructs the <Sorgula> XML body for the KPS TumKutukDogrulamaServisi.
//
// The generated XML format is as follows:
//
// <Sorgula xmlns="...">
//
//	<kriterListesi>
//	  <TumKutukDogrulamaSorguKriteri>
//	    <Ad>...</Ad>
//	    <DogumAy>...</DogumAy>
//	    <DogumGun>...</DogumGun>
//	    <DogumYil>...</DogumYil>
//	    <KimlikNo>...</KimlikNo>
//	    <Soyad>...</Soyad>
//	    <TCKKSeriNo>...</TCKKSeriNo>
//	  </TumKutukDogrulamaSorguKriteri>
//	</kriterListesi>
//
// </Sorgula>
//
// Notes:
//   - If DogumAy or DogumGun are left empty, "0" will automatically be used, as required by the service.
//   - TCKKSeriNo is optional.
//   - XML contents are escaped.
//
// Returns: The full XML body as a string.
func BuildTumKutukBody(r QueryRequest) string {
	var sb strings.Builder

	// Root opening
	sb.WriteString(fmt.Sprintf(
		`<Sorgula xmlns="%s" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">`,
		bodyNS,
	))

	sb.WriteString(`<kriterListesi><TumKutukDogrulamaSorguKriteri>`)

	// Fields
	sb.WriteString(fmt.Sprintf(`<Ad>%s</Ad>`, xmlEscape(r.FirstName)))
	sb.WriteString(fmt.Sprintf(`<DogumAy>%s</DogumAy>`, xmlEscape(defaultZero(r.BirthMonth))))
	sb.WriteString(fmt.Sprintf(`<DogumGun>%s</DogumGun>`, xmlEscape(defaultZero(r.BirthDay))))
	sb.WriteString(fmt.Sprintf(`<DogumYil>%s</DogumYil>`, xmlEscape(r.BirthYear)))
	sb.WriteString(fmt.Sprintf(`<KimlikNo>%s</KimlikNo>`, xmlEscape(r.TCNo)))
	sb.WriteString(fmt.Sprintf(`<Soyad>%s</Soyad>`, xmlEscape(r.LastName)))

	// Optional field
	sb.WriteString(fmt.Sprintf(`<TCKKSeriNo>%s</TCKKSeriNo>`, xmlEscape(r.SerialNumber)))

	// Closing
	sb.WriteString(`</TumKutukDogrulamaSorguKriteri></kriterListesi></Sorgula>`)

	return sb.String()
}
