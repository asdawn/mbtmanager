*译文仅保留骨架部分，并添加了注释。本程序直接参考MBTiles规范最新版本（1.3）。*
# MBTiles 1.3

## 摘要

MBTiles是实用[SQLite](http://sqlite.org/)存储瓦片地图的一种规范。

MBTiles文件（即**tilesets**）应遵循以下规范以确保兼容性。

## 数据库规范

MBTiles文件应使用SQLite数据库的[version 3.0.0](http://sqlite.org/formatchng.html) 及以上版本。仅允许使用SQLite的核心功能，不允许使用扩展。

MBTiles文件可以使用[官方定义幻数the officially assigned magic number](http://www.sqlite.org/src/artifact?ci=trunk&filename=magic.txt)进行标识，文件头Offset 68，取值0x4d504258表示应用类型为`MBTiles tileset`。

*注释：设置方法为：```PRAGMA application_id =1297105496 ;```*

## 数据库

注意：这里数据库模式采用接口式描述，即只要返回的查询结果符合要求，实际采用表或者视图均可，后续描述不特别区分`表`和`视图`两个术语，统称为`表`。
*注释：无特殊需求时建议仅使用表，否则修改、更新会比较麻烦*

## 字符集

所有`text`类型的字段应使用UTF-8字符编码。

### 元数据部分

#### 模式

数据库必须包含名为 `metadata`的表或视图。

`metadata`严格由两列组成：
+ 字段名`name`，类型`text`
+ 字段名`value`，类型`text`

典型的建表语句：

   ``` CREATE TABLE metadata (name text, value text);```

#### 内容

元数据表采用key/value式存储。

**2个必选属性对：**

* `name`: 瓦片数据集的名称。
* `format`: 瓦片数据的格式，可以为`pbf`, `jpg`, `png`, `webp`或[IETF media type](https://www.iana.org/assignments/media-types/media-types.xhtml)列出的其他类型。

其中，`pbf` 对应于[MMVT瓦片格式](https://github.com/mapbox/vector-tile-spec/)，是一种GZIP压缩的矢量瓦片格式。

**4个应该包含（强烈建议）的属性对：**

* `bounds`: WGS坐标系下的瓦片边界，采用OpenLayers边界格式表述(左, 下, 右, 上)。例如，全球的范围表示为：`-180.0,-85,180,85`。

*注释：由于Web墨卡托坐标系的限制，普通地图纬度的极限大致在南北纬85度*

* `center` 默认视图中心点，由经度、维度、缩放级别组成，例如`-122.1906,37.7599,11`。
* `minzoom` : 存储的瓦片数据的最小缩放级别。
* `maxzoom` : 存储的瓦片数据的最大缩放级别。

*注释：缩放级zoom即xyz瓦片编码中的z*

**可能包含的属性对：**

* `attribution` : 地图或样式的版权/归属说明，应为HTML字符串。
* `description` : 地图内容描述。
* `type` : 地图类型，取值为`overlay` 或 `baselayer`，分别对应于覆盖物和底图。
* `version` : 瓦片地图数据的版本号（数字）

**当格式为`pbf`时, `metadata` 表必须包含属性对:**

* `json`: JSON字符串，列出矢量瓦片中的图层以及图层中要素属性的名称与类型， 详见 [下文](#矢量瓦片元数据) 。

`metadata` 表允许包含符合[UTFGrid-based interaction](https://github.com/mapbox/utfgrid-spec) 的属性对，以及用于其他目的属性对。

### 瓦片部分

#### 模式

数据库必须包含名为`tiles`的表。

该表必须包含：
+ 字段名`zoom_level`， 类型`integer`
+ 字段名`tile_column`， 类型`integer`
+ 字段名`tile_row`，类型`integer`
+ 字段名`tile_data`，类型`blob`

典型的`tiles`表建表语句：

 ```CREATE TABLE tiles (zoom_level integer, tile_column integer, tile_row integer, tile_data blob);```

为提高瓦片访问效率，该表可以带有索引：

```CREATE UNIQUE INDEX tile_index on tiles (zoom_level, tile_column, tile_row);```

*注释：强烈建议添加索引*

#### 内容

`zoom_level`、`tile_column`和`tile_row`三列是瓦片的编号，必须遵循[OSGEO瓦片地图服务规范](http://wiki.osgeo.org/wiki/Tile_Map_Service_Specification)，使用[global-mercator](http://wiki.osgeo.org/wiki/Tile_Map_Service_Specification#global-mercator) (aka Spherical Mercator) 配置。

注意在TMS瓦片地图模式中，Y轴方向与Web地图常用的XYZ瓦片编码相反，因此我们通常使用的编码为11/327/791的瓦片存储时
`zoom_level`=11，`tile_column`=327，但是`tile_row`=1256（1256 = 2^11 - 1 - 791）。

*注释：XYZ瓦片编码即谷歌地图式编码，路径一般为Z/X/Y形式，Z即缩放级别，X为列号，Y为行号*

`tile_data`必须以`blob`形式存储瓦片影像或矢量瓦片的原始数据。

*注释：即每条记录存储一个瓦片，数据用`tile_data`，编号用`zoom_level`、`tile_column`和`tile_row`*


**UTFGrid（即属性值瓦片格式）部分暂时跳过！**


## 矢量瓦片元数据

如上文所述，MBTiles文件用于矢量瓦片存储时，`metadata`表必须有`name`为`json`的记录，用UTF-8编码的JSON字符串存储瓦片涉及的图层和属性的基本信息。

### 矢量图层

存储矢量瓦片的MBTiles文件的`metadata`表中，`name`为`json`的记录取值为一个JSON字符串，该JSON必须包含一个名为`vector_layers`，取值为JSON对象数据的键值对。对象数据中，每个对象为一个矢量图层的描述信息，并且必须包含以下键值对：

* `id` (string): The layer ID, which is referred to as the `name` of the layer in the [Mapbox Vector Tile spec](https://github.com/mapbox/vector-tile-spec/tree/master/2.1#41-layers).
* `fields`: 存储着图层中属性名称、类型键值对的JSON对象。类型为`"Number"`、`"Boolean"`或者`"String"`。可变类型标记为`"String"`。

每个图层对象还可以包含以下键值对：

* `description`: 图层的描述文本

每个图层对象还可以包含以下键值对：

* `minzoom` : 图层有瓦片数据的最小缩放级别
* `maxzoom`: 图层有瓦片数据的最大缩放级别

任意图层的`minzoom` >= MBTiles文件的`minzoom`，
任意图层的`maxzoom` <= MBTiles文件的`maxzoom`。

这些键值对用来应对同一个MBTiles文件中不同图层有数据的缩放级别也不同的情况，例如"minor roads"图层仅在较高的缩放级别上有数据。

**注释：说的这么麻烦，看看下边的示例就明白了，JSON字符串的大致结构是**
```
{  
   //矢量瓦片必选的vector_layers
   "vector_layers": [
      //第一个图层的描述
      {    
         "id": "图层的ID",
         "description": 图层描述信息,
         "minzoom": 图层瓦片数据的最小缩放级别,
         "maxzoom": 图层瓦片数据的最大缩放级别,
         "fields": {
            "属性1名称": "属性1类型",
            //其余属性的名称:类型
          }
       },
      //图层2、3、4...的描述
   ],
   //可选的tilestates
   "tilestats": {......}
}
``` 
**注释：矢量瓦片允许包含多个图层，不过本程序不关心瓦片内容，只关心瓦片的存储与整个MBTiles的Metadata，且只考虑单图层矢量瓦片**。

### Tilestats

`json`记录对应的`value`（JSON字符串）中，允许包含一个名为`tilestats`的JSON键/值对，其取值为一个"geostats"式对象（参见 [mapbox-geostats](https://github.com/mapbox/mapbox-geostats#output-the-stats)）。与`vector_layers`类似, 它列出了瓦片数据集涉及的图层、各图层的属性，但额外为追加了属性的范例值以及数值属性的取值范围。

**注释：这个JSON字符串中，"vector_layers"必选，"geostats"可选，后者的内容更详细**


### 示例

包含美国县界及主要道路的矢量瓦片数据（源于[TIGER](https://www.census.gov/geo/maps-data/data/tiger-line.html)） 的`metadata`表的内容如下：

*注释：美国的县级别差不多相当于中国的地级市。反正页面打不开，只看看格式就好了*

* `name`: `TIGER 2016`
* `format`: `pbf`
* `bounds`: `-179.231086,-14.601813,179.859681,71.441059`
* `center`: `-84.375000,36.466030,5`
* `minzoom`: `0`
* `maxzoom`: `5`
* `attribution`: `United States Census`
* `description`: `US Census counties and primary roads`
* `type`: `overlay`
* `version`: `2`
* `json`:
```
    {
        "vector_layers": [
            {
                "id": "tl_2016_us_county",
                "description": "Census counties",
                "minzoom": 0,
                "maxzoom": 5,
                "fields": {
                    "ALAND": "Number",
                    "AWATER": "Number",
                    "GEOID": "String",
                    "MTFCC": "String",
                    "NAME": "String"
                }
            },
            {
                "id": "tl_2016_us_primaryroads",
                "description": "Census primary roads",
                "minzoom": 0,
                "maxzoom": 5,
                "fields": {
                    "FULLNAME": "String",
                    "LINEARID": "String",
                    "MTFCC": "String",
                    "RTTYP": "String"
                }
            }
        ],
        "tilestats": {
            "layerCount": 2,
            "layers": [
                {
                    "layer": "tl_2016_us_county",
                    "count": 3233,
                    "geometry": "Polygon",
                    "attributeCount": 5,
                    "attributes": [
                        {
                            "attribute": "ALAND",
                            "count": 6,
                            "type": "number",
                            "values": [
                                1000508839,
                                1001065264,
                                1001787870,
                                1002071716,
                                1002509543,
                                1003451714
                            ],
                            "min": 82093,
                            "max": 376825063576
                        },
                        {
                            "attribute": "AWATER",
                            "count": 6,
                            "type": "number",
                            "values": [
                                0,
                                100091246,
                                10017651,
                                100334057,
                                10040117,
                                1004128585
                            ],
                            "min": 0,
                            "max": 25190628850
                        },
                        {
                            "attribute": "GEOID",
                            "count": 6,
                            "type": "string",
                            "values": [
                                "01001",
                                "01003",
                                "01005",
                                "01007",
                                "01009",
                                "01011"
                            ]
                        },
                        {
                            "attribute": "MTFCC",
                            "count": 1,
                            "type": "string",
                            "values": [
                                "G4020"
                            ]
                        },
                        {
                            "attribute": "NAME",
                            "count": 6,
                            "type": "string",
                            "values": [
                                "Abbeville",
                                "Acadia",
                                "Accomack",
                                "Ada",
                                "Adair",
                                "Adams"
                            ]
                        }
                    ]
                },
                {
                    "layer": "tl_2016_us_primaryroads",
                    "count": 12509,
                    "geometry": "LineString",
                    "attributeCount": 4,
                    "attributes": [
                        {
                            "attribute": "FULLNAME",
                            "count": 6,
                            "type": "string",
                            "values": [
                                "1- 80",
                                "10",
                                "10-Hov Fwy",
                                "12th St",
                                "14 Th St",
                                "17th St NE"
                            ]
                        },
                        {
                            "attribute": "LINEARID",
                            "count": 6,
                            "type": "string",
                            "values": [
                                "1101000363000",
                                "1101000363004",
                                "1101019172643",
                                "1101019172644",
                                "1101019172674",
                                "1101019172675"
                            ]
                        },
                        {
                            "attribute": "MTFCC",
                            "count": 1,
                            "type": "string",
                            "values": [
                                "S1100"
                            ]
                        },
                        {
                            "attribute": "RTTYP",
                            "count": 6,
                            "type": "string",
                            "values": [
                                "C",
                                "I",
                                "M",
                                "O",
                                "S",
                                "U"
                            ]
                        }
                    ]
                }
            ]
        }
    }

```

## 未来的改进方向

未来版本中，`metadata`表可能新增一条`name`为 `compression`的记录，用于标明瓦片数据的压缩格式。

未来版本中，`metadata`中`name`为`bounds`、`minzoom`和`maxzoom`的属性对将不再是强烈推荐，而是必选。

未来版本中，`metadata`中`name`为`json` 的记录的功能将由外部定义替代。
