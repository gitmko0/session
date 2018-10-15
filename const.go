package session

import "time"

const version = "v1.0.0"

const defaultSessionKeyName = "sessionid"
const defaultDomain = ""
const defaultExpires = 2 * time.Hour
const defaultGCLifetime = 30 * time.Second
const defaultSessionLifetime int64 = 60
const defaultSecure = true
const defaultSessionIDInURLQuery = false
const defaultSessionIDInHTTPHeader = false
const defaultCookieLen uint32 = 32

const cookieCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
