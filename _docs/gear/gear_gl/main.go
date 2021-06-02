package main

import (
	. "fmt"
	"gl"
	. "math"
	"sdl"
	"time"
)

var (
	T0     int64 = 0
	Frames int64 = 0
)

func Vertex3f(a, b, c float64) {
	gl.Vertex3f(gl.GLfloat(a), gl.GLfloat(b), gl.GLfloat(c))
}

func Normal3f(a, b, c float64) {
	gl.Normal3f(gl.GLfloat(a), gl.GLfloat(b), gl.GLfloat(c))
}

func gear(inner_radius, outer_radius, width, teeth, tooth_depth float64) {
	var i float64
	var r0, r1, r2, angle, da, u, v, length float64

	r0 = inner_radius
	r1 = outer_radius - tooth_depth/2.0
	r2 = outer_radius + tooth_depth/2.0

	da = 2.0 * Pi / teeth / 4.0

	gl.ShadeModel(gl.FLAT)

	gl.Normal3f(0.0, 0.0, 1.0)

	gl.Begin(gl.QUAD_STRIP)

	for i = 0; i <= teeth; i++ {
		angle = i * 2.0 * Pi / teeth
		Vertex3f(r0*Cos(angle), r0*Sin(angle), width*0.5)
		Vertex3f(r1*Cos(angle), r1*Sin(angle), width*0.5)
		if i < teeth {
			Vertex3f(r0*Cos(angle), r0*Sin(angle), width*0.5)
			Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), width*0.5)
		}
	}

	gl.End()

	gl.Begin(gl.QUADS)
	da = 2.0 * Pi / teeth / 4.0

	for i = 0; i < teeth; i++ {
		angle = i * 2.0 * Pi / teeth

		Vertex3f(r1*Cos(angle), r1*Sin(angle), width*0.5)
		Vertex3f(r2*Cos(angle+da), r2*Sin(angle+da), width*0.5)
		Vertex3f(r2*Cos(angle+2*da), r2*Sin(angle+2*da), width*0.5)
		Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), width*0.5)
	}
	gl.End()

	gl.Normal3f(0.0, 0.0, -1.0)

	/* draw back face */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = i * 2.0 * Pi / teeth
		Vertex3f(r1*Cos(angle), r1*Sin(angle), -width*0.5)
		Vertex3f(r0*Cos(angle), r0*Sin(angle), -width*0.5)

		if i < teeth {
			Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), -width*0.5)
			Vertex3f(r0*Cos(angle), r0*Sin(angle), -width*0.5)
		}
	}
	gl.End()

	/* draw back sides of teeth */
	gl.Begin(gl.QUADS)
	da = 2.0 * Pi / teeth / 4.0
	for i = 0; i < teeth; i++ {
		angle = i * 2.0 * Pi / teeth

		Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), -width*0.5)
		Vertex3f(r2*Cos(angle+2*da), r2*Sin(angle+2*da), -width*0.5)
		Vertex3f(r2*Cos(angle+da), r2*Sin(angle+da), -width*0.5)
		Vertex3f(r1*Cos(angle), r1*Sin(angle), -width*0.5)
	}
	gl.End()

	/* draw outward faces of teeth */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i < teeth; i++ {
		angle = i * 2.0 * Pi / teeth

		Vertex3f(r1*Cos(angle), r1*Sin(angle), width*0.5)
		Vertex3f(r1*Cos(angle), r1*Sin(angle), -width*0.5)
		u = r2*Cos(angle+da) - r1*Cos(angle)
		v = r2*Sin(angle+da) - r1*Sin(angle)
		length = Sqrt(u*u + v*v)
		u /= length
		v /= length
		Normal3f(v, -u, 0.0)
		Vertex3f(r2*Cos(angle+da), r2*Sin(angle+da), width*0.5)
		Vertex3f(r2*Cos(angle+da), r2*Sin(angle+da), -width*0.5)
		Normal3f(Cos(angle), Sin(angle), 0.0)
		Vertex3f(r2*Cos(angle+2*da), r2*Sin(angle+2*da), width*0.5)
		Vertex3f(r2*Cos(angle+2*da), r2*Sin(angle+2*da), -width*0.5)
		u = r1*Cos(angle+3*da) - r2*Cos(angle+2*da)
		v = r1*Sin(angle+3*da) - r2*Sin(angle+2*da)
		Normal3f(v, -u, 0.0)
		Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), width*0.5)
		Vertex3f(r1*Cos(angle+3*da), r1*Sin(angle+3*da), -width*0.5)
		Normal3f(Cos(angle), Sin(angle), 0.0)
	}

	Vertex3f(r1*Cos(0), r1*Sin(0), width*0.5)
	Vertex3f(r1*Cos(0), r1*Sin(0), -width*0.5)

	gl.End()

	gl.ShadeModel(gl.SMOOTH)

	/* draw inside radius cylinder */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = i * 2.0 * Pi / teeth
		Normal3f(-Cos(angle), -Sin(angle), 0.0)
		Vertex3f(r0*Cos(angle), r0*Sin(angle), -width*0.5)
		Vertex3f(r0*Cos(angle), r0*Sin(angle), width*0.5)
	}
	gl.End()
}

