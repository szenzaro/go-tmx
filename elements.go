package tmx

import "encoding/xml"

// Tmx document - The <tmx> element encloses all the other elements of the document.
type Tmx struct {
	XMLName xml.Name `xml:"tmx"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Header  Header   `xml:"header"`
	Body    Body     `xml:"body"`
}

// Header represents the <header> element. It contains information pertaining to the whole document.
type Header struct {
	Text                string `xml:",chardata"`
	Creationtool        string `xml:"creationtool,attr"`
	Creationtoolversion string `xml:"creationtoolversion,attr"`
	Datatype            string `xml:"datatype,attr"`
	Segtype             string `xml:"segtype,attr"`
	Adminlang           string `xml:"adminlang,attr"`
	Srclang             string `xml:"srclang,attr"`
	OTmf                string `xml:"o-tmf,attr"`
	Creationdate        string `xml:"creationdate,attr"`
	Creationid          string `xml:"creationid,attr"`
	Changedate          string `xml:"changedate,attr"`
	Changeid            string `xml:"changeid,attr"`
	OEncoding           string `xml:"o-encoding,attr"`
	Note                []Note `xml:"note"`
	Prop                []Prop `xml:"prop"`
	Ude                 []Ude  `xml:"ude"`
}

// Ude represents User-Defined Encoding - The <ude> element is used to specify a set of user-defined characters and/or, optionally their mapping from Unicode to the user-defined encoding.
type Ude struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
	Base string `xml:"base,attr"`
	Map  []Map  `xml:"map"`
}

// Map represents the <map/> element. It is used to specify a user-defined character and some of its properties.
type Map struct {
	Text    string `xml:",chardata"`
	Unicode string `xml:"unicode,attr"`
	Code    string `xml:"code,attr"`
	Ent     string `xml:"ent,attr"`
	Subst   string `xml:"subst,attr"`
}

// Body represents the <body> element. It encloses the main data, the set of <tu> elements that are comprised within the file.
type Body struct {
	Text string `xml:",chardata"`
	Tu   []Tu   `xml:"tu"`
}

// Tuv represents a Translation Unit Variant - The <tuv> element specifies text in a given language.
type Tuv struct {
	Text                string `xml:",chardata"`
	Lang                string `xml:"lang,attr"`
	Creationdate        string `xml:"creationdate,attr"`
	Creationid          string `xml:"creationid,attr"`
	Changedate          string `xml:"changedate,attr"`
	Changeid            string `xml:"changeid,attr"`
	Seg                 Seg    `xml:"seg"`
	Note                []Note `xml:"note"`
	Prop                []Prop `xml:"prop"`
	OEncoding           string `xml:"o-encoding,attr"`
	Datatype            string `xml:"datatype,attr"`
	Usagecount          string `xml:"usagecount,attr"`
	Lastusagedate       string `xml:"lastusagedate,attr"`
	Creationtool        string `xml:"creationtool,attr"`
	Creationtoolversion string `xml:"creationtoolversion,attr"`
	OTmf                string `xml:"o-tmf,attr"`
}

// Tu represents a translation unit.  The <tu> element contains the data for a given translation unit.
type Tu struct {
	Text                string `xml:",chardata"`
	Tuid                string `xml:"tuid,attr"`
	Datatype            string `xml:"datatype,attr"`
	Usagecount          string `xml:"usagecount,attr"`
	Lastusagedate       string `xml:"lastusagedate,attr"`
	Srclang             string `xml:"srclang,attr"`
	Note                []Note `xml:"note"`
	Prop                []Prop `xml:"prop"`
	Tuv                 []Tuv  `xml:"tuv"`
	OEncoding           string `xml:"o-encoding,attr"`
	Creationtool        string `xml:"creationtool,attr"`
	Creationtoolversion string `xml:"creationtoolversion,attr"`
	Creationdate        string `xml:"creationdate,attr"`
	Creationid          string `xml:"creationid,attr"`
	Changedate          string `xml:"changedate,attr"`
	Segtype             string `xml:"segtype,attr"`
	Changeid            string `xml:"changeid,attr"`
	OTmf                string `xml:"o-tmf,attr"`
}

type Prop struct {
	Text      string `xml:",chardata"`
	Type      string `xml:"type,attr"`
	OEncoding string `xml:"o-encoding,attr"`
	Lang      string `xml:"lang,attr"`
}

// Note - The <note> element. It is used for comments.
type Note struct {
	Text      string `xml:",chardata"`
	Lang      string `xml:"lang,attr"`
	OEncoding string `xml:"o-encoding,attr"`
}

// Seg represents the <seg> element. It contains the text of the given segment.
// There is no length limitation to the content of a <seg> element.
// All spacing and line-breaking characters are significant within a <seg> element.
type Seg struct {
	Text string `xml:",chardata"`
	// 	Bpt []Bpt
	// 	Ept    []Ept
	// 	It     []It
	// 	Ph     []Ph
	// 	Hi     []Hi
}
