namespace * contentatom.media
namespace java com.gu.contentatom.thrift.atom.media
#@namespace scala com.gu.contentatom.thrift.atom.media

typedef i64 Version

enum Platform {
  YOUTUBE,
  FACEBOOK,
  DAILYMOTION,
  MAINSTREAM,
  URL
}

enum AssetType {
  AUDIO,
  VIDEO
}

enum Category {
    DOCUMENTARY,
    EXPLAINER,
    FEATURE,
    NEWS,
    HOSTED // commercial content supplied by advertiser
}

struct Asset {
  1: required AssetType assetType
  2: required Version version
  3: required string id
  4: required Platform platform
  5: optional string mimeType
}

struct MediaAtom {
  /* the unique ID will be stored in the `atom` data, and this should correspond to the pluto ID */
  2: required list<Asset> assets
  3: optional Version activeVersion
  4: required string title
  5: required Category category
  6: optional string plutoProjectId
  7: optional i64 duration // milliseconds
  8: optional string source
  9: optional string posterUrl
}
