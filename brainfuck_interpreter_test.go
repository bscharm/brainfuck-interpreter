package main

import (
	"bytes"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BrainfuckInterpreter", func() {
	Describe("NewMachine", func() {
		It("returns a pointer to a new Machine instance", func() {
			machine := NewMachine("code", os.Stdin, os.Stdout)
			Ω(*machine).Should(BeAssignableToTypeOf(Machine{}))
		})
	})

	Describe("readChar", func() {
		It("reads into buf and stores the value in memory at current dp", func() {
			reader := strings.NewReader(">")
			machine := &Machine{
				input: reader,
				buf:   make([]byte, 1),
				dp:    4,
			}
			machine.readChar()
			Ω(machine.buf[0]).Should(Equal(byte(62)))
			Ω(machine.memory[4]).Should(Equal(62))
		})
	})

	Describe("putChar", func() {
		It("puts current value in memory at dp into buf and writes to output", func() {
			var b bytes.Buffer
			machine := &Machine{
				output: &b,
				buf:    make([]byte, 1),
				dp:     4,
			}
			machine.memory[4] = 62
			machine.putChar()
			Ω(b.Len()).Should(Equal(1))
			Ω(b.String()).Should(Equal(">"))
			Ω(b.Next(1)).Should(Equal(machine.buf))
		})
	})
})
