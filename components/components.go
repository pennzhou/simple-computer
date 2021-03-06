package components

import (
	"github.com/djhworld/simple-computer/circuit"
	"github.com/djhworld/simple-computer/utils"
)

const BUS_WIDTH = 16

type Component interface {
	ConnectOutput(Component)
	SetInputWire(int, bool)
	GetOutputWire(int) bool
}

type Enabler struct {
	inputs  [BUS_WIDTH]circuit.Wire
	gates   [BUS_WIDTH]circuit.ANDGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewEnabler() *Enabler {
	e := new(Enabler)

	for i, _ := range e.gates {
		e.gates[i] = *circuit.NewANDGate()
	}
	return e
}

func (e *Enabler) ConnectOutput(b Component) {
	e.next = b
}

func (e *Enabler) GetOutputWire(index int) bool {
	return e.outputs[index].Get()
}

func (e *Enabler) SetInputWire(index int, value bool) {
	e.inputs[index].Update(value)
}

func (e *Enabler) Update(enable bool) {
	for i := 0; i < len(e.gates); i++ {
		e.gates[i].Update(e.inputs[i].Get(), enable)
		e.outputs[i].Update(e.gates[i].Output())
	}

	if e.next != nil {
		for i := 0; i < len(e.outputs); i++ {
			e.next.SetInputWire(i, e.outputs[i].Get())
		}
	}
}

// TODO not sure if this is exactly how this should look...
type LeftShifter struct {
	inputs   [BUS_WIDTH]circuit.Wire
	outputs  [BUS_WIDTH]circuit.Wire
	shiftIn  circuit.Wire
	shiftOut circuit.Wire
	next     Component
}

func NewLeftShifter() *LeftShifter {
	return new(LeftShifter)
}

func (l *LeftShifter) ConnectOutput(b Component) {
	l.next = b
}

func (l *LeftShifter) GetOutputWire(index int) bool {
	return l.outputs[index].Get()
}

func (l *LeftShifter) SetInputWire(index int, value bool) {
	l.inputs[index].Update(value)
}

func (l *LeftShifter) ShiftOut() bool {
	return l.shiftOut.Get()
}

func (l *LeftShifter) Update(shiftIn bool) {
	l.shiftIn.Update(shiftIn)
	l.shiftOut.Update(l.inputs[0].Get())
	l.outputs[0].Update(l.inputs[1].Get())
	l.outputs[1].Update(l.inputs[2].Get())
	l.outputs[2].Update(l.inputs[3].Get())
	l.outputs[3].Update(l.inputs[4].Get())
	l.outputs[4].Update(l.inputs[5].Get())
	l.outputs[5].Update(l.inputs[6].Get())
	l.outputs[6].Update(l.inputs[7].Get())
	l.outputs[7].Update(l.inputs[8].Get())
	l.outputs[8].Update(l.inputs[9].Get())
	l.outputs[9].Update(l.inputs[10].Get())
	l.outputs[10].Update(l.inputs[11].Get())
	l.outputs[11].Update(l.inputs[12].Get())
	l.outputs[12].Update(l.inputs[13].Get())
	l.outputs[13].Update(l.inputs[14].Get())
	l.outputs[14].Update(l.inputs[15].Get())
	l.outputs[15].Update(l.shiftIn.Get())
}

type RightShifter struct {
	inputs   [BUS_WIDTH]circuit.Wire
	shiftIn  circuit.Wire
	shiftOut circuit.Wire
	outputs  [BUS_WIDTH]circuit.Wire
	next     Component
}

func NewRightShifter() *RightShifter {
	return new(RightShifter)
}

func (r *RightShifter) ConnectOutput(b Component) {
	r.next = b
}

func (r *RightShifter) GetOutputWire(index int) bool {
	return r.outputs[index].Get()
}

func (r *RightShifter) SetInputWire(index int, value bool) {
	r.inputs[index].Update(value)
}

func (r *RightShifter) ShiftOut() bool {
	return r.shiftOut.Get()
}

func (r *RightShifter) Update(shiftIn bool) {
	r.shiftIn.Update(shiftIn)
	r.outputs[0].Update(r.shiftIn.Get())
	r.outputs[1].Update(r.inputs[0].Get())
	r.outputs[2].Update(r.inputs[1].Get())
	r.outputs[3].Update(r.inputs[2].Get())
	r.outputs[4].Update(r.inputs[3].Get())
	r.outputs[5].Update(r.inputs[4].Get())
	r.outputs[6].Update(r.inputs[5].Get())
	r.outputs[7].Update(r.inputs[6].Get())
	r.outputs[8].Update(r.inputs[7].Get())
	r.outputs[9].Update(r.inputs[8].Get())
	r.outputs[10].Update(r.inputs[9].Get())
	r.outputs[11].Update(r.inputs[10].Get())
	r.outputs[12].Update(r.inputs[11].Get())
	r.outputs[13].Update(r.inputs[12].Get())
	r.outputs[14].Update(r.inputs[13].Get())
	r.outputs[15].Update(r.inputs[14].Get())
	r.shiftOut.Update(r.inputs[15].Get())
}

type IsZero struct {
	inputs  [BUS_WIDTH]circuit.Wire
	orer    ORer
	notGate circuit.NOTGate
	output  circuit.Wire
}

func NewIsZero() *IsZero {
	z := new(IsZero)
	z.orer = *NewORer()
	z.notGate = *circuit.NewNOTGate()

	return z
}

func (z *IsZero) Reset() {
	z.output.Update(false)
}

func (z *IsZero) ConnectOutput(b Component) {
	// noop
}

func (z *IsZero) GetOutputWire(index int) bool {
	// only 1 wire
	return z.output.Get()
}

func (z *IsZero) SetInputWire(index int, value bool) {
	z.inputs[index].Update(value)
}

func (z *IsZero) Update() {
	for i, _ := range z.inputs {
		z.orer.SetInputWire(i, z.inputs[i].Get())
		z.orer.SetInputWire(i+BUS_WIDTH, z.inputs[i].Get())
	}
	z.orer.Update()

	for i, _ := range z.orer.outputs {
		if z.orer.outputs[i].Get() {
			z.notGate.Update(true)
			z.output.Update(z.notGate.Output())
			return
		} else {
			z.notGate.Update(false)
		}
	}

	z.output.Update(z.notGate.Output())

}

type NOTer struct {
	inputs  [BUS_WIDTH]circuit.Wire
	gates   [BUS_WIDTH]circuit.NOTGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewNOTer() *NOTer {
	n := new(NOTer)

	for i, _ := range n.gates {
		n.gates[i] = *circuit.NewNOTGate()
	}

	return n
}

func (n *NOTer) ConnectOutput(b Component) {
	n.next = b
}

func (n *NOTer) GetOutputWire(index int) bool {
	return n.outputs[index].Get()
}

func (n *NOTer) SetInputWire(index int, value bool) {
	n.inputs[index].Update(value)
}

func (n *NOTer) Update() {
	for i, _ := range n.gates {
		n.gates[i].Update(n.inputs[i].Get())
		n.outputs[i].Update(n.gates[i].Output())
	}
}

type ANDer struct {
	inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]circuit.ANDGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewANDer() *ANDer {
	a := new(ANDer)

	for i, _ := range a.gates {
		a.gates[i] = *circuit.NewANDGate()
	}

	return a
}

func (a *ANDer) ConnectOutput(b Component) {
	a.next = b
}

func (a *ANDer) GetOutputWire(index int) bool {
	return a.outputs[index].Get()
}

func (a *ANDer) SetInputWire(index int, value bool) {
	a.inputs[index].Update(value)
}

func (a *ANDer) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i, _ := range a.gates {
		a.gates[i].Update(a.inputs[awire].Get(), a.inputs[bwire].Get())
		a.outputs[i].Update(a.gates[i].Output())
		awire++
		bwire++
	}
}

type ORer struct {
	inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]circuit.ORGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewORer() *ORer {
	o := new(ORer)

	for i, _ := range o.gates {
		o.gates[i] = *circuit.NewORGate()
	}

	return o
}

func (o *ORer) ConnectOutput(b Component) {
	o.next = b
}

func (o *ORer) GetOutputWire(index int) bool {
	return o.outputs[index].Get()
}

func (o *ORer) SetInputWire(index int, value bool) {
	o.inputs[index].Update(value)
}

func (o *ORer) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i, _ := range o.gates {
		o.gates[i].Update(o.inputs[awire].Get(), o.inputs[bwire].Get())
		o.outputs[i].Update(o.gates[i].Output())
		awire++
		bwire++
	}
}

