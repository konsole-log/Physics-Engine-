package physics

import "slices"

// force generator used to add a force to one or more particles
type ForceGenerator interface {
	// update the force applied to the given particle
	UpdateForce(particle *Particle, duration float64)
}

// keeps track of one force generator and the particle it applies to
type forceRegistration struct {
	particle *Particle
	fg       ForceGenerator
}

// Holds all the force generators and the particles they apply to
type ForceRegistry struct {
	registrations []forceRegistration
}

// registers the given force generator to apply to the given particle
func (fr *ForceRegistry) Add(particle *Particle, fg ForceGenerator) {
	fr.registrations = append(fr.registrations, forceRegistration{particle: particle, fg: fg})
}

// removes the given registered pair from the registry
func (fr *ForceRegistry) Remove(particle *Particle, fg ForceGenerator) {
	for i, reg := range fr.registrations {
		if reg.particle == particle && reg.fg == fg {
			fr.registrations = slices.Delete(fr.registrations, i, i+1)
			return
		}
	}
}

// clears all registrations from the registry
func (fr *ForceRegistry) Clear() {
	fr.registrations = nil
}

// calls all force generators to update the forces of their corresponding particles
func (fr *ForceRegistry) UpdateForces(duration float64) {
	for _, reg := range fr.registrations {
		reg.fg.UpdateForce(reg.particle, duration)
	}
}

// Gravity force generator. One instance can be used for multiple particles
type Gravity struct {
	gravity Vector3
}

func NewGravity(g Vector3) *Gravity {
	return &Gravity{gravity: g}
}

// applies the gravitational force to the given particle
func (g *Gravity) UpdateForce(p *Particle, duration float64) {
	if p.inverseMass <= 0 {
		return
	}
	p.AddForce(g.gravity.Scale(p.GetMass()))
}

// Drag Force Generator. One instance can be used for multiple particles
type Drag struct {
	k1, k2 float64
}

func NewDrag(k1, k2 float64) *Drag {
	return &Drag{k1: k1, k2: k2}
}

func (d *Drag) UpdateForce(p *Particle, duration float64) {
	speed := p.velocity.Magnitude()
	if speed == 0 {
		return
	}
	coeff := d.k1*speed + d.k2*speed*speed
	force := p.velocity.Scale(-coeff / speed)
	p.AddForce(force)
}
