package main

import (
	"delay"
	"rtos"

	"arch/cortexm"
	"arch/cortexm/scb"

	"nrf51/hal/clock"
	"nrf51/hal/gpio"
	"nrf51/hal/irq"
	"nrf51/hal/rtc"
	"nrf51/hal/system"
	"nrf51/hal/timer"
)

//emgo:const
var leds = [...]byte{18, 19, 20, 21, 22}

var (
	p0   = gpio.P0
	t0   = timer.Timer0
	rtc0 = rtc.RTC0
)

const period = 2 * 32768 // 2s

func init() {
	system.Setup(clock.Xtal, clock.Xtal, true)

	for _, led := range leds {
		p0.SetMode(int(led), gpio.Out)
	}

	t0.SetPrescaler(8) // 62500 Hz
	t0.SetCC(1, 65526/2)
	t0.Event(timer.COMPARE0).EnableIRQ()
	t0.Event(timer.COMPARE1).EnableIRQ()
	rtos.IRQ(irq.Timer0).Enable()
	t0.Task(timer.START).Trigger()

	rtc0.SetPRESCALER(0) // 32768 Hz
	rtc0.Event(rtc.COMPARE1).EnableIRQ()
	rtos.IRQ(irq.RTC0).Enable()
	rtc0.SetCC(1, period)
	rtc0.Task(rtc.START).Trigger()
}

func blink(led byte, dly int) {
	p0.SetPin(int(led))
	delay.Loop(dly)
	p0.ClearPin(int(led))
}

func timerISR() {
	if e := t0.Event(timer.COMPARE0); e.IsSet() {
		e.Clear()
		blink(leds[3], 1e3)
	}
	if e := t0.Event(timer.COMPARE1); e.IsSet() {
		e.Clear()
		blink(leds[4], 1e3)
	}
}

func rtcISR() {
	rtc0.Event(rtc.COMPARE1).Clear()
	rtc0.SetCC(1, rtc0.CC(1)+period)
	blink(leds[2], 1e3)
}

func main() {
	// Sleep forever.
	scb.SCB.SLEEPONEXIT().Set()
	cortexm.DSB() // not necessary on Cortex-M0,M3,M4
	cortexm.WFI()
	// Execution should never reach there so LED0 should never light up.
	p0.SetPin(int(leds[0]))
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.Timer0: timerISR,
	irq.RTC0:   rtcISR,
}
