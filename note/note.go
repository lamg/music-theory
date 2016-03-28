// A Note is used to represent the relative duration and pitch of a sound.
package note

// Note models a musical note
type Note struct {
	Class     Class   // Class of pitch
	Octave    Octave  // Octave #
}

// Named note returns a Note model
func Named(text string) (n *Note) {
	n = &Note{}

	// First the name, including octave shift.
	n.Class, n.Octave = NameOf(text)

	// Last, add the originally named octave.
	n.Octave += OctaveOf(text)

	return
}

// OfClass pitch returns a Note model
func OfClass(class Class) (n *Note) {
	n = &Note{}
	n.Class = class
	return
}

// ClassNamed returns a pitch Class
func ClassNamed(text string) Class {
	n := Named(text)
	return n.Class
}
