package hvue

import "github.com/gopherjs/gopherjs/js"

// Define name as a computed property.  Note that name *must not* be set in
// data for this to work.  It's probably best if it's not even a slot in the
// struct.  Only access it via vm.Get/Set.  You could create an accessor; see
// the computed-with-setter example.
func Computed(name string, f func(vm *VM) interface{}) option {
	return func(c *Config) {
		if c.Computed == js.Undefined {
			c.Computed = NewObject()
		}
		c.Computed.Set(name, jsCallWithVM(f))
	}
}

// Define name as a computed property with explicit get & set.  Note that name
// *must not* be set in data for this to work.  It's probably best if it's not
// even a slot in the struct.  Only access it via vm.Get/Set.  You could
// create an accessor; see the computed-with-setter example.
func ComputedWithGetSet(name string, get func(vm *VM) interface{}, set func(vm *VM, newValue *js.Object)) option {
	return func(c *Config) {
		if c.Computed == js.Undefined {
			c.Computed = NewObject()
		}
		c.Computed.Set(name,
			js.M{
				"get": jsCallWithVM(get),
				"set": js.MakeFunc(
					func(this *js.Object, args []*js.Object) interface{} {
						vm := &VM{Object: this}
						set(vm, args[0])
						return nil
					})})
	}
}
