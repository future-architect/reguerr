package gen

import "fmt"

// Check unique message code
func Validate(decls []*Decl) error {
	unique := make(map[string]struct{}, len(decls))
	for _, decl := range decls {
		if _, ok := unique[decl.Code]; ok {
			return fmt.Errorf("duplicated message code: %v", decl.Code)
		}
		unique[decl.Code] = struct{}{}

		// call Build() check. reguerr DSL must call Build() function.
		if !decl.CallBuild {
			return fmt.Errorf("reguerr DSL requires Build() function call at the end: ^%s", decl.Name)
		}
	}

	return nil
}
