# Metadata

## Valid types
  * binary
  * boolean
  * date
  * datetime
  * decimal
  * float
  * geolocation
  * smallint (-32768 to +32767)
  * integer (-2147483648 to +2147483647)
  * biginteger (-9223372036854775808 to 9223372036854775807)
  * json
  * password
  * password_confirmation
  * primary_key
  * string
  * text
  * time
  * timestamp
  * timestamptz
  * uuid

## PropertyDef
  * Name               string      `yaml:"name"`
  * Type               string      `yaml:"type"`
  * Length             int         `yaml:"length"`
  * IsVirtual          bool        `yaml:"isVirtual"`
  * IsKey              bool        `yaml:"isKey"`
  * IsUnique           bool        `yaml:"isUnique"`
  * AdmitNull          bool        `yaml:"admitNull"`
  * Ref                PropertyRef `yaml:"references"`


## PropertyRef
  * Model    string `yaml:"model"`
  * Property string `yaml:"property"`

