Warning: not ready to use yet!

# mbtmanager
A command line tool for accessing .mbtiles files. 

*In fact, it will be a set of independent executables.*

## Functionality

0. Parameters

`mbtpathIN`: the path of a .mbtiles file as the input.

`txtpathIN`: the path of a text file as the input.

`tiledirIN`: the dirctory of tiles as the input. 

`tileIN`: the path of a tile file as the input.

`mbtpathOUT`: the path of a .mbtiles file as the output.

`txtpathOUT`: the path of a text file as the output.

`tiledirOUT`: the directory to save tiles.

`tileOUT`: the path of a tile file as the output.

`zoom`: zoom level of one or some tiles.

`x`: the column number of a tile or some tiles.

`y`: the row number of a tile or some tiles.

`xyzlist`: a list file of z,x,y numbers of tiles, one line one tile, no header, values seperated by `Tab`.

1. Read a .mbtiles file
+ Show metadata

        mbtinfo-metadata [mbtpathIN]
        mbtinfo-name [mbtpath]
        mbtinfo-format [mbtpathIN]
        mbtinfo-bounds [mbtpathIN]
        mbtinfo-center [mbtpathIN]
        mbtinfo-minzoom [mbtpathIN]
        mbtinfo-maxzoom [mbtpathIN]
        mbtinfo-attribution [mbtpathIN]
        mbtinfo-description [mbtpathIN]
        mbtinfo-type [mbtpathIN]
        mbtinfo-version [mbtpathIN]
        mbtinfo-json [mbtpathIN]
        mbtinfo-which [mbtpathIN]

+ Export metadata

        mbtinfo-metadata [mbtpathIN] [txtpathOUT]
        mbtinfo-name [mbtpath] [txtpathOUT]
        mbtinfo-format [mbtpathIN] [txtpathOUT]
        mbtinfo-bounds [mbtpathIN] [txtpathOUT]
        mbtinfo-center [mbtpathIN] [txtpathOUT]
        mbtinfo-minzoom [mbtpathIN] [txtpathOUT]
        mbtinfo-maxzoom [mbtpathIN] [txtpathOUT]
        mbtinfo-attribution [mbtpathIN] [txtpathOUT]
        mbtinfo-description [mbtpathIN] [txtpathOUT]
        mbtinfo-type [mbtpathIN] [txtpathOUT]
        mbtinfo-version [mbtpathIN] [txtpathOUT]
        mbtinfo-json [mbtpathIN] [txtpathOUT]
        mbtinfo-which [mbtpathIN] [txtpathOUT]

+ List statistics of tiles



+ Extract one specified tile according to xyz

        mbtextract-xyz [mbtpathIN] [zoom] [x] [y] [tileOUT]

+ Extract tiles according to a xyz file

        mbtextract-list [mbtpathIN] [xyzlist] [tiledirOUT]

+ Extract tiles of spedified zoom level

        mbtextract-z [mbtpathIN] [xyzlist] [tiledirOUT]

+ Extract all tiles

        mbtextract-all [mbtpathIN] [tiledirOUT]

2. Write

+ Create an empty .mbtiles file

+ Import metadata from file

+ Import metadata from another .mbtiles file

+ Import tiles from a folder to .mbtiles file

+ Import tiles from another .mbtiles file

+ Update metadata

+ Update one tile according to xyz

+ Update tiles according to a xyz file

+ Update tiles of specified zoom level

3. Make

ogr2ogr is a good tool for creating MVT .mbtiles, see [GDAL document](https://gdal.org/drivers/vector/mvt.html), it can utilize all cores of your CPU.

4. Serve

consbio/mbtileserver is a good tool for serving mbtiles, see [mbtileserver](https://github.com/consbio/mbtileserver), it can once serve all .mbtiles in a folder. 

5. View

You can view the data with any one or some of belowing tools:

+ QGIS with .mbtiles plugins

QGIS is a **ready to use** desktop GIS, it can run on computers with X86, X64 and ARM CPUs and Linux, Windows, MacOS and other Unix OSs.
see [QGIS website](https://www.qgis.org/).

+ Openlayers

Openlayers is a javascript WebGIS library, there's a [vector tile demo](https://openlayers.org/en/latest/examples/vector-tile-info.html). 
You have to release the xyz tile service first, or extract tiles into a folder.
However nowadays the you can not copy, modify the URL and run the demo directly. I'll post a demo later.

+ Mapbox

Mapbox also provides a WebGIS library, [Mapbox GL](https://github.com/mapbox/mapbox-gl-js). In fact, the full name of the [MVT tile format](https://github.com/mapbox/vector-tile-spec) is *Mapbox Vector Tile*. And even the [mbtiles format](https://github.com/mapbox/mbtiles-spec).
Mapbox has a lot of on-line mapping services, however not all of them are free.

6. Styling

Not in this project!
Modify Mapbox's tool is a practice idea

7. Encryption

Not suggested, and not in consideration. 
You can directly process the tile files, this program will not test if these files are valid.
