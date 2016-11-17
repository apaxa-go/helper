package mimeh

// Common WEB MIME types.
const (
	MimeTextHtmlUtf8          = "text/html; charset=UTF-8" // HTML, ".html"
	MimeTextCssUtf8           = "text/css; charset=UTF-8"  // CSS, ".css"
	MimeTextLessUtf8          = "text/css; charset=UTF-8"  // Less (http://lesscss.org/), ".less", not sureabout mime type, but will good for development
	MimeApplicationJavascript = "application/javascript"   // JavaScript, ".js"
	MimeApplicationVndDart    = "application/vnd-dart"     // DartLang, ".dart"
)

// Fonts MIME types.
const (
	// TODO someday fonts mime types may be change - http://dev.w3.org/webfonts/WOFF2/spec/#IMT
	MimeApplicationVndMsFontObject = "application/vnd.ms-fontobject" // Embedded OpenType font, ".eot"
	MimeApplicationXFontTtf        = "application/x-font-ttf"        // TrueType font, ".ttf"
	MimeApplicationXFontWoff       = "application/x-font-woff"       // Web Open Font Format, ".woff"
	MimeFontWoff2                  = "font/woff2"                    // Web Open Font Format 2, ".woff2", currently there is not mime type for woff2, using draft
)

// Images MIME types.
const (
	MimeImageSvgXml = "image/svg+xml" // Scalable Vector Graphics, ".svg"
	MimeImageJpeg   = "image/jpeg"    // Joint Photographic Experts Group image, ".jpeg"
	MimeImagePng    = "image/png"     // Portable Network Graphics, ".png"
	MimeImageGif    = "image/gif"     // Graphics Interchange Format, ".gif"
)

// Arbitrary data MIME type.
const MimeApplicationOctetStream = "application/octet-stream" // Any binary data

// Shorthands

// Common WEB MIME types shorthands.
const (
	MimeHtml = MimeTextHtmlUtf8
	MimeCss  = MimeTextCssUtf8
	MimeLess = MimeTextLessUtf8
	MimeJs   = MimeApplicationJavascript
	MimeDart = MimeApplicationVndDart
)

// Fonts MIME types shorthands.
const (
	MimeFontEot  = MimeApplicationVndMsFontObject
	MimeFontTtf  = MimeApplicationXFontTtf
	MimeFontWoff = MimeApplicationXFontWoff
)

// Images MIME types shorthands.
const (
	MimeImageSvg = MimeImageSvgXml
)

// Arbitary data MIME type shorthands.
const (
	MimeBinary   = MimeApplicationOctetStream
	MimeFallback = MimeApplicationOctetStream
)
