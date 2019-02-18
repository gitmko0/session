package session

import "time"

const defaultSessionKeyName = "sessionid"
const defaultDomain = ""
const defaultExpires = 2 * time.Hour
const defaultGCLifetime = 60 * time.Second
const defaultSessionLifetime int64 = int64(defaultExpires)
const defaultSecure = true
const defaultSessionIDInURLQuery = false
const defaultSessionIDInHTTPHeader = false
const defaultCookieLen uint32 = 32
