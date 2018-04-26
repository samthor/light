package light

// Bit describes a toggleable light.
type Bit struct {
	ref      *Ref
	priority int
	color    HasColor
}

// NewBit creates a new Bit light.
func NewBit(priority int, color HasColor) *Bit {
	return &Bit{
		priority: priority,
		color:    color,
	}
}

// Update updates the config of this Bit light.
func (b *Bit) Update(color HasColor) {
	b.color = color
	if b.ref != nil {
		b.enable() // reenable if it was enabled
	}
}

// Set enables or disables the Bit light.
func (b *Bit) Set(on bool) {
	alreadyOn := b.ref != nil
	if on == alreadyOn {
		// nothing
	} else if on {
		b.enable()
	} else {
		b.ref.Cancel()
		b.ref = nil
	}
}

func (b *Bit) enable() {
	b.ref.Cancel() // safe anyway
	ref := Add(Task{
		Priority: b.priority,
		Color:    b.color,
	})
	b.ref = &ref
}
