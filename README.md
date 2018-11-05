# iiif-flat-metadata

Read an IIIF Manifest (from local file or remote URL) and output flattened metadata (multiple values are concatenated with `|`) suitable to be ingested in SOLR or edited with OpenRefine

### Install

    go get github.com/atomotic/iiif-flat-metadata

### Example: URL

```
iiif-flat-metadata http://www.e-codices.unifr.ch/metadata/iiif/fmb-cb-0100/manifest.json  | jq
{
  "@id": "http://www.e-codices.unifr.ch/metadata/iiif/fmb-cb-0100/manifest.json",
  "Attribution": "e-codices - Virtual Manuscript Library of Switzerland",
  "Century": "14th century",
  "Collection Name": "Fondation Martin Bodmer",
  "DOI": "10.5076/e-codices-cb-0100",
  "Date of Origin (English)": "14th century",
  "Description": "",
  "Dimensions": "42.5 x 26.5 cm",
  "Document Type": "Manuscript",
  "Label": "Cologny, Fondation Martin Bodmer, Cod. Bodmer 100",
  "Location": "Cologny",
  "Material": "Parchment",
  "Number of Pages": "292",
  "Online Since": "2013-04-23",
  "Persons": "Author: Accursius, Franciscus Senior; Commentator: Accursius, Franciscus Senior; Author: Angelus Boncambius; Former possessor: Bodmer, Martin; Seller: Kraus, Hans P.",
  "Place of Origin (English)": "Italy",
  "Shelfmark": "Cod. Bodmer 100",
  "Summary (English)": "This 14th century Italian manuscript, probably from Bologna, contains the <i>Digestum Vetus</i>, a fundamental work which attests to the 14th century’s interest in the history of Roman law. It comprises various reference texts, which are systematically accompanied by the <i>Glossa ordinaria</i>, the so-called \"Magna glossa\" by Franciscus Accursius, an interlinear gloss and the gloss of the Gloss, which are works of explanation and instruction for the use of the text. Many manicules or fists (lat <i>manicula, ae</i>: small hands) testify to the assiduous labor which a large number of readers have performed on this dry text. This manuscript contains numerous pecia marks. A detached page (<a href=\"http://www.e-codices.unifr.ch/en/fmb/cb-0100/37ar\">f. 37bis</a>) contains a poem to the reader by the Italian jurist Angelus Boncambius (about 1450).",
  "Text Language": "Latin | French",
  "Title (English)": "Justinian I, Digestum Vetus"
}
```

### Example: File
```
parallel iiif-flat-metadata "{} > {#}.json" :::: list.txt
```