type XORer struct {
	inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]circuit.XORGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewXORer() *XORer {
	o := new(XORer)

	for i, _ := range o.gates {
		o.gates[i] = *circuit.NewXORGate()
	}

	return o
}

func (o *XORer) ConnectOutput(b Component) {
	o.next = b
}

func (o *XORer) GetOutputWire(index int) bool {
	return o.outputs[index].Get()
}

func (o *XORer) SetInputWire(index int, value bool) {
	o.inputs[index].Update(value)
}

func (o *XORer) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i, _ := range o.gates {
		o.gates[i].Update(o.inputs[awire].Get(), o.inputs[bwire].Get())
		o.outputs[i].Update(o.gates[i].Output())
		awire++
		bwire++
	}
}

type Compare2 struct {
	inputA      circuit.Wire
	inputB      circuit.Wire
	xor1        circuit.XORGate
	not1        circuit.NOTGate
	and1        circuit.ANDGate
	andgate3    ANDGate3
	or1         circuit.ORGate
	out         circuit.Wire
	equalIn     circuit.Wire
	equalOut    circuit.Wire
	isLargerIn  circuit.Wire
	isLargerOut circuit.Wire
}

func NewCompare2() *Compare2 {
	c := new(Compare2)

	c.inputA = *circuit.NewWire("a", false)
	c.inputB = *circuit.NewWire("b", false)
	c.equalIn = *circuit.NewWire("eqin", false)
	c.equalOut = *circuit.NewWire("eqin", false)
	c.isLargerIn = *circuit.NewWire("largerin", false)
	c.isLargerOut = *circuit.NewWire("largerin", false)
	c.out = *circuit.NewWire("out", false)

	c.xor1 = *circuit.NewXORGate()
	c.not1 = *circuit.NewNOTGate()
	c.and1 = *circuit.NewANDGate()
	c.andgate3 = *NewANDGate3()
	c.or1 = *circuit.NewORGate()

	return c
}

