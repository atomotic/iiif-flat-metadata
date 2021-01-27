## Case 1: folder with all JSON

Open terminal, write:

`for manifest in * ; do iiif-flat-metadata $manifest; done > ../lista.json`

Open `lista.json` file with Openrefine.
