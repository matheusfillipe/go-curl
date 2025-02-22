//go:generate /usr/bin/env python ./misc/codegen.py

package curl
/*
#include <curl/curl.h>
#include "compat.h"
*/
import "C"

// CURLcode
const (
	E_UNSUPPORTED_PROTOCOL    = C.CURLE_UNSUPPORTED_PROTOCOL
	E_FAILED_INIT             = C.CURLE_FAILED_INIT
	E_URL_MALFORMAT           = C.CURLE_URL_MALFORMAT
	E_URL_MALFORMAT_USER      = C.CURLE_URL_MALFORMAT_USER
	E_COULDNT_RESOLVE_PROXY   = C.CURLE_COULDNT_RESOLVE_PROXY
	E_COULDNT_RESOLVE_HOST    = C.CURLE_COULDNT_RESOLVE_HOST
	E_COULDNT_CONNECT         = C.CURLE_COULDNT_CONNECT
	E_FTP_WEIRD_SERVER_REPLY  = C.CURLE_FTP_WEIRD_SERVER_REPLY
	E_FTP_ACCESS_DENIED       = C.CURLE_FTP_ACCESS_DENIED
	E_FTP_USER_PASSWORD_INCORRECT = C.CURLE_FTP_USER_PASSWORD_INCORRECT
	E_FTP_WEIRD_PASS_REPLY    = C.CURLE_FTP_WEIRD_PASS_REPLY
	E_FTP_WEIRD_USER_REPLY    = C.CURLE_FTP_WEIRD_USER_REPLY
	E_FTP_WEIRD_PASV_REPLY    = C.CURLE_FTP_WEIRD_PASV_REPLY
	E_FTP_WEIRD_227_FORMAT    = C.CURLE_FTP_WEIRD_227_FORMAT
	E_FTP_CANT_GET_HOST       = C.CURLE_FTP_CANT_GET_HOST
	E_FTP_CANT_RECONNECT      = C.CURLE_FTP_CANT_RECONNECT
	E_FTP_COULDNT_SET_BINARY  = C.CURLE_FTP_COULDNT_SET_BINARY
	E_PARTIAL_FILE            = C.CURLE_PARTIAL_FILE
	E_FTP_COULDNT_RETR_FILE   = C.CURLE_FTP_COULDNT_RETR_FILE
	E_FTP_WRITE_ERROR         = C.CURLE_FTP_WRITE_ERROR
	E_FTP_QUOTE_ERROR         = C.CURLE_FTP_QUOTE_ERROR
	E_HTTP_RETURNED_ERROR     = C.CURLE_HTTP_RETURNED_ERROR
	E_WRITE_ERROR             = C.CURLE_WRITE_ERROR
	E_MALFORMAT_USER          = C.CURLE_MALFORMAT_USER
	E_FTP_COULDNT_STOR_FILE   = C.CURLE_FTP_COULDNT_STOR_FILE
	E_READ_ERROR              = C.CURLE_READ_ERROR
	E_OUT_OF_MEMORY           = C.CURLE_OUT_OF_MEMORY
	E_OPERATION_TIMEOUTED     = C.CURLE_OPERATION_TIMEOUTED
	E_FTP_COULDNT_SET_ASCII   = C.CURLE_FTP_COULDNT_SET_ASCII
	E_FTP_PORT_FAILED         = C.CURLE_FTP_PORT_FAILED
	E_FTP_COULDNT_USE_REST    = C.CURLE_FTP_COULDNT_USE_REST
	E_FTP_COULDNT_GET_SIZE    = C.CURLE_FTP_COULDNT_GET_SIZE
	E_HTTP_RANGE_ERROR        = C.CURLE_HTTP_RANGE_ERROR
	E_HTTP_POST_ERROR         = C.CURLE_HTTP_POST_ERROR
	E_SSL_CONNECT_ERROR       = C.CURLE_SSL_CONNECT_ERROR
	E_BAD_DOWNLOAD_RESUME     = C.CURLE_BAD_DOWNLOAD_RESUME
	E_FILE_COULDNT_READ_FILE  = C.CURLE_FILE_COULDNT_READ_FILE
	E_LDAP_CANNOT_BIND        = C.CURLE_LDAP_CANNOT_BIND
	E_LDAP_SEARCH_FAILED      = C.CURLE_LDAP_SEARCH_FAILED
	E_LIBRARY_NOT_FOUND       = C.CURLE_LIBRARY_NOT_FOUND
	E_FUNCTION_NOT_FOUND      = C.CURLE_FUNCTION_NOT_FOUND
	E_ABORTED_BY_CALLBACK     = C.CURLE_ABORTED_BY_CALLBACK
	E_BAD_FUNCTION_ARGUMENT   = C.CURLE_BAD_FUNCTION_ARGUMENT
	E_BAD_CALLING_ORDER       = C.CURLE_BAD_CALLING_ORDER
	E_INTERFACE_FAILED        = C.CURLE_INTERFACE_FAILED
	E_BAD_PASSWORD_ENTERED    = C.CURLE_BAD_PASSWORD_ENTERED
	E_UNKNOWN_TELNET_OPTION   = C.CURLE_UNKNOWN_TELNET_OPTION
	E_OBSOLETE                = C.CURLE_OBSOLETE
	E_SSL_PEER_CERTIFICATE    = C.CURLE_SSL_PEER_CERTIFICATE
	E_GOT_NOTHING             = C.CURLE_GOT_NOTHING
	E_SSL_ENGINE_NOTFOUND     = C.CURLE_SSL_ENGINE_NOTFOUND
	E_SSL_ENGINE_SETFAILED    = C.CURLE_SSL_ENGINE_SETFAILED
	E_SEND_ERROR              = C.CURLE_SEND_ERROR
	E_RECV_ERROR              = C.CURLE_RECV_ERROR
	E_SHARE_IN_USE            = C.CURLE_SHARE_IN_USE
	E_SSL_CERTPROBLEM         = C.CURLE_SSL_CERTPROBLEM
	E_SSL_CIPHER              = C.CURLE_SSL_CIPHER
	E_SSL_CACERT              = C.CURLE_SSL_CACERT
	E_BAD_CONTENT_ENCODING    = C.CURLE_BAD_CONTENT_ENCODING
	E_LDAP_INVALID_URL        = C.CURLE_LDAP_INVALID_URL
	E_FILESIZE_EXCEEDED       = C.CURLE_FILESIZE_EXCEEDED
	E_FTP_SSL_FAILED          = C.CURLE_FTP_SSL_FAILED
	E_SEND_FAIL_REWIND        = C.CURLE_SEND_FAIL_REWIND
	E_SSL_ENGINE_INITFAILED   = C.CURLE_SSL_ENGINE_INITFAILED
	E_LOGIN_DENIED            = C.CURLE_LOGIN_DENIED
	E_TFTP_NOTFOUND           = C.CURLE_TFTP_NOTFOUND
	E_TFTP_PERM               = C.CURLE_TFTP_PERM
	E_TFTP_DISKFULL           = C.CURLE_TFTP_DISKFULL
	E_TFTP_ILLEGAL            = C.CURLE_TFTP_ILLEGAL
	E_TFTP_UNKNOWNID          = C.CURLE_TFTP_UNKNOWNID
	E_TFTP_EXISTS             = C.CURLE_TFTP_EXISTS
	E_TFTP_NOSUCHUSER         = C.CURLE_TFTP_NOSUCHUSER
	E_CONV_FAILED             = C.CURLE_CONV_FAILED
	E_CONV_REQD               = C.CURLE_CONV_REQD
	E_SSL_CACERT_BADFILE      = C.CURLE_SSL_CACERT_BADFILE
	E_OPERATION_TIMEDOUT      = C.CURLE_OPERATION_TIMEDOUT
	E_HTTP_NOT_FOUND          = C.CURLE_HTTP_NOT_FOUND
	E_HTTP_PORT_FAILED        = C.CURLE_HTTP_PORT_FAILED
	E_ALREADY_COMPLETE        = C.CURLE_ALREADY_COMPLETE
	E_FTP_PARTIAL_FILE        = C.CURLE_FTP_PARTIAL_FILE
	E_FTP_BAD_DOWNLOAD_RESUME = C.CURLE_FTP_BAD_DOWNLOAD_RESUME
)