func (g *Compare2) Equal() bool {
	return g.equalOut.Get()
}

func (g *Compare2) Larger() bool {
	return g.isLargerOut.Get()
}

func (g *Compare2) Output() bool {
	return g.out.Get()
}

func (g *Compare2) Update(inputA, inputB, equalIn, isLargerIn bool) {
	g.inputA.Update(inputA)
	g.inputB.Update(inputB)
	g.equalIn.Update(equalIn)
	g.isLargerIn.Update(isLargerIn)

	g.xor1.Update(g.inputA.Get(), g.inputB.Get())
	g.not1.Update(g.xor1.Output())
	g.and1.Update(g.not1.Output(), g.equalIn.Get())
	g.equalOut.Update(g.and1.Output())

	g.andgate3.Update(g.equalIn.Get(), g.inputA.Get(), g.xor1.Output())
	g.or1.Update(g.andgate3.Output(), g.isLargerIn.Get())
	g.isLargerOut.Update(g.or1.Output())

	g.out.Update(g.xor1.Output())
}

type Comparator struct {
	inputs       [BUS_WIDTH * 2]circuit.Wire
	equalIn      circuit.Wire
	aIsLargerIn  circuit.Wire
	compares     [BUS_WIDTH]Compare2
	outputs      [BUS_WIDTH]circuit.Wire
	equalOut     circuit.Wire
	aIsLargerOut circuit.Wire
	next         Component
}

func NewComparator() *Comparator {
	c := new(Comparator)

	for i, _ := range c.compares {
		c.compares[i] = *NewCompare2()
	}

	return c
}

func (c *Comparator) ConnectOutput(b Component) {
	c.next = b
}

func (c *Comparator) GetOutputWire(index int) bool {
	return c.outputs[index].Get()
}

func (c *Comparator) SetInputWire(index int, value bool) {
	c.inputs[index].Update(value)
}

func (c *Comparator) Update() {
	// these start out as 1 and 0 respectively
	c.equalIn.Update(true)
	c.aIsLargerIn.Update(false)

	// top 16 bits are <b>, bottom 16 bits are <a>
	awire := 0
	bwire := BUS_WIDTH

	for i := range c.compares {
		c.compares[i].Update(c.inputs[awire].Get(), c.inputs[bwire].Get(), c.equalIn.Get(), c.aIsLargerIn.Get())
		c.outputs[i].Update(c.compares[i].Output())
		c.equalOut.Update(c.compares[i].Equal())
		c.aIsLargerOut.Update(c.compares[i].Larger())

		c.equalIn.Update(c.compares[i].Equal())
		c.aIsLargerIn.Update(c.compares[i].Larger())
		awire++
		bwire++
	}
}

