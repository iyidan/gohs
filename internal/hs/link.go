package hs

/*
#cgo linux CFLAGS: -I${SRCDIR}/clibs/hs_centos72_install/include/hs
#cgo linux LDFLAGS: -L${SRCDIR}/clibs/hs_centos72_install/lib64/ -lhs
#cgo darwin CFLAGS: -I${SRCDIR}/clibs/hs_darwin_install/include/hs
#cgo darwin LDFLAGS: -L${SRCDIR}/clibs/hs_darwin_install/lib/ -lhs
#cgo linux LDFLAGS: -lm -lstdc++
*/
import "C"