// easy.Setopt(flag, ...)
const (
	OPT_FILE                      = C.CURLOPT_FILE
	OPT_URL                       = C.CURLOPT_URL
	OPT_PORT                      = C.CURLOPT_PORT
	OPT_PROXY                     = C.CURLOPT_PROXY
	OPT_USERPWD                   = C.CURLOPT_USERPWD
	OPT_PROXYUSERPWD              = C.CURLOPT_PROXYUSERPWD
	OPT_RANGE                     = C.CURLOPT_RANGE
	OPT_INFILE                    = C.CURLOPT_INFILE
	OPT_ERRORBUFFER               = C.CURLOPT_ERRORBUFFER
	OPT_WRITEFUNCTION             = C.CURLOPT_WRITEFUNCTION
	OPT_READFUNCTION              = C.CURLOPT_READFUNCTION
	OPT_TIMEOUT                   = C.CURLOPT_TIMEOUT
	OPT_INFILESIZE                = C.CURLOPT_INFILESIZE
	OPT_POSTFIELDS                = C.CURLOPT_POSTFIELDS
	OPT_REFERER                   = C.CURLOPT_REFERER
	OPT_FTPPORT                   = C.CURLOPT_FTPPORT
	OPT_USERAGENT                 = C.CURLOPT_USERAGENT
	OPT_LOW_SPEED_TIME            = C.CURLOPT_LOW_SPEED_TIME
	OPT_RESUME_FROM               = C.CURLOPT_RESUME_FROM
	OPT_COOKIE                    = C.CURLOPT_COOKIE
	OPT_HTTPHEADER                = C.CURLOPT_HTTPHEADER
	OPT_HTTPPOST                  = C.CURLOPT_HTTPPOST
	OPT_SSLCERT                   = C.CURLOPT_SSLCERT
	OPT_SSLCERTPASSWD             = C.CURLOPT_SSLCERTPASSWD
	OPT_SSLKEYPASSWD              = C.CURLOPT_SSLKEYPASSWD
	OPT_CRLF                      = C.CURLOPT_CRLF
	OPT_QUOTE                     = C.CURLOPT_QUOTE
	OPT_WRITEHEADER               = C.CURLOPT_WRITEHEADER
	OPT_COOKIEFILE                = C.CURLOPT_COOKIEFILE
	OPT_SSLVERSION                = C.CURLOPT_SSLVERSION
	OPT_TIMECONDITION             = C.CURLOPT_TIMECONDITION
	OPT_TIMEVALUE                 = C.CURLOPT_TIMEVALUE
	OPT_CUSTOMREQUEST             = C.CURLOPT_CUSTOMREQUEST
	OPT_STDERR                    = C.CURLOPT_STDERR
	OPT_POSTQUOTE                 = C.CURLOPT_POSTQUOTE
	OPT_WRITEINFO                 = C.CURLOPT_WRITEINFO
	OPT_VERBOSE                   = C.CURLOPT_VERBOSE
	OPT_HEADER                    = C.CURLOPT_HEADER
	OPT_NOPROGRESS                = C.CURLOPT_NOPROGRESS
	OPT_NOBODY                    = C.CURLOPT_NOBODY
	OPT_FAILONERROR               = C.CURLOPT_FAILONERROR
	OPT_UPLOAD                    = C.CURLOPT_UPLOAD
	OPT_POST                      = C.CURLOPT_POST
	OPT_FTPLISTONLY               = C.CURLOPT_FTPLISTONLY
	OPT_FTPAPPEND                 = C.CURLOPT_FTPAPPEND
	OPT_NETRC                     = C.CURLOPT_NETRC
	OPT_FOLLOWLOCATION            = C.CURLOPT_FOLLOWLOCATION
	OPT_TRANSFERTEXT              = C.CURLOPT_TRANSFERTEXT
	OPT_PUT                       = C.CURLOPT_PUT
	OPT_PROGRESSFUNCTION          = C.CURLOPT_PROGRESSFUNCTION
	OPT_PROGRESSDATA              = C.CURLOPT_PROGRESSDATA
	OPT_AUTOREFERER               = C.CURLOPT_AUTOREFERER
	OPT_PROXYPORT                 = C.CURLOPT_PROXYPORT
	OPT_POSTFIELDSIZE             = C.CURLOPT_POSTFIELDSIZE
	OPT_HTTPPROXYTUNNEL           = C.CURLOPT_HTTPPROXYTUNNEL
	OPT_INTERFACE                 = C.CURLOPT_INTERFACE
	OPT_KRB4LEVEL                 = C.CURLOPT_KRB4LEVEL
	OPT_SSL_VERIFYPEER            = C.CURLOPT_SSL_VERIFYPEER
	OPT_CAINFO                    = C.CURLOPT_CAINFO
	OPT_MAXREDIRS                 = C.CURLOPT_MAXREDIRS
	OPT_FILETIME                  = C.CURLOPT_FILETIME
	OPT_TELNETOPTIONS             = C.CURLOPT_TELNETOPTIONS
	OPT_MAXCONNECTS               = C.CURLOPT_MAXCONNECTS
	OPT_CLOSEPOLICY               = C.CURLOPT_CLOSEPOLICY
	OPT_FRESH_CONNECT             = C.CURLOPT_FRESH_CONNECT
	OPT_FORBID_REUSE              = C.CURLOPT_FORBID_REUSE
	OPT_RANDOM_FILE               = C.CURLOPT_RANDOM_FILE
	OPT_EGDSOCKET                 = C.CURLOPT_EGDSOCKET
	OPT_CONNECTTIMEOUT            = C.CURLOPT_CONNECTTIMEOUT
	OPT_HEADERFUNCTION            = C.CURLOPT_HEADERFUNCTION
	OPT_HTTPGET                   = C.CURLOPT_HTTPGET
	OPT_SSL_VERIFYHOST            = C.CURLOPT_SSL_VERIFYHOST
	OPT_COOKIEJAR                 = C.CURLOPT_COOKIEJAR
	OPT_SSL_CIPHER_LIST           = C.CURLOPT_SSL_CIPHER_LIST
	OPT_HTTP_VERSION              = C.CURLOPT_HTTP_VERSION
	OPT_FTP_USE_EPSV              = C.CURLOPT_FTP_USE_EPSV
	OPT_SSLCERTTYPE               = C.CURLOPT_SSLCERTTYPE
	OPT_SSLKEY                    = C.CURLOPT_SSLKEY
	OPT_SSLKEYTYPE                = C.CURLOPT_SSLKEYTYPE
	OPT_SSLENGINE                 = C.CURLOPT_SSLENGINE
	OPT_SSLENGINE_DEFAULT         = C.CURLOPT_SSLENGINE_DEFAULT
	OPT_DNS_USE_GLOBAL_CACHE      = C.CURLOPT_DNS_USE_GLOBAL_CACHE
	OPT_DNS_CACHE_TIMEOUT         = C.CURLOPT_DNS_CACHE_TIMEOUT
	OPT_PREQUOTE                  = C.CURLOPT_PREQUOTE
	OPT_DEBUGFUNCTION             = C.CURLOPT_DEBUGFUNCTION
	OPT_DEBUGDATA                 = C.CURLOPT_DEBUGDATA
	OPT_COOKIESESSION             = C.CURLOPT_COOKIESESSION
	OPT_CAPATH                    = C.CURLOPT_CAPATH
	OPT_BUFFERSIZE                = C.CURLOPT_BUFFERSIZE
	OPT_NOSIGNAL                  = C.CURLOPT_NOSIGNAL
	OPT_SHARE                     = C.CURLOPT_SHARE
	OPT_PROXYTYPE                 = C.CURLOPT_PROXYTYPE
	OPT_ENCODING                  = C.CURLOPT_ENCODING
	OPT_PRIVATE                   = C.CURLOPT_PRIVATE
	OPT_HTTP200ALIASES            = C.CURLOPT_HTTP200ALIASES
	OPT_UNRESTRICTED_AUTH         = C.CURLOPT_UNRESTRICTED_AUTH
	OPT_FTP_USE_EPRT              = C.CURLOPT_FTP_USE_EPRT
	OPT_HTTPAUTH                  = C.CURLOPT_HTTPAUTH
	OPT_SSL_CTX_FUNCTION          = C.CURLOPT_SSL_CTX_FUNCTION
	OPT_SSL_CTX_DATA              = C.CURLOPT_SSL_CTX_DATA
	OPT_FTP_CREATE_MISSING_DIRS   = C.CURLOPT_FTP_CREATE_MISSING_DIRS
	OPT_PROXYAUTH                 = C.CURLOPT_PROXYAUTH
	OPT_IPRESOLVE                 = C.CURLOPT_IPRESOLVE
	OPT_RESOLVE                   = C.CURLOPT_RESOLVE
	OPT_MAXFILESIZE               = C.CURLOPT_MAXFILESIZE
	OPT_INFILESIZE_LARGE          = C.CURLOPT_INFILESIZE_LARGE
	OPT_RESUME_FROM_LARGE         = C.CURLOPT_RESUME_FROM_LARGE
	OPT_MAXFILESIZE_LARGE         = C.CURLOPT_MAXFILESIZE_LARGE
	OPT_NETRC_FILE                = C.CURLOPT_NETRC_FILE
	OPT_FTP_SSL                   = C.CURLOPT_FTP_SSL
	OPT_POSTFIELDSIZE_LARGE       = C.CURLOPT_POSTFIELDSIZE_LARGE
	OPT_TCP_NODELAY               = C.CURLOPT_TCP_NODELAY
	OPT_FTPSSLAUTH                = C.CURLOPT_FTPSSLAUTH
	OPT_IOCTLFUNCTION             = C.CURLOPT_IOCTLFUNCTION
	OPT_IOCTLDATA                 = C.CURLOPT_IOCTLDATA
	OPT_FTP_ACCOUNT               = C.CURLOPT_FTP_ACCOUNT
	OPT_COOKIELIST                = C.CURLOPT_COOKIELIST
	OPT_IGNORE_CONTENT_LENGTH     = C.CURLOPT_IGNORE_CONTENT_LENGTH
	OPT_FTP_SKIP_PASV_IP          = C.CURLOPT_FTP_SKIP_PASV_IP
	OPT_FTP_FILEMETHOD            = C.CURLOPT_FTP_FILEMETHOD
	OPT_LOCALPORT                 = C.CURLOPT_LOCALPORT
	OPT_LOCALPORTRANGE            = C.CURLOPT_LOCALPORTRANGE
	OPT_CONNECT_ONLY              = C.CURLOPT_CONNECT_ONLY
	OPT_CONV_FROM_NETWORK_FUNCTION = C.CURLOPT_CONV_FROM_NETWORK_FUNCTION
	OPT_CONV_TO_NETWORK_FUNCTION  = C.CURLOPT_CONV_TO_NETWORK_FUNCTION
	OPT_CONV_FROM_UTF8_FUNCTION   = C.CURLOPT_CONV_FROM_UTF8_FUNCTION
	OPT_MAX_SEND_SPEED_LARGE      = C.CURLOPT_MAX_SEND_SPEED_LARGE
	OPT_MAX_RECV_SPEED_LARGE      = C.CURLOPT_MAX_RECV_SPEED_LARGE
	OPT_FTP_ALTERNATIVE_TO_USER   = C.CURLOPT_FTP_ALTERNATIVE_TO_USER
	OPT_SOCKOPTFUNCTION           = C.CURLOPT_SOCKOPTFUNCTION
	OPT_SOCKOPTDATA               = C.CURLOPT_SOCKOPTDATA
	OPT_SSL_SESSIONID_CACHE       = C.CURLOPT_SSL_SESSIONID_CACHE
	OPT_WRITEDATA                 = C.CURLOPT_WRITEDATA
	OPT_READDATA                  = C.CURLOPT_READDATA
	OPT_HEADERDATA                = C.CURLOPT_HEADERDATA
)