func (g *Comparator) Equal() bool {
	return g.equalOut.Get()
}

func (g *Comparator) Larger() bool {
	return g.aIsLargerOut.Get()
}

type BusOne struct {
	inputBus  *Bus
	outputBus *Bus
	inputs    [BUS_WIDTH]circuit.Wire
	bus1      circuit.Wire
	andGates  [BUS_WIDTH - 1]circuit.ANDGate
	notGate   circuit.NOTGate
	orGate    circuit.ORGate
	outputs   [BUS_WIDTH]circuit.Wire
	next      Component
}

func NewBusOne(inputBus, outputBus *Bus) *BusOne {
	b := new(BusOne)
	b.inputBus = inputBus
	b.outputBus = outputBus

	for i, _ := range b.andGates {
		b.andGates[i] = *circuit.NewANDGate()
	}

	b.notGate = *circuit.NewNOTGate()
	b.orGate = *circuit.NewORGate()

	return b
}

func (b *BusOne) ConnectOutput(bc Component) {
	b.next = bc
}

func (b *BusOne) GetOutputWire(index int) bool {
	return b.outputs[index].Get()
}

func (b *BusOne) SetInputWire(index int, value bool) {
	b.inputs[index].Update(value)
}

func (b *BusOne) Enable() {
	b.bus1.Update(true)
}

func (b *BusOne) Disable() {
	b.bus1.Update(false)
}

func (b *BusOne) Update() {
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		b.inputs[i].Update(b.inputBus.GetOutputWire(i))
	}

	b.notGate.Update(b.bus1.Get())

	b.andGates[0].Update(b.inputs[0].Get(), b.notGate.Output())
	b.andGates[1].Update(b.inputs[1].Get(), b.notGate.Output())
	b.andGates[2].Update(b.inputs[2].Get(), b.notGate.Output())
	b.andGates[3].Update(b.inputs[3].Get(), b.notGate.Output())
	b.andGates[4].Update(b.inputs[4].Get(), b.notGate.Output())
	b.andGates[5].Update(b.inputs[5].Get(), b.notGate.Output())
	b.andGates[6].Update(b.inputs[6].Get(), b.notGate.Output())
	b.andGates[7].Update(b.inputs[7].Get(), b.notGate.Output())
	b.andGates[8].Update(b.inputs[8].Get(), b.notGate.Output())
	b.andGates[9].Update(b.inputs[9].Get(), b.notGate.Output())
	b.andGates[10].Update(b.inputs[10].Get(), b.notGate.Output())
	b.andGates[11].Update(b.inputs[11].Get(), b.notGate.Output())
	b.andGates[12].Update(b.inputs[12].Get(), b.notGate.Output())
	b.andGates[13].Update(b.inputs[13].Get(), b.notGate.Output())
	b.andGates[14].Update(b.inputs[14].Get(), b.notGate.Output())
	b.orGate.Update(b.inputs[15].Get(), b.bus1.Get())

	b.outputs[0].Update(b.andGates[0].Output())
	b.outputs[1].Update(b.andGates[1].Output())
	b.outputs[2].Update(b.andGates[2].Output())
	b.outputs[3].Update(b.andGates[3].Output())
	b.outputs[4].Update(b.andGates[4].Output())
	b.outputs[5].Update(b.andGates[5].Output())
	b.outputs[6].Update(b.andGates[6].Output())
	b.outputs[7].Update(b.andGates[7].Output())
	b.outputs[8].Update(b.andGates[8].Output())
	b.outputs[9].Update(b.andGates[9].Output())
	b.outputs[10].Update(b.andGates[10].Output())
	b.outputs[11].Update(b.andGates[11].Output())
	b.outputs[12].Update(b.andGates[12].Output())
	b.outputs[13].Update(b.andGates[13].Output())
	b.outputs[14].Update(b.andGates[14].Output())
	b.outputs[15].Update(b.orGate.Output())

	for i := BUS_WIDTH - 1; i >= 0; i-- {
		b.outputBus.SetInputWire(i, b.outputs[i].Get())
	}
}

func (b *BusOne) String() string {
	var output uint16
	var x int = 0
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		if b.outputs[i].Get() {
			output = output | (1 << uint16(x))
		} else {
			output = output & ^(1 << uint16(x))
		}
		x++
	}
	return utils.ValueToString(output)
}
