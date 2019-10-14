Warning: not ready to use yet!

# mbtmanager
A command line tool for accessing .mbtiles files. 

*In fact, it will be a set of independent executables.*

## Functionality

1. Read a .mbtiles file
+ Show metadata
+ Export metadata
+ List statistics of tiles
+ Extract one specified tile according to xyz
+ Extract tiles according to a xyz file
+ Extract tiles of spedified zoom level
+ Extract all tiles

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

Openlayers is a javascript WebGIS library, there's a [vector tile demo](). 
You have to release the xyz tile service first, or extract tiles into a folder.
However nowadays the you can not copy, modify the URL and run the demo directly. I'll post a demo later.

+ Mapbox

Mapbox also provides a WebGIS library, [Mapbox GL](https://github.com/mapbox/mapbox-gl-js). In fact, the full name of the [MVT tile format](https://github.com/mapbox/vector-tile-spec) is *Mapbox Vector Tile*. And even the [mbtiles format](https://github.com/mapbox/mbtiles-spec).
Mapbox has a lot of on-line mapping services, however not all of them are free.

6. Styling

Not in this project!