var (
	view_rotx, view_roty, view_rotz gl.GLfloat = 20.0, 30.0, 0.0
	gear1, gear2, gear3             gl.GLuint
	angle                           gl.GLfloat = 0.0
)

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Rotatef(view_rotx, 1.0, 0.0, 0.0)
	gl.Rotatef(view_roty, 0.0, 1.0, 0.0)
	gl.Rotatef(view_rotz, 0.0, 0.0, 1.0)

	gl.PushMatrix()
	gl.Translatef(-3.0, -2.0, 0.0)
	gl.Rotatef(angle, 0.0, 0.0, 1.0)
	gl.CallList(gear1)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translatef(3.1, -2.0, 0.0)
	gl.Rotatef(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	gl.CallList(gear2)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translatef(-3.1, 4.2, 0.0)
	gl.Rotatef(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	gl.CallList(gear3)
	gl.PopMatrix()

	gl.PopMatrix()

	sdl.GL_SwapBuffers()

	Frames++

	t := time.Seconds()

	if t-T0 >= 5 {
		seconds := t - T0
		fps := Frames / seconds
		Println(Frames, "frames in", seconds, "seconds =", fps, "FPS")
		T0 = t
		Frames = 0
	}
}

func idle() {
	angle += 2.0
	sdl.GL_SwapBuffers()
}

func setup() {
	pos := [4]gl.GLfloat{5.0, 5.0, 10.0, 0.0}
	red := [4]gl.GLfloat{0.8, 0.1, 0.0, 1.0}
	green := [4]gl.GLfloat{0.0, 0.8, 0.2, 1.0}
	blue := [4]gl.GLfloat{0.2, 0.2, 1.0, 1.0}

	gl.Lightfv(gl.LIGHT0, gl.POSITION, &pos[0])
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	/* make the gears */
	gear1 = gl.GenLists(1)
	gl.NewList(gear1, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, &red[0])
	gear(1.0, 4.0, 1.0, 20, 0.7)
	gl.EndList()

	gear2 = gl.GenLists(1)
	gl.NewList(gear2, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, &green[0])
	gear(0.5, 2.0, 2.0, 10, 0.7)
	gl.EndList()

	gear3 = gl.GenLists(1)
	gl.NewList(gear3, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, &blue[0])
	gear(1.3, 2.0, 0.5, 10, 0.7)
	gl.EndList()

	gl.Enable(gl.NORMALIZE)
}

func reshape(width, height gl.GLfloat) {
	h := gl.GLdouble(height / width)

	gl.Viewport(0, 0, gl.GLsizei(width), gl.GLsizei(height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-1.0, 1.0, -h, h, 5.0, 60.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -40.0)
}

func main() {
	sdl.Init(sdl.INIT_VIDEO)

	var screen = sdl.SetVideoMode(300, 300, 32, sdl.OPENGL)

	if screen == nil {
		panic(sdl.GetError())
	}

	if gl.Init() != 0 {
		panic(gl.GetError())
	}

	reshape(300, 300)
	setup()
	fps := NewFramerate()
	fps.SetFramerate(300)

	var running = true

	for running {
		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
			default:
				{
					keys := sdl.GetKeyState()
					shift := keys[sdl.K_LSHIFT] == 1 || keys[sdl.K_RSHIFT] == 1

					for n, i := range keys {
						if i == 1 {
							switch n {
							case sdl.K_z:
								if shift {
									view_rotz += 5.0
								} else {
									view_rotz -= 5.0
								}
							case sdl.K_ESCAPE:
								running = false
							case sdl.K_UP:
								view_rotx += 5.0
							case sdl.K_DOWN:
								view_rotx -= 5.0
							case sdl.K_LEFT:
								view_roty += 5.0
							case sdl.K_RIGHT:
								view_roty -= 5.0
							}
						}
					}
				}
			}
		}

		draw()
		idle()
		fps.FramerateDelay()
	}

	sdl.Quit()
}

/*
A pure Go version of SDL_framerate
*/

const (
	FPS_UPPER_LIMIT = 1000
	FPS_LOWER_LIMIT = 1
	FPS_DEFAULT     = 30
)

type FPSmanager struct {
	framecount uint32
	rateticks  float
	lastticks  uint32
	rate       uint32
}

func NewFramerate() *FPSmanager {
	return &FPSmanager{
		framecount: 0,
		rate:       FPS_DEFAULT,
		rateticks:  (1000.0 / float(FPS_DEFAULT)),
		lastticks:  sdl.GetTicks(),
	}
}

func (manager *FPSmanager) SetFramerate(rate uint32) {
	if rate >= FPS_LOWER_LIMIT && rate <= FPS_UPPER_LIMIT {
		manager.framecount = 0
		manager.rate = rate
		manager.rateticks = 1000.0 / float(rate)
	} else {
	}
}

func (manager *FPSmanager) GetFramerate() uint32 {
	return manager.rate
}

func (manager *FPSmanager) FramerateDelay() {
	var current_ticks, target_ticks, the_delay uint32

	// next frame
	manager.framecount++

	// get/calc ticks
	current_ticks = sdl.GetTicks()
	target_ticks = manager.lastticks + uint32(float(manager.framecount)*manager.rateticks)

	if current_ticks <= target_ticks {
		the_delay = target_ticks - current_ticks
		sdl.Delay(the_delay)
	} else {
		manager.framecount = 0
		manager.lastticks = sdl.GetTicks()
	}
}
