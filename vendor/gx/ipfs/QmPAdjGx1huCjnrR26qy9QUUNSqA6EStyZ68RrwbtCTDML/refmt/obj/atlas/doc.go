/*
	Atlas types are used to define how to map Go values into refmt token streams.

	Atlas information may be autogenerated based on struct tags automatically,
	but you can also specify custom AtlasEntry info to use advanced features
	and define custom transformations.

	An Atlas is a collection of AtlasEntry (plus some internal indexing).
	Typical usage is to declare an AtlasEntry for your structs (often near by the
	struct definition), then

	Building an AtlasEntry for some type called `Formula` looks like this:

		atlas.BuildEntry(Formula{}).StructMap().Autogenerate().Complete()

	Building an AtlasEntry always starts with `atlas.BuildEntry(x)` where `x` is
	a dummy object used to convey type information.
	The next function in the chain declares what kind of behavior we're going
	to use to turn that type of object into its serial form.
	(In the above example, we're declaring that we want refmt to see the `Formula`
	type as a struct and traverse its fields.  There are many other options!)
	Subsequent functions are specific to what kind of walking and mapping we've
	chosen.  For struct walking, this may involve declaring fields and custom serial
	names to map them to; for a "Transform" we'd instead have to provide callbacks
	to do the transformation from the `Formula` type to some other type; etcetera.
	The final function in the chain is always called `Complete`, and returns
	a ready-to-use AtlasEntry.

	Building a complete Atlas for a whole suite of serializable types is as
	easy as putting a bunch of them together:

		atlas.Build(
			atlas.BuildEntry(Foo{}).StructMap().Autogenerate().Complete(),
			atlas.BuildEntry(Bar{}).StructMap().Autogenerate().Complete(),
			atlas.BuildEntry(Baz{}).StructMap().Autogenerate().Complete(),
		)

	You can put your entire protocol into one Atlas.
	It's also possible to build several different Atlases each with different
	sets of AtlasEntry.  This may be useful if you have a protocol where some
	messages are not valid during some phases of communication, and you would
	like to use the Atlas as a form of whitelisting for what can be
	marshalled/unmarshalled.
*/
package atlas