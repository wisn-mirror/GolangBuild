package main

import (
	"fmt"
	"math"
)

type MoveBean struct {
	x ,y float32
}

func (m MoveBean)add(addm MoveBean) MoveBean {
	return MoveBean{
		m.x+addm.x,
		m.y+addm.y,
	}
}

func (m MoveBean)Sub(subm MoveBean)MoveBean  {
	return MoveBean{
		m.x-subm.x,
		m.y-subm.y,
	}
}

func (m MoveBean)scale(f float32) MoveBean {
	return MoveBean{
		m.x*f,
		m.y*f,
	}
}
func (m MoveBean )distanceTo(other MoveBean)float32  {
	dx:=m.x-other.x
	dy:=m.y-other.y
	return float32(math.Sqrt(float64(dx*dx+dy*dy)))
}
func (m MoveBean) Mormalize()MoveBean  {
	msg:=m.x*m.x+m.y*m.y
	if msg>0{
		oneovermsg:=1/float32(math.Sqrt(float64(msg)))
		return MoveBean{
			m.x*oneovermsg,
			m.y*oneovermsg,
		}
	}
	return MoveBean{0,0}
}

type Player struct {
	current MoveBean
	target MoveBean
	speed float32
}

func( player *Player) MoveTo(movebean MoveBean)  {
	player.target=movebean
}

func (player *Player)Pos() MoveBean {
	return player.current
}
func (player *Player)isArrived()bool  {
	return player.current.distanceTo(player.target)<player.speed
}
func (player *Player)update()  {
	if !player.isArrived(){
		//当前位置的方向
		dir:=player.target.Sub(player.current).Mormalize()
		newPos:=player.current.add(dir.scale(player.speed))
		player.current=newPos
	}
}
func newPlayer(speed float32)*Player  {
	return &Player{
		speed:speed,
	}
}


func main() {
	p:=newPlayer(0.4)
	p.MoveTo(MoveBean{4,5})
	for !p.isArrived(){
		p.update()
		fmt.Println(p.Pos())
	}
}
