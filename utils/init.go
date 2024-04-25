package utils

import (
	"github.com/dunpju/higo-utils/utils/convutil"
	"github.com/dunpju/higo-utils/utils/dirutil"
	"github.com/dunpju/higo-utils/utils/encodeutil"
	"github.com/dunpju/higo-utils/utils/fileutil"
	"github.com/dunpju/higo-utils/utils/gobutil"
	"github.com/dunpju/higo-utils/utils/hostutil"
	"github.com/dunpju/higo-utils/utils/maputil"
	"github.com/dunpju/higo-utils/utils/modutil"
	"github.com/dunpju/higo-utils/utils/phputil"
	"github.com/dunpju/higo-utils/utils/protoutil"
	"github.com/dunpju/higo-utils/utils/randomutil"
	"github.com/dunpju/higo-utils/utils/rsautil"
	"github.com/dunpju/higo-utils/utils/runtimeutil"
	"github.com/dunpju/higo-utils/utils/sliceutil"
	"github.com/dunpju/higo-utils/utils/stringutil"
	"github.com/dunpju/higo-utils/utils/structutil"
	"github.com/dunpju/higo-utils/utils/timeutil"
	"github.com/dunpju/higo-utils/utils/tlsutil"
	"github.com/dunpju/higo-utils/utils/tokenutil"
	"github.com/dunpju/higo-utils/utils/ufuncutil"
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
	Mod     = &mod{}
	Struct  = &_struct{}
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
type mod struct{ modutil.Mod }
type _struct struct{ structutil.Struct }
