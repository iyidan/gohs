//go:build chimera
// +build chimera

package ch

/*
#cgo linux LDFLAGS: -L../../clibs/hs_centos72_install/lib64/ -lpcre
#cgo linux CFLAGS: -I../../clibs/hs_centos72_install/include/hs
#cgo linux LDFLAGS: -L../../clibs/hs_centos72_install/lib64/ -lhs
#cgo linux LDFLAGS: -L../../clibs/hs_centos72_install/lib64/ -lchimera -lpcre -lhs -lstdc++ -lm -lc -lgcc_s -lgcc
#cgo linux LDFLAGS: -lm -lstdc++ -lpcre
#cgo darwin LDFLAGS: -L../../clibs/hs_darwin_install/lib/ -lpcre
#cgo darwin CFLAGS: -I../../clibs/hs_darwin_install/include/hs
#cgo darwin LDFLAGS: -L../../clibs/hs_darwin_install/lib/ -lhs
#cgo darwin LDFLAGS: -L../../clibs/hs_darwin_install/lib/ -lchimera
*/
import "C"
