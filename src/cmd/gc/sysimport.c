char*	sysimport =
	"package sys\n"
	"type sys._e002 {}\n"
	"type sys.any 24\n"
	"type sys._e003 *sys.any\n"
	"type sys._o207 {_e205 sys._e003}\n"
	"type sys.uint32 6\n"
	"type sys._i209 {_e206 sys.uint32}\n"
	"type sys._e001 (sys._e002 sys._o207 sys._i209)\n"
	"var !sys.mal sys._e001\n"
	"type sys._e005 {}\n"
	"type sys._e006 {}\n"
	"type sys._e007 {}\n"
	"type sys._e004 (sys._e005 sys._e006 sys._e007)\n"
	"var !sys.breakpoint sys._e004\n"
	"type sys._e009 {}\n"
	"type sys._e010 {}\n"
	"type sys.int32 5\n"
	"type sys._i215 {_e214 sys.int32}\n"
	"type sys._e008 (sys._e009 sys._e010 sys._i215)\n"
	"var !sys.panicl sys._e008\n"
	"type sys._e012 {}\n"
	"type sys._e013 {}\n"
	"type sys.bool 12\n"
	"type sys._i220 {_e219 sys.bool}\n"
	"type sys._e011 (sys._e012 sys._e013 sys._i220)\n"
	"var !sys.printbool sys._e011\n"
	"type sys._e015 {}\n"
	"type sys._e016 {}\n"
	"type sys.float64 10\n"
	"type sys._i225 {_e224 sys.float64}\n"
	"type sys._e014 (sys._e015 sys._e016 sys._i225)\n"
	"var !sys.printfloat sys._e014\n"
	"type sys._e018 {}\n"
	"type sys._e019 {}\n"
	"type sys.int64 7\n"
	"type sys._i230 {_e229 sys.int64}\n"
	"type sys._e017 (sys._e018 sys._e019 sys._i230)\n"
	"var !sys.printint sys._e017\n"
	"type sys._e021 {}\n"
	"type sys._e022 {}\n"
	"type sys._e023 25\n"
	"type sys.string *sys._e023\n"
	"type sys._i235 {_e234 sys.string}\n"
	"type sys._e020 (sys._e021 sys._e022 sys._i235)\n"
	"var !sys.printstring sys._e020\n"
	"type sys._e025 {}\n"
	"type sys._e026 {}\n"
	"type sys.uint8 2\n"
	"type sys._e027 *sys.uint8\n"
	"type sys._i240 {_e239 sys._e027}\n"
	"type sys._e024 (sys._e025 sys._e026 sys._i240)\n"
	"var !sys.printpointer sys._e024\n"
	"type sys._e029 {}\n"
	"type sys._o247 {_e244 sys.string}\n"
	"type sys._i249 {_e245 sys.string _e246 sys.string}\n"
	"type sys._e028 (sys._e029 sys._o247 sys._i249)\n"
	"var !sys.catstring sys._e028\n"
	"type sys._e031 {}\n"
	"type sys._o257 {_e254 sys.int32}\n"
	"type sys._i259 {_e255 sys.string _e256 sys.string}\n"
	"type sys._e030 (sys._e031 sys._o257 sys._i259)\n"
	"var !sys.cmpstring sys._e030\n"
	"type sys._e033 {}\n"
	"type sys._o268 {_e264 sys.string}\n"
	"type sys._i270 {_e265 sys.string _e266 sys.int32 _e267 sys.int32}\n"
	"type sys._e032 (sys._e033 sys._o268 sys._i270)\n"
	"var !sys.slicestring sys._e032\n"
	"type sys._e035 {}\n"
	"type sys._o279 {_e276 sys.uint8}\n"
	"type sys._i281 {_e277 sys.string _e278 sys.int32}\n"
	"type sys._e034 (sys._e035 sys._o279 sys._i281)\n"
	"var !sys.indexstring sys._e034\n"
	"type sys._e037 {}\n"
	"type sys._o288 {_e286 sys.string}\n"
	"type sys._i290 {_e287 sys.int64}\n"
	"type sys._e036 (sys._e037 sys._o288 sys._i290)\n"
	"var !sys.intstring sys._e036\n"
	"type sys._e039 {}\n"
	"type sys._o297 {_e294 sys.string}\n"
	"type sys._e040 *sys.uint8\n"
	"type sys._i299 {_e295 sys._e040 _e296 sys.int32}\n"
	"type sys._e038 (sys._e039 sys._o297 sys._i299)\n"
	"var !sys.byteastring sys._e038\n"
	"type sys._e042 {}\n"
	"type sys._e043 <>\n"
	"type sys._o308 {_e304 sys._e043}\n"
	"type sys._e044 *sys.uint8\n"
	"type sys._e045 *sys.uint8\n"
	"type sys._s315 {}\n"
	"type sys._e046 *sys._s315\n"
	"type sys._i310 {_e305 sys._e044 _e306 sys._e045 _e307 sys._e046}\n"
	"type sys._e041 (sys._e042 sys._o308 sys._i310)\n"
	"var !sys.mkiface sys._e041\n"
	"type sys._e048 {}\n"
	"type sys._o319 {_e318 sys.int32}\n"
	"type sys._e049 {}\n"
	"type sys._e047 (sys._e048 sys._o319 sys._e049)\n"
	"var !sys.argc sys._e047\n"
	"type sys._e051 {}\n"
	"type sys._o323 {_e322 sys.int32}\n"
	"type sys._e052 {}\n"
	"type sys._e050 (sys._e051 sys._o323 sys._e052)\n"
	"var !sys.envc sys._e050\n"
	"type sys._e054 {}\n"
	"type sys._o328 {_e326 sys.string}\n"
	"type sys._i330 {_e327 sys.int32}\n"
	"type sys._e053 (sys._e054 sys._o328 sys._i330)\n"
	"var !sys.argv sys._e053\n"
	"type sys._e056 {}\n"
	"type sys._o336 {_e334 sys.string}\n"
	"type sys._i338 {_e335 sys.int32}\n"
	"type sys._e055 (sys._e056 sys._o336 sys._i338)\n"
	"var !sys.envv sys._e055\n"
	"type sys._e058 {}\n"
	"type sys._o345 {_e342 sys.int32 _e343 sys.float64}\n"
	"type sys._i347 {_e344 sys.float64}\n"
	"type sys._e057 (sys._e058 sys._o345 sys._i347)\n"
	"var !sys.frexp sys._e057\n"
	"type sys._e060 {}\n"
	"type sys._o354 {_e351 sys.float64}\n"
	"type sys._i356 {_e352 sys.int32 _e353 sys.float64}\n"
	"type sys._e059 (sys._e060 sys._o354 sys._i356)\n"
	"var !sys.ldexp sys._e059\n"
	"type sys._e062 {}\n"
	"type sys._o364 {_e361 sys.float64 _e362 sys.float64}\n"
	"type sys._i366 {_e363 sys.float64}\n"
	"type sys._e061 (sys._e062 sys._o364 sys._i366)\n"
	"var !sys.modf sys._e061\n"
	"type sys._e064 {}\n"
	"type sys._e066 [sys.any] sys.any\n"
	"type sys._e065 *sys._e066\n"
	"type sys._o370 {hmap sys._e065}\n"
	"type sys._i372 {keysize sys.uint32 valsize sys.uint32 keyalg sys.uint32 valalg sys.uint32 hint sys.uint32}\n"
	"type sys._e063 (sys._e064 sys._o370 sys._i372)\n"
	"var !sys.newmap sys._e063\n"
	"type sys._e068 {}\n"
	"type sys._o380 {val sys.any}\n"
	"type sys._e070 [sys.any] sys.any\n"
	"type sys._e069 *sys._e070\n"
	"type sys._i382 {hmap sys._e069 key sys.any}\n"
	"type sys._e067 (sys._e068 sys._o380 sys._i382)\n"
	"var !sys.mapaccess1 sys._e067\n"
	"type sys._e072 {}\n"
	"type sys._o387 {val sys.any pres sys.bool}\n"
	"type sys._e074 [sys.any] sys.any\n"
	"type sys._e073 *sys._e074\n"
	"type sys._i389 {hmap sys._e073 key sys.any}\n"
	"type sys._e071 (sys._e072 sys._o387 sys._i389)\n"
	"var !sys.mapaccess2 sys._e071\n"
	"type sys._e076 {}\n"
	"type sys._e077 {}\n"
	"type sys._e079 [sys.any] sys.any\n"
	"type sys._e078 *sys._e079\n"
	"type sys._i394 {hmap sys._e078 key sys.any val sys.any}\n"
	"type sys._e075 (sys._e076 sys._e077 sys._i394)\n"
	"var !sys.mapassign1 sys._e075\n"
	"type sys._e081 {}\n"
	"type sys._e082 {}\n"
	"type sys._e084 [sys.any] sys.any\n"
	"type sys._e083 *sys._e084\n"
	"type sys._i400 {hmap sys._e083 key sys.any val sys.any pres sys.bool}\n"
	"type sys._e080 (sys._e081 sys._e082 sys._i400)\n"
	"var !sys.mapassign2 sys._e080\n"
	"type sys._e086 {}\n"
	"type sys._o410 {_e407 sys.string _e408 sys.bool}\n"
	"type sys._i412 {_e409 sys.string}\n"
	"type sys._e085 (sys._e086 sys._o410 sys._i412)\n"
	"var !sys.readfile sys._e085\n"
	"))\n"
;