// easy.Getinfo(flag)
const (
	INFO_TEXT                 = C.CURLINFO_TEXT
	INFO_EFFECTIVE_URL        = C.CURLINFO_EFFECTIVE_URL
	INFO_RESPONSE_CODE        = C.CURLINFO_RESPONSE_CODE
	INFO_TOTAL_TIME           = C.CURLINFO_TOTAL_TIME
	INFO_NAMELOOKUP_TIME      = C.CURLINFO_NAMELOOKUP_TIME
	INFO_CONNECT_TIME         = C.CURLINFO_CONNECT_TIME
	INFO_PRETRANSFER_TIME     = C.CURLINFO_PRETRANSFER_TIME
	INFO_SIZE_UPLOAD          = C.CURLINFO_SIZE_UPLOAD
	INFO_SIZE_DOWNLOAD        = C.CURLINFO_SIZE_DOWNLOAD
	INFO_SPEED_DOWNLOAD       = C.CURLINFO_SPEED_DOWNLOAD
	INFO_SPEED_UPLOAD         = C.CURLINFO_SPEED_UPLOAD
	INFO_HEADER_SIZE          = C.CURLINFO_HEADER_SIZE
	INFO_REQUEST_SIZE         = C.CURLINFO_REQUEST_SIZE
	INFO_SSL_VERIFYRESULT     = C.CURLINFO_SSL_VERIFYRESULT
	INFO_FILETIME             = C.CURLINFO_FILETIME
	INFO_CONTENT_LENGTH_DOWNLOAD = C.CURLINFO_CONTENT_LENGTH_DOWNLOAD
	INFO_CONTENT_LENGTH_UPLOAD = C.CURLINFO_CONTENT_LENGTH_UPLOAD
	INFO_STARTTRANSFER_TIME   = C.CURLINFO_STARTTRANSFER_TIME
	INFO_CONTENT_TYPE         = C.CURLINFO_CONTENT_TYPE
	INFO_REDIRECT_TIME        = C.CURLINFO_REDIRECT_TIME
	INFO_REDIRECT_COUNT       = C.CURLINFO_REDIRECT_COUNT
	INFO_PRIVATE              = C.CURLINFO_PRIVATE
	INFO_HTTP_CONNECTCODE     = C.CURLINFO_HTTP_CONNECTCODE
	INFO_HTTPAUTH_AVAIL       = C.CURLINFO_HTTPAUTH_AVAIL
	INFO_PROXYAUTH_AVAIL      = C.CURLINFO_PROXYAUTH_AVAIL
	INFO_OS_ERRNO             = C.CURLINFO_OS_ERRNO
	INFO_NUM_CONNECTS         = C.CURLINFO_NUM_CONNECTS
	INFO_SSL_ENGINES          = C.CURLINFO_SSL_ENGINES
	INFO_COOKIELIST           = C.CURLINFO_COOKIELIST
	INFO_LASTSOCKET           = C.CURLINFO_LASTSOCKET
	INFO_FTP_ENTRY_PATH       = C.CURLINFO_FTP_ENTRY_PATH
	INFO_LASTONE              = C.CURLINFO_LASTONE
	INFO_HTTP_CODE            = C.CURLINFO_HTTP_CODE
)

// Auth
const (
	AUTH_NONE                      = C.CURLAUTH_NONE & (1<<32 - 1)
	AUTH_BASIC                     = C.CURLAUTH_BASIC & (1<<32 - 1)
	AUTH_DIGEST                    = C.CURLAUTH_DIGEST & (1<<32 - 1)
	AUTH_GSSNEGOTIATE              = C.CURLAUTH_GSSNEGOTIATE & (1<<32 - 1)
	AUTH_NTLM                      = C.CURLAUTH_NTLM & (1<<32 - 1)
	AUTH_ANY                       = C.CURLAUTH_ANY & (1<<32 - 1)
	AUTH_ANYSAFE                   = C.CURLAUTH_ANYSAFE & (1<<32 - 1)
)

// generated ends
