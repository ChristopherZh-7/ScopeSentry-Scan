// Code generated by 'yaegi extract github.com/Autumn-27/ScopeSentry-Scan/internal/contextmanager'. DO NOT EDIT.

package symbols

import (
	"github.com/Autumn-27/ScopeSentry-Scan/internal/contextmanager"
	"reflect"
)

func init() {
	Symbols["github.com/Autumn-27/ScopeSentry-Scan/internal/contextmanager/contextmanager"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"GlobalContextManagers": reflect.ValueOf(&contextmanager.GlobalContextManagers).Elem(),
		"NewContextManager":     reflect.ValueOf(contextmanager.NewContextManager),

		// type definitions
		"ContextManager": reflect.ValueOf((*contextmanager.ContextManager)(nil)),
	}
}
