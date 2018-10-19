## Folder con già i JSON scaricati

Sul terminale, scrivere:

`for manifest in * ; do iiif-flat-metadata $manifest; done > ../lista.json`

Aprire il file `lista.json` con Openrefine.
