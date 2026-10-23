package physics

import (
	"errors"
	"math"
)

// here inverseMass means 1/m because acceleration = 1/m * f
type Particle struct {
	position, velocity, acceleration, forceAccum Vector3
	damping, inverseMass                         float64
}

func (p *Particle) Integrate(duration float64) error {
	if duration <= 0.0 {
		return errors.New("duration must be greater than 0")
	}
	//update position using formula p=p+p*dt
	p.position = p.position.AddScaledVector(p.velocity, duration)

	//workout the acceleration from the force.
	//we use formula as a = f*1/m or f*inverseMass
	resultingAcc := p.acceleration
	resultingAcc = resultingAcc.AddScaledVector(p.forceAccum, p.inverseMass)

	//update linear velocity from the acceleartion as v = v+a*dt
	p.velocity = p.velocity.AddScaledVector(resultingAcc, duration)

	//Impose drag as v = v*(damping^dt)
	p.velocity = p.velocity.Scale(math.Pow(p.damping, duration))
	p.ClearAccumulator()
	return nil
}

func (p *Particle) ClearAccumulator() {
	p.forceAccum = Vector3{}
}

func (p *Particle) AddForce(force Vector3) {
	p.forceAccum = p.forceAccum.Add(force)
}

func (p Particle) GetMass() float64 {
	return (1 / p.inverseMass)
}
