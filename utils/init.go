package utils

import (
	"github.com/dengpju/higo-utils/utils/convutil"
	"github.com/dengpju/higo-utils/utils/dirutil"
	"github.com/dengpju/higo-utils/utils/encodeutil"
	"github.com/dengpju/higo-utils/utils/fileutil"
	"github.com/dengpju/higo-utils/utils/gobutil"
	"github.com/dengpju/higo-utils/utils/hostutil"
	"github.com/dengpju/higo-utils/utils/maputil"
	"github.com/dengpju/higo-utils/utils/phputil"
	"github.com/dengpju/higo-utils/utils/protoutil"
	"github.com/dengpju/higo-utils/utils/randomutil"
	"github.com/dengpju/higo-utils/utils/rsautil"
	"github.com/dengpju/higo-utils/utils/runtimeutil"
	"github.com/dengpju/higo-utils/utils/sliceutil"
	"github.com/dengpju/higo-utils/utils/stringutil"
	"github.com/dengpju/higo-utils/utils/timeutil"
	"github.com/dengpju/higo-utils/utils/tlsutil"
	"github.com/dengpju/higo-utils/utils/tokenutil"
	"github.com/dengpju/higo-utils/utils/ufuncutil"
)

var (
	Convert = &convert{}
	Dir     = &dir{}
	Encode  = &encode{}
	File    = &file{}
	Gob     = &gob{}
	Host    = &host{}
	Map     = &maps{}
	PHP     = &php{}
	Proto   = &php{}
	Random  = &random{}
	Rsa     = &rsa{}
	Runtime = &runtime{}
	Slice   = &sliceu{}
	String  = &str{}
	Time    = &timer{}
	TLS     = &tls{}
	Token   = &token{}
	Ufunc   = &ufunc{}
)

type convert struct{ convutil.Convert }
type dir struct{ dirutil.Dire }
type encode struct{ encodeutil.Encode }
type file struct{ fileutil.Fileutil }
type gob struct{ gobutil.Gob }
type host struct{ hostutil.Host }
type maps struct{ maputil.Maputil }
type php struct{ phputil.PHP }
type proto struct{ protoutil.Proto }
type random struct{ randomutil.Randomizer }
type rsa struct{ rsautil.Rsautil }
type runtime struct{ runtimeutil.Runtime }
type sliceu struct{ sliceutil.Sliceutil }
type str struct{ stringutil.Stringutil }
type timer struct{ timeutil.Timeutil }
type tls struct{ tlsutil.TLSutil }
type token struct{ tokenutil.Tokenutil }
type ufunc struct{ ufuncutil.Ufuncutil }
