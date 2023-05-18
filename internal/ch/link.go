//go:build chimera
// +build chimera

package ch

/*
#cgo linux LDFLAGS: -L${SRCDIR}/clibs/hs_centos72_install/lib64/ -lpcre
#cgo linux CFLAGS: -I${SRCDIR}/clibs/hs_centos72_install/include/hs
#cgo linux LDFLAGS: -L${SRCDIR}/clibs/hs_centos72_install/lib64/ -lhs
#cgo linux LDFLAGS: -L${SRCDIR}/clibs/hs_centos72_install/lib64/ -lchimera
#cgo darwin LDFLAGS: -L${SRCDIR}/clibs/hs_darwin_install/lib/ -lpcre
#cgo darwin CFLAGS: -I${SRCDIR}/clibs/hs_darwin_install/include/hs
#cgo darwin LDFLAGS: -L${SRCDIR}/clibs/hs_darwin_install/lib/ -lhs
#cgo darwin LDFLAGS: -L${SRCDIR}/clibs/hs_darwin_install/lib/ -lchimera
*/
import "C"
