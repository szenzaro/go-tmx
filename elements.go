package tmx

import "encoding/xml"

// Tmx document - The <tmx> element encloses all the other elements of the document.
type Tmx struct {
	XMLName xml.Name `xml:"tmx,omitempty"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr,omitempty"`
	Header  Header   `xml:"header,omitempty"`
	Body    Body     `xml:"body,omitempty"`
}

// Header represents the <header> element. It contains information pertaining to the whole document.
type Header struct {
	Text                string `xml:",chardata"`
	Creationtool        string `xml:"creationtool,attr,omitempty"`
	Creationtoolversion string `xml:"creationtoolversion,attr,omitempty"`
	Datatype            string `xml:"datatype,attr,omitempty"`
	Segtype             string `xml:"segtype,attr,omitempty"`
	Adminlang           string `xml:"adminlang,attr,omitempty"`
	Srclang             string `xml:"srclang,attr,omitempty"`
	OTmf                string `xml:"o-tmf,attr,omitempty"`
	Creationdate        string `xml:"creationdate,attr,omitempty"`
	Creationid          string `xml:"creationid,attr,omitempty"`
	Changedate          string `xml:"changedate,attr,omitempty"`
	Changeid            string `xml:"changeid,attr,omitempty"`
	OEncoding           string `xml:"o-encoding,attr,omitempty"`
	Note                []Note `xml:"note,omitempty"`
	Prop                []Prop `xml:"prop,omitempty"`
	Ude                 []Ude  `xml:"ude,omitempty"`
}

// Ude represents User-Defined Encoding - The <ude> element is used to specify a set of user-defined characters and/or, optionally their mapping from Unicode to the user-defined encoding.
type Ude struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr,omitempty"`
	Base string `xml:"base,attr,omitempty"`
	Map  []Map  `xml:"map,omitempty"`
}

// Map represents the <map/> element. It is used to specify a user-defined character and some of its properties.
type Map struct {
	Text    string `xml:",chardata"`
	Unicode string `xml:"unicode,attr,omitempty"`
	Code    string `xml:"code,attr,omitempty"`
	Ent     string `xml:"ent,attr,omitempty"`
	Subst   string `xml:"subst,attr,omitempty"`
}

// Body represents the <body> element. It encloses the main data, the set of <tu> elements that are comprised within the file.
type Body struct {
	Text string `xml:",chardata"`
	Tu   []Tu   `xml:"tu"`
}

// Tuv represents a Translation Unit Variant - The <tuv> element specifies text in a given language.
type Tuv struct {
	Text                string `xml:",chardata"`
	Lang                string `xml:"lang,attr,omitempty"`
	Creationdate        string `xml:"creationdate,attr,omitempty"`
	Creationid          string `xml:"creationid,attr,omitempty"`
	Changedate          string `xml:"changedate,attr,omitempty"`
	Changeid            string `xml:"changeid,attr,omitempty"`
	Seg                 Seg    `xml:"seg"`
	Note                []Note `xml:"note,omitempty"`
	Prop                []Prop `xml:"prop,omitempty"`
	OEncoding           string `xml:"o-encoding,attr,omitempty"`
	Datatype            string `xml:"datatype,attr,omitempty"`
	Usagecount          string `xml:"usagecount,attr,omitempty"`
	Lastusagedate       string `xml:"lastusagedate,attr,omitempty"`
	Creationtool        string `xml:"creationtool,attr,omitempty"`
	Creationtoolversion string `xml:"creationtoolversion,attr,omitempty"`
	OTmf                string `xml:"o-tmf,attr,omitempty"`
}

// Tu represents a translation unit.  The <tu> element contains the data for a given translation unit.
type Tu struct {
	Text                string `xml:",chardata"`
	Tuid                string `xml:"tuid,attr,omitempty"`
	Datatype            string `xml:"datatype,attr,omitempty"`
	Usagecount          string `xml:"usagecount,attr,omitempty"`
	Lastusagedate       string `xml:"lastusagedate,attr,omitempty"`
	Srclang             string `xml:"srclang,attr,omitempty"`
	Note                []Note `xml:"note,omitempty"`
	Prop                []Prop `xml:"prop,omitempty"`
	Tuv                 []Tuv  `xml:"tuv"`
	OEncoding           string `xml:"o-encoding,attr,omitempty"`
	Creationtool        string `xml:"creationtool,attr,omitempty"`
	Creationtoolversion string `xml:"creationtoolversion,attr,omitempty"`
	Creationdate        string `xml:"creationdate,attr,omitempty"`
	Creationid          string `xml:"creationid,attr,omitempty"`
	Changedate          string `xml:"changedate,attr,omitempty"`
	Segtype             string `xml:"segtype,attr,omitempty"`
	Changeid            string `xml:"changeid,attr,omitempty"`
	OTmf                string `xml:"o-tmf,attr,omitempty"`
}

type Prop struct {
	Text      string `xml:",chardata"`
	Type      string `xml:"type,attr,omitempty"`
	OEncoding string `xml:"o-encoding,attr,omitempty"`
	Lang      string `xml:"lang,attr,omitempty"`
}

// Note - The <note> element. It is used for comments.
type Note struct {
	Text      string `xml:",chardata"`
	Lang      string `xml:"lang,attr,omitempty"`
	OEncoding string `xml:"o-encoding,attr,omitempty"`
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
